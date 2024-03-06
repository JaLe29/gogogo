package main

import (
	db "bastard-proxy/db"
	container "bastard-proxy/pkg/container"
	"bastard-proxy/pkg/metrics"
	router "bastard-proxy/pkg/router"
	"bastard-proxy/pkg/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	config "bastard-proxy/pkg/config"

	"github.com/coocood/freecache"
)

func main() {

	conf, err := config.Load()

	if err != nil {
		panic(err)
	}

	c := container.InitContainer(conf)

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
		requestHost := strings.TrimPrefix(r.Host, "www.")
		fmt.Println("req to: " + requestHost)
		if proxyObj, ok := (*container.DomainMap)[requestHost]; ok {
			clientIp := utils.ReadUserIP(r)

			fmt.Println("req to: " + requestHost + ", target:" + proxyObj.TargetProxy + ", clientIp: " + clientIp)

			c := *container.Metrics
			c.HandleActivity(metrics.Activity{Ip: clientIp, ProxyId: proxyObj.Id, Host: requestHost})

			hasAllowMap := (*container.AllowMap)[proxyObj.Id] != nil
			isGuardEnabled := (*container.GuardMap)[proxyObj.Id]
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

			if isGuardEnabled {
				authToken, errorAuthToken := r.Cookie("ap-token")

				if authToken == nil || errorAuthToken != nil {
					loginForm := []byte(
						`
						<!DOCTYPE html>
						<html lang="cs">
						<head>
							<meta charset="UTF-8">
							<meta name="viewport" content="width=device-width, initial-scale=1.0">
							<title>Přihlašovací formulář</title>
							<script>

								function submitform() {
										var form = document.getElementById('form');
										var xhr = new XMLHttpRequest();
										var formData = new FormData(form);
										//open the request
										xhr.open('POST', '/__system/login')
										xhr.setRequestHeader("Content-Type", "application/json");

										//send the form data
										xhr.send(JSON.stringify(Object.fromEntries(formData)));

										xhr.onreadystatechange = function() {
											if (xhr.readyState == XMLHttpRequest.DONE) {
												form.reset(); //reset form after AJAX success or do something else
												if (xhr.response === "OK") {
													location.reload();
												}
											}
										}
										//Fail the onsubmit to avoid page refresh.
										return false;
								}
							</script>
						</head>
							<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0;">

							<div style="background-color: #fff; padding: 20px; border-radius: 5px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); width: 300px;">
								<form id="form" style="display: flex; flex-direction: column; gap: 10px;">
									<h2 style="text-align: center; margin: 0 0 20px 0;">Přihlášení</h2>
									<div style="display: flex; flex-direction: column;">
										<label for="email" style="margin-bottom: 5px;">E-mail</label>
										<input type="email" id="email" name="email" required style="padding: 10px; border: 1px solid #ccc; border-radius: 4px;">
									</div>
										<div style="display: flex; flex-direction: column;">
										<label for="password" style="margin-bottom: 5px;">Heslo</label>
										<input type="password" id="password" name="password" required style="padding: 10px; border: 1px solid #ccc; border-radius: 4px;">
										<input type="submit" hidden onsubmit="submitform()" />
									</div>
									<button value="Submit" type="button" onclick="submitform()" style="padding: 10px; border: none; border-radius: 4px; background-color: #007bff; color: white; cursor: pointer;">Přihlásit se</button>
								</form>
							</div>

							</body>
						</html>
					`,
					)

					w.WriteHeader(http.StatusOK)
					w.Write(loginForm)
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

		} else {
			log.Println("Invalid request " + requestHost)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("400 - Bad Request!"))
		}
	})

	http.HandleFunc("/__system/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "ap-token",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/__system/metrics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// days := 7
		// container.PrismaClient.Activity.FindMany(db.Activity.CreatedAt.Lt(time.Now().AddDate(0, 0, -1*days))).Delete().Exec(container.Context)

		m := *container.Metrics
		m.Handler().ServeHTTP(w, r)

		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte("OK"))
	})

	http.HandleFunc("/__system/login", func(w http.ResponseWriter, r *http.Request) {

		proxyId := (*container.DomainMap)[r.Host].Id

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		type LoginPayload struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		l := &LoginPayload{}
		err := json.NewDecoder(r.Body).Decode(l)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		guardResponse, err := container.PrismaClient.Guard.FindFirst(
			db.Guard.ProxyID.Equals(proxyId),
			db.Guard.Email.Equals(l.Email),
		).Exec(container.Context)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		fmt.Println(guardResponse)
		if guardResponse == nil {
			http.Error(w, "Unauthorized 1", http.StatusUnauthorized)
			return
		}

		newsPswHash, _ := utils.HashPassword(l.Password)
		fmt.Println(utils.CheckPasswordHash(l.Password, newsPswHash))

		if !utils.CheckPasswordHash(l.Password, guardResponse.Password) {
			http.Error(w, "Unauthorized 2", http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "ap-token",
			Value:  utils.GenerateJwtToken(guardResponse.ID),
			Path:   "/",
			MaxAge: container.Config.CokieTTL,
		})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.AddRoute("/_system/health", container, router.RouterMap{Get: router.HealthHandler})

	log.Println("--------------------------------- Proxy server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	// -------------------------------
}
