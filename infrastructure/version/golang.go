package version

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type GolangVersionInfo struct {
	Version string `json:"version"`
}

func GetLatestGolangVersion() (string, error) {
	response, err := http.Get("https://golang.org/dl/?mode=json")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var versions []GolangVersionInfo
	err = json.Unmarshal(body, &versions)
	if err != nil {
		return "", err
	}

	// Find the latest stable version
	r := regexp.MustCompile(`^go[0-9]+\.[0-9]+(\.[0-9]+)?$`)
	for _, v := range versions {
		if r.MatchString(v.Version) {
			return v.Version, nil
		}
	}

	return "", fmt.Errorf("latest stable Golang version not found")
}
