package router

import (
	"bastard-proxy/db"
	container "bastard-proxy/pkg/container"
	"bastard-proxy/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	metricsPkg "bastard-proxy/pkg/metrics"
	openapi "bastard-proxy/pkg/openapi"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

type RouterMap struct {
	Get    func(container.AppContainer, http.ResponseWriter, *http.Request)
	Post   func(container.AppContainer, http.ResponseWriter, *http.Request)
	Delete func(container.AppContainer, http.ResponseWriter, *http.Request)
}

func AddRoute(path string, c container.AppContainer, routerMap RouterMap) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, PATCH, POST, DELETE, OPTIONS")

		if routerMap.Get != nil && (*r).Method == "GET" {
			(routerMap.Get)(c, w, r)
			return
		}

		if routerMap.Post != nil && (*r).Method == "POST" {
			(routerMap.Post)(c, w, r)
			return
		}

		if routerMap.Post != nil && (*r).Method == "DELETE" {
			(routerMap.Delete)(c, w, r)
			return
		}

		if (*r).Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Bad Request!"))

	})

}

// ------------------- system
func HealthHandler(c container.AppContainer, w http.ResponseWriter, r *http.Request) {

	type Response struct {
		Status string `json:"status"`
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Response{"OK"})

}

// ------------------- system

var createMetricsMiddleware = func(metrics metricsPkg.Metrics) func(c *gin.Context) {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process the request
		c.Next()

		// Get the original path
		path := c.Request.URL.Path

		metricsLabels := metricsPkg.Labels{
			Method: c.Request.Method,
			Route:  path,
			Status: strconv.Itoa(c.Writer.Status()),
		}

		// Record the metrics
		metrics.HandlerExecutionTime(metricsLabels).Observe(float64(time.Since(startTime).Milliseconds()))
	}
}

func InitRouter(c container.AppContainer) {
	c.Logger.Info("Starting rest server")

	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()

	ginRouter.Use(createMetricsMiddleware(*c.Metrics))

	clientDir := "./dist"

	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "GET", "OPTIONS", "DELETE", "PATCH"},
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

	openapi.RegisterHandlers(ginRouter, NewOpenapiServer(c))

	// Start and run the server
	err := ginRouter.Run(":5000")
	if err != nil {
		panic(err)
	}
}

//

type OpenApiServer struct {
	container container.AppContainer
}

// DeleteApiGuardExcludeGuardId implements openapi.ServerInterface.
func (aps *OpenApiServer) DeleteApiGuardExcludeGuardId(c *gin.Context, guardId string, params openapi.DeleteApiGuardExcludeGuardIdParams) {
	panic("implement me")
}

// GetApiGuardExcludeGuardId implements openapi.ServerInterface.
func (aps *OpenApiServer) GetApiGuardExcludeGuardId(c *gin.Context, guardId string) {
	panic("implement me")
}

// PostApiGuardExcludeGuardId implements openapi.ServerInterface.
func (aps *OpenApiServer) PostApiGuardExcludeGuardId(c *gin.Context, guardId string) {
	panic("implement me")
}

// GetApiGuardProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) GetApiGuardProxyId(c *gin.Context, proxyId string) {
	res, _ := aps.container.PrismaClient.Guard.FindMany(db.Guard.ProxyID.Equals(proxyId)).Exec(aps.container.Context)

	// clear password from response
	for i := range res {
		res[i].Password = ""
	}

	c.JSON(http.StatusOK, res)
}

// DeleteApiGuardProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) DeleteApiGuardProxyId(c *gin.Context, proxyId string, params openapi.DeleteApiGuardProxyIdParams) {
	aps.container.PrismaClient.Guard.FindUnique(db.Guard.ID.Equals(params.Id)).Delete().Exec(aps.container.Context)

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// PostApiGuardProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) PostApiGuardProxyId(c *gin.Context, proxyId string) {
	type Guard struct {
		Email    string `json:"email" validate:"nonzero"`
		Password string `json:"password" validate:"nonzero"`
	}

	var g Guard
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pswHash, _ := utils.HashPassword(g.Password)

	aps.container.PrismaClient.Guard.CreateOne(
		db.Guard.Email.Set(g.Email),
		db.Guard.Password.Set(pswHash),
		db.Guard.ProxyID.Set(proxyId),
	).Exec(aps.container.Context)

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// PatchApiProxy implements openapi.ServerInterface.
func (aps *OpenApiServer) PatchApiProxy(c *gin.Context, params openapi.PatchApiProxyParams) {
	type Proxy struct {
		Disable bool `json:"disable" validate:"nonzero"`
		Cache   bool `json:"cache" validate:"nonzero"`
	}

	var d Proxy
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(d)
	fmt.Println(params.Id)
	aps.container.PrismaClient.Proxy.FindUnique(db.Proxy.ID.Equals(params.Id)).Update(
		db.Proxy.Disable.Set(d.Disable),
		db.Proxy.Cache.Set(d.Cache),
	).Exec(aps.container.Context)

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// DeleteApiAllowProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) DeleteApiAllowProxyId(c *gin.Context, proxyId string, params openapi.DeleteApiAllowProxyIdParams) {
	aps.container.PrismaClient.Allow.FindUnique(db.Allow.ID.Equals(params.Id)).Delete().Exec(aps.container.Context)

	// aps.container.RefetchBlockMap()

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// GetApiAllowProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) GetApiAllowProxyId(c *gin.Context, proxyId string) {
	res, _ := aps.container.PrismaClient.Allow.FindMany(db.Allow.ProxyID.Equals(proxyId)).Exec(aps.container.Context)

	c.JSON(http.StatusOK, res)
}

// PostApiAllowProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) PostApiAllowProxyId(c *gin.Context, proxyId string) {
	type Allow struct {
		Ip string `json:"ip" validate:"nonzero"`
	}

	var a Allow
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	aps.container.PrismaClient.Allow.CreateOne(
		db.Allow.IP.Set(a.Ip),
		db.Allow.ProxyID.Set(proxyId),
	).Exec(aps.container.Context)

	// aps.container.RefetchBlockMap()

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// DeleteApiBlockProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) DeleteApiBlockProxyId(c *gin.Context, proxyId string, params openapi.DeleteApiBlockProxyIdParams) {
	aps.container.PrismaClient.Block.FindUnique(db.Block.ID.Equals(params.Id)).Delete().Exec(aps.container.Context)

	aps.container.RefetchBlockMap()

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// GetApiBlockProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) GetApiBlockProxyId(c *gin.Context, proxyId string) {
	res, _ := aps.container.PrismaClient.Block.FindMany(db.Block.ProxyID.Equals(proxyId)).Exec(aps.container.Context)

	c.JSON(http.StatusOK, res)
}

// PostApiBlockProxyId implements openapi.ServerInterface.
func (aps *OpenApiServer) PostApiBlockProxyId(c *gin.Context, proxyId string) {
	type Block struct {
		Ip string `json:"ip" validate:"nonzero"`
	}

	var b Block
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	aps.container.PrismaClient.Block.CreateOne(
		db.Block.IP.Set(b.Ip),
		db.Block.ProxyID.Set(proxyId),
	).Exec(aps.container.Context)

	aps.container.RefetchBlockMap()

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

// GetApiBlock implements openapi.ServerInterface.
func (aps *OpenApiServer) GetApiBlock(c *gin.Context) {
	res, _ := aps.container.PrismaClient.Block.FindMany().Exec(aps.container.Context)

	c.JSON(http.StatusOK, res)
}

// DeleteProxy implements openapi.ServerInterface.
func (aps *OpenApiServer) DeleteApiProxy(c *gin.Context, params openapi.DeleteApiProxyParams) {
	aps.container.PrismaClient.Proxy.FindUnique(db.Proxy.ID.Equals(params.Id)).Delete().Exec(aps.container.Context)

	aps.container.RefetchDomainMap()

	c.JSON(http.StatusOK, &openapi.SuccessResponse{Message: "OK"})
}

func (aps *OpenApiServer) GetApiProxy(c *gin.Context) {
	res, _ := aps.container.PrismaClient.Proxy.FindMany().Exec(aps.container.Context)
	fmt.Println(res)
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
		db.Proxy.Disable.Set(false),
		db.Proxy.Cache.Set(false),
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
