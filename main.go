package main

import (
	db "bastard-proxy/db"
	container "bastard-proxy/pkg/container"
	router "bastard-proxy/pkg/router"
	"bastard-proxy/pkg/utils"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	c := container.InitContainer()

	go startProxy(c)
	router.InitRouter(c)
}

func startProxy(container container.AppContainer) {
	container.Logger.Info("Starting proxy server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if proxyObj, ok := (*container.DomainMap)[r.Host]; ok {
			clientIp := utils.ReadUserIP(r)

			hasAllowMap := (*container.AllowMap)[proxyObj.Id] != nil

			if hasAllowMap && !(*container.AllowMap)[proxyObj.Id][clientIp] {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403 - Forbidden!"))
				return
			}

			if !hasAllowMap {
				if (*container.BlockMap)[proxyObj.Id] != nil && ((*container.BlockMap)[proxyObj.Id][clientIp] || (*container.BlockMap)[proxyObj.Id]["*"]) {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte("403 - Forbidden!"))
					return
				}
			}

			proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: proxyObj.TargetProxy})

			proxy.ServeHTTP(w, r)

			container.PrismaClient.Activity.CreateOne(
				db.Activity.IP.Set(clientIp),
				db.Activity.ProxyID.Set(proxyObj.Id),
			).Exec(container.Context)

		} else {
			log.Println("Invalid request " + r.Host)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("400 - Bad Request!"))
		}
	})

	router.AddRoute("/_system/health", container, router.RouterMap{Get: router.HealthHandler})

	log.Println("--------------------------------- Proxy server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	// -------------------------------
}
