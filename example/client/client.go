package main

import (
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

const (
	authServerURL = "http://localhost:9096"
)

var (
	config = oauth2.Config{
		ClientID:     "222222",
		ClientSecret: "22222222",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:9094/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}
	globalToken *oauth2.Token // Non-concurrent security
)

func main() {
	//授权码模式Authorization Code
	//访问第三方授权页
	http.HandleFunc("/", index)
	//由三方鉴权服务重定向返回，拿到code，并请求和验证token
	http.HandleFunc("/oauth2", oAuth2)
	//刷新验证码
	http.HandleFunc("/refresh", refresh)
	http.HandleFunc("/try", try)

	//密码模式Resource Owner Password Credentials
	http.HandleFunc("/pwd", pwd)

	//客户端模式Client Credentials
	http.HandleFunc("/client", client)

	log.Println("Client is running at 9094 port.Please open http://localhost:9094")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
