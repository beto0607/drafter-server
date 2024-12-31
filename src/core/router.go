package core

import (
	asset_slice "drafter/src/slices/asset"
	project_slice "drafter/src/slices/project"
	"log"
	"net/http"
)

func InitRouting() *http.ServeMux {
	apiRouter := doApiRouting()
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))
	return router
}

func doApiRouting() *http.ServeMux {
	log.Println("Doing API routing...")
	router := http.NewServeMux()
	asset_slice.AssetAPIRouting(router)
	project_slice.ProjectAPIRouting(router)

	return router
}
