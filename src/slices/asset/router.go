package asset_slice

import (
	"log"
	"net/http"
)

func AssetAPIRouting(router *http.ServeMux) {
	router.HandleFunc("GET /assets/{id}", GetAsset)
	router.HandleFunc("POST /assets", CreateAsset)
	router.HandleFunc("DELETE /blobs/{id}", DeleteAsset)
	log.Println("Assets API added")
}
