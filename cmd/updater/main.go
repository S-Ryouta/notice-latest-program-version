package main

import (
	"net/http"

	"github.com/yourusername/yourproject/infrastructure/db"
	"github.com/yourusername/yourproject/usecase/version"
)

func CheckAndUpdateVersionHandler(w http.ResponseWriter, r *http.Request) {
	client := db.CreateRedisClient()
	defer client.Close()

	versionInteractor := version.NewVersionInteractor(client)
	versionInteractor.CheckAndUpdateVersion()
}

func main() {
	http.HandleFunc("/", CheckAndUpdateVersionHandler)
	http.ListenAndServe(":8080", nil)
}
