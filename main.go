package main

import (
	"encoding/json"
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
		log.Fatal("VAULT_ADDR must be set.")
	}
	vaultSecret := os.Getenv("VAULT_SECRET")
	if vaultSecret == "" {
		log.Fatal("VAULT_SECRET must be set.")
	}
	vaultRole := os.Getenv("VAULT_ROLE")
	if vaultRole == "" {
		log.Fatal("VAULT_ROLE must be set.")
	}

	var err error

	secret, err = gcpss.FetchVaultSecret(vaultAddr, vaultSecret, vaultRole)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	http.HandleFunc("/", getSecret)
	http.ListenAndServe(":8080", nil)
}

func getSecret(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal(secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}