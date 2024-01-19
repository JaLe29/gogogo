package main

import (
	db "bastard-proxy/db"
	container "bastard-proxy/pkg/container"
	router "bastard-proxy/pkg/router"
	"bastard-proxy/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/validator.v2"
)

func main() {
	godotenv.Load()

	c := container.InitContainer()
	c.RefetchDomainMap()

	go startGin(c)
	startProxy(c)
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
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	ginRouter.Use(static.Serve("/", static.LocalFile(clientDir, true)))
	ginRouter.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.File(clientDir + "/index.html")
		}
		//default 404 page not found
	})

	// Setup route group for the API
	api := ginRouter.Group("/api")
	{
		api.GET("/proxy", func(c *gin.Context) {
			res, _ := container.PrismaClient.Proxy.FindMany().Exec(container.Context)

			c.JSON(http.StatusOK, res)
		})

		api.DELETE("/proxy/:id", func(c *gin.Context) {
			id := c.Param("id")
			res, _ := container.PrismaClient.Proxy.FindUnique(db.Proxy.ID.Equals(id)).Delete().Exec(container.Context)

			c.JSON(http.StatusOK, res)
		})

		api.POST("/proxy", func(c *gin.Context) {
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

			if err := validator.Validate(p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			res, _ := container.PrismaClient.Proxy.CreateOne(
				db.Proxy.Source.Set(p.Source),
				db.Proxy.Target.Set(p.Target),
			).Exec(container.Context)

			container.RefetchDomainMap()

			c.JSON(http.StatusOK, res)
		})
	}
	log.Println("--------------------------------- Gin server started on port 5000")
	// Start and run the server
	ginRouter.Run(":5000")

}
