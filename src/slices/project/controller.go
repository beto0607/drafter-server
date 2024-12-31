package project_slice

import (
	"drafter/src/utils"
	"encoding/json"
	"log"
	"net/http"
)

func GetProject(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("id")
	if len(projectId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	project, err := FindProjectEntry(projectId)
	if err != nil {
		log.Printf("Couldn't find %s\n", projectId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(projectId)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var projectDto ProjectDto

	err := json.NewDecoder(r.Body).Decode(&projectDto)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if validationErr := utils.Validator.Struct(&projectDto.Data); validationErr != nil {
		log.Println(validationErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newProject, err := SaveProjectEntry(projectDto.Data)

	if err != nil {
		log.Println("Couln't insert document")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProject)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("id")

	if len(projectId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var projectDto ProjectDto

	err := json.NewDecoder(r.Body).Decode(&projectDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = UpdateProjectEntry(projectId, &projectDto.Data)
	if err != nil {
		log.Println("Couln't update document")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projectDto.Data)

}

func DeleteProject(w http.ResponseWriter, r *http.Request)      {}
func MoveProjectToTrash(w http.ResponseWriter, r *http.Request) {}
