/*
 * wcafe
 *
 * wcafe store
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	StoresApiService := openapi.NewStoresApiService()
	StoresApiController := openapi.NewStoresApiController(StoresApiService)

	router := openapi.NewRouter(StoresApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}