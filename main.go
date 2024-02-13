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
	// "github.com/golang-jwt/jwt/v5"
)

func main() {
	/*
		secret := []byte("dssdadsasdasdasda")
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "john",
			"foo": 2,
		})
		s, e := t.SignedString(secret)
		fmt.Println(e)
		fmt.Println(s)

		// sample token string taken from the New example
		tokenString := s

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secret, nil
		})
		if err != nil {
			log.Fatal(err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims)
		} else {
			fmt.Println(err)
		}
	*/

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

			authToken, errorAuthToken := r.Cookie("ap-token")
			// fmt.Println(authToken.Value)
			if authToken == nil || errorAuthToken != nil {
				// adminDomain := os.Getenv("ADMIN_DOMAIN")
				loginForm := []byte(
					`
						<!DOCTYPE html>
						<html lang="cs">
						<head>
							<meta charset="UTF-8">
							<meta name="viewport" content="width=device-width, initial-scale=1.0">
							<title>Přihlašovací formulář</title>
							<script>	

								function setCookie(name,value,days) {
									var expires = "";
									if (days) {
										var date = new Date();
										date.setTime(date.getTime() + (days*24*60*60*1000));
										expires = "; expires=" + date.toUTCString();
									}
									document.cookie = name + "=" + (value || "")  + expires + "; path=/";
								}

								function getCookie(name) {
									var nameEQ = name + "=";
									var ca = document.cookie.split(';');
									for(var i=0;i < ca.length;i++) {
										var c = ca[i];
										while (c.charAt(0)==' ') c = c.substring(1,c.length);
										if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
									}
									return null;
								}

								function eraseCookie(name) {   
									document.cookie = name +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
								}

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

	http.HandleFunc("/__system/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "ap-token",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/__system/login", func(w http.ResponseWriter, r *http.Request) {

		http.SetCookie(w, &http.Cookie{
			Name:   "ap-token",
			Value:  utils.GenerateJwtToken(),
			Path:   "/",
			MaxAge: 3600,
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
