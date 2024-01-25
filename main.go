package main

import (
	db "bastard-proxy/db"
	container "bastard-proxy/pkg/container"
	openapi "bastard-proxy/pkg/openapi"
	router "bastard-proxy/pkg/router"
	"bastard-proxy/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	middleware "github.com/oapi-codegen/gin-middleware"
)

type OpenApiServer struct {
	container container.AppContainer
}

// DeleteProxy implements openapi.ServerInterface.
func (aps *OpenApiServer) DeleteApiProxy(c *gin.Context, params openapi.DeleteApiProxyParams) {
	aps.container.PrismaClient.Proxy.FindUnique(db.Proxy.ID.Equals(params.Id)).Delete().Exec(aps.container.Context)

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// GetActivity implements openapi.ServerInterface.
func (aps *OpenApiServer) GetApiActivity(c *gin.Context) {
	res, _ := aps.container.PrismaClient.Activity.FindMany().Exec(aps.container.Context)

	c.JSON(http.StatusOK, res)
}

func (aps *OpenApiServer) GetApiProxy(c *gin.Context) {
	res, _ := aps.container.PrismaClient.Proxy.FindMany().Exec(aps.container.Context)

	c.JSON(http.StatusOK, res)
}

func (aps *OpenApiServer) PostApiProxy(c *gin.Context) {
	type Proxy struct {
		Target string `json:"target" validate:"nonzero"`
		Source string `json:"source" validate:"nonzero"`
	}

	var p Proxy
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, _ := aps.container.PrismaClient.Proxy.CreateOne(
		db.Proxy.Source.Set(p.Source),
		db.Proxy.Target.Set(p.Target),
	).Exec(aps.container.Context)

	aps.container.RefetchDomainMap()

	c.JSON(http.StatusOK, res)
}

// Struktura implementující ServerInterface
var _ openapi.ServerInterface = (*OpenApiServer)(nil)

func NewOpenapiServer(container container.AppContainer) *OpenApiServer {
	return &OpenApiServer{
		container: container,
	}
}

func newSwaggerServer() *openapi3.T {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil

	return swagger
}

func main() {
	godotenv.Load()

	c := container.InitContainer()
	c.RefetchDomainMap()

	go startProxy(c)
	startGin(c)
}

func startProxy(container container.AppContainer) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println((*container.DomainMap))
		if proxyObj, ok := (*container.DomainMap)[r.Host]; ok {
			// clientIp := utils.ReadUserIP(r)

			proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: proxyObj.TargetProxy})
			// if clientCounterMap[target] == nil {
			// 	clientCounterMap[target] = map[string]int32{}
			// }
			// clientCounterMap[target][clientIp]++

			proxy.ServeHTTP(w, r)

			container.PrismaClient.Activity.CreateOne(
				db.Activity.IP.Set(utils.ReadUserIP(r)),
				db.Activity.ProxyID.Set(proxyObj.Id),
			).Exec(container.Context)

		} else {
			log.Println("Invalid request " + r.Host)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("400 - Bad Request!"))
		}
	})

	router.AddRoute("/_system/health", container, router.RouterMap{Get: router.HealthHandler})
	// router.AddRoute("/_system/proxy", container, router.RouterMap{
	// 	Get:    router.GetProxy,
	// 	Post:   router.PostProxy,
	// 	Delete: router.DeleteProxy,
	// })

	log.Println("--------------------------------- Proxy server started on port 8080")
	http.ListenAndServe(":8080", nil)

	// -------------------------------
}

func startGin(container container.AppContainer) {
	ginRouter := gin.Default()

	clientDir := "./dist"

	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "GET", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	ginRouter.Use(static.Serve("/", static.LocalFile(clientDir, true)))

	ginRouter.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.File(clientDir + "/index.html")
		}
	})

	ginRouter.Use(middleware.OapiRequestValidator(newSwaggerServer()))

	openapi.RegisterHandlers(ginRouter, NewOpenapiServer(container))

	// Start and run the server
	ginRouter.Run(":5000")

}
