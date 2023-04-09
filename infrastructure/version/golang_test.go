package version_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/S-Ryouta/notice-latest-program-version/infrastructure/version"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestVersion(t *testing.T) {
	testCases := []struct {
		name               string
		responseStatusCode int
		responseBody       string
		expectError        bool
		expectedVersion    string
	}{
		{
			name:               "Success",
			responseStatusCode: http.StatusOK,
			responseBody: `[{
				"cycle": "1.16.x",
				"latest": "1.16.8",
				"releaseDate": "2021-02-16",
				"latestReleaseDate": "2021-10-25",
				"lts": false
			}]`,
			expectError:     false,
			expectedVersion: "1.16.8",
		},
		{
			name:               "ErrorParsingResponseBody",
			responseStatusCode: http.StatusOK,
			responseBody:       `invalid JSON`,
			expectError:        true,
			expectedVersion:    "",
		},
		{
			name:               "ErrorNon200Response",
			responseStatusCode: http.StatusInternalServerError,
			responseBody: `{
				"message": "Internal server error"
			}`,
			expectError:     true,
			expectedVersion: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.responseStatusCode)
				fmt.Fprintln(w, tc.responseBody)
			}))
			defer ts.Close()

			// Override the EndOfLifeUrl for testing
			originalEndOfLifeUrl := version.EndOfLifeUrl
			version.EndOfLifeUrl = ts.URL

			versionGetter := version.NewDefaultVersionGetter()
			gotVersion, err := versionGetter.GetLatestVersion()

			// Restore the original EndOfLifeUrl
			version.EndOfLifeUrl = originalEndOfLifeUrl

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedVersion, gotVersion)
			}
		})
	}
}
