package main

import (
	"crypto/sha256"
	"encoding/base64"
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
	http.HandleFunc("/", index)

	http.HandleFunc("/oauth2", oAuth2)

	http.HandleFunc("/refresh", refresh)

	http.HandleFunc("/try", try)

	http.HandleFunc("/pwd", pwd)

	http.HandleFunc("/client", client)

	log.Println("Client is running at 9094 port.Please open http://localhost:9094")
	log.Fatal(http.ListenAndServe(":9094", nil))
}

func genCodeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))
	return base64.URLEncoding.EncodeToString(s256[:])
}
