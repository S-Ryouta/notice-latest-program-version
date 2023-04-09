package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var (
	EndOfLifeUrl = "https://endoflife.date"
)

type EndOfLifeResponse struct {
	Cycle             string `json:"cycle"`
	Latest            string `json:"latest"`
	ReleaseDate       string `json:"releaseDate"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	Lts               bool   `json:"lts"`
}

type EndOfLifeErrorResponse struct {
	Message string `json:"message"`
}

type VersionGetter interface {
	GetLatestVersion() (string, error)
}

type defaultVersionGetter struct{}

func (d *defaultVersionGetter) GetLatestVersion() (string, error) {
	response, err := http.Get(EndOfLifeUrl + "/api/go.json")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		var endOfLifeErrorResponse EndOfLifeErrorResponse

		err = json.Unmarshal(body, &endOfLifeErrorResponse)
		if err != nil {
			return "", fmt.Errorf("Error response unmarshal failed: %+v \n", string(body))
		}
		errorMessage := fmt.Sprintf("connection is failed for EndOfLife. status code: %d, error_message: %s",
			response.StatusCode,
			endOfLifeErrorResponse.Message,
		)
		return "", fmt.Errorf("EndOfLifeErrorResponse: " + errorMessage)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var versions []EndOfLifeResponse
	err = json.Unmarshal(body, &versions)
	if err != nil {
		return "", err
	}

	// Find the latest stable version
	r := regexp.MustCompile(`^[0-9]+\.[0-9]+(\.[0-9]+)?$`)
	for _, v := range versions {
		if r.MatchString(v.Latest) {
			return v.Latest, nil
		}
	}

	return "", fmt.Errorf("latest stable Golang version not found")
}

func NewDefaultVersionGetter() VersionGetter {
	return &defaultVersionGetter{}
}
