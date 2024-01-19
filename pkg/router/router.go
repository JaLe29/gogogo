package router

import (
	db "bastard-proxy/db"
	container "bastard-proxy/pkg/container"
	utils "bastard-proxy/pkg/utils"
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
// ------------------- proxy
func GetProxy(c container.AppContainer, w http.ResponseWriter, r *http.Request) {
	res, _ := c.PrismaClient.Proxy.FindMany().Exec(c.Context)
	json.NewEncoder(w).Encode(&res)
}

func PostProxy(c container.AppContainer, w http.ResponseWriter, r *http.Request) {
	type Proxy struct {
		Target string `json:"target" validate:"required"`
		Source string `json:"source" validate:"required"`
	}

	var p Proxy
	if !utils.ValidateAndProcessData(w, r, &p) {
		return
	}

	res, _ := c.PrismaClient.Proxy.CreateOne(
		db.Proxy.Source.Set(p.Source),
		db.Proxy.Target.Set(p.Target),
	).Exec(c.Context)

	c.RefetchDomainMap()

	json.NewEncoder(w).Encode(&res)
}

func DeleteProxy(c container.AppContainer, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id != "" {
		c.PrismaClient.Proxy.FindUnique(
			db.Proxy.ID.Equals(id),
		).Delete().Exec(c.Context)

		type Response struct {
			Status string `json:"status"`
		}

		c.RefetchDomainMap()

		json.NewEncoder(w).Encode(&Response{"OK"})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Bad Request!"))
		return
	}

}

// ------------------- proxy
