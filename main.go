package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BESTSELLER/go-vault/gcpss"
)

var secret string

func init() {
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		log.Fatal("VAULT_ADDR environment variable must be set.")
	}
	vaultSecret := os.Getenv("VAULT_SECRET")
	if vaultSecret == "" {
		log.Fatal("VAULT_SECRET environment variable must be set.")
	}
	vaultRole := os.Getenv("VAULT_ROLE")
	if vaultRole == "" {
		log.Fatal("VAULT_ROLE environment variable must be set.")
	}

	var err error

	secret, err = gcpss.FetchVaultSecret(vaultAddr, vaultSecret, vaultRole)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	log.Println(secret)
	http.HandleFunc("/", getSecret)
	http.ListenAndServe(":8080", nil)
}

func getSecret(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(secret))
}
