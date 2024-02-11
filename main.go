package main

import (
	db "bastard-proxy/db"
	container "bastard-proxy/pkg/container"
	router "bastard-proxy/pkg/router"
	"bastard-proxy/pkg/utils"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/coocood/freecache"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	c := container.InitContainer()

	go startProxy(c)
	router.InitRouter(c)
}

type transport struct {
	http.RoundTripper
	cache *freecache.Cache
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1)
	body := io.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))

	header := resp.Header.Get("Ap-Cache-Control")
	if header != "" {
		headerIntValue, err := strconv.Atoi(header)
		if err == nil {
			fmt.Println("Caching for ", headerIntValue, " seconds", " - ", req.URL.String())
			t.cache.Set([]byte(req.URL.String()), b, headerIntValue)
		}
	}

	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	return resp, nil
}

func startProxy(container container.AppContainer) {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)

	container.Logger.Info("Starting proxy server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if proxyObj, ok := (*container.DomainMap)[r.Host]; ok {
			clientIp := utils.ReadUserIP(r)

			hasAllowMap := (*container.AllowMap)[proxyObj.Id] != nil
			isDisabled := proxyObj.Disable

			if (hasAllowMap && !(*container.AllowMap)[proxyObj.Id][clientIp]) || isDisabled {
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

			isCacheEnabled := proxyObj.Cache

			if isCacheEnabled {
				cacheResponse, _ := cache.Get([]byte("http://" + proxyObj.TargetProxy + r.URL.String()))
				// fmt.Println(cacheResponse)
				if cacheResponse != nil {
					// fmt.Println("Cache hit")
					w.Header().Add("Ap-Cache-Status", "HIT")
					w.WriteHeader(http.StatusOK)
					w.Write(cacheResponse)
					return
				}
			}

			proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: proxyObj.TargetProxy})

			if isCacheEnabled {
				proxy.Transport = &transport{http.DefaultTransport, cache}
			}

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
