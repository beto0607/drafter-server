package project_slice

import (
	"drafter/src/utils"
	"log"
	"net/http"
)

func ProjectAPIRouting(router *http.ServeMux) {
	router.HandleFunc("OPTIONS /projects", utils.PreflightHandler)
	router.HandleFunc("OPTIONS /projects/{id}", utils.PreflightHandler)
	router.HandleFunc("GET /projects/{id}", utils.CorsMiddleware(GetProject))
	router.HandleFunc("POST /projects", utils.CorsMiddleware(CreateProject))
	router.HandleFunc("PUT /projects/{id}", utils.CorsMiddleware(UpdateProject))
	router.HandleFunc("DELETE /projects/{id}", utils.CorsMiddleware(DeleteProject))
	router.HandleFunc("DELETE /projects/{id}/moveToTrash", utils.CorsMiddleware(MoveProjectToTrash))
	log.Println("Project API added")
}
