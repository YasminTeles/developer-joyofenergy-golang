package standard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

var (
	GitCommit string
	GitTag    string
	BuildData string
)

func Version(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")

	version := getVersion()

	err := json.NewEncoder(response).Encode(version)
	if err != nil {
		log.Fatal("Error happened in JSON marshal.")
	}
}

func getVersion() map[string]string {
	version := make(map[string]string)
	version["version"] = GitTag
	version["build"] = strings.Replace(BuildData, "\t", " ", 1)
	version["commit"] = GitCommit

	return version
}

func Healthcheck(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(response, "Working!")
}
