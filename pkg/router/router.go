package router

import (
	container "bastard-proxy/pkg/container"
	"encoding/json"
	"net/http"
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
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

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
