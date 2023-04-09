package version_test

import (
	"errors"
	"testing"

	"github.com/S-Ryouta/notice-latest-program-version/domain/entity"
	"github.com/S-Ryouta/notice-latest-program-version/infrastructure/version"
	usecaseVersion "github.com/S-Ryouta/notice-latest-program-version/usecase/version"
	"github.com/stretchr/testify/mock"
)

type MockVersionRepository struct {
	mock.Mock
}

func (m *MockVersionRepository) GetVersion() (*entity.Version, error) {
	args := m.Called()
	return args.Get(0).(*entity.Version), args.Error(1)
}

func (m *MockVersionRepository) SaveVersion(version *entity.Version) error {
	args := m.Called(version)
	return args.Error(0)
}

type mockVersionGetter struct {
	version.VersionGetter
	getLatestVersionFunc func() (string, error)
}

func (m *mockVersionGetter) GetLatestVersion() (string, error) {
	return m.getLatestVersionFunc()
}

func TestCheckAndUpdateVersion(t *testing.T) {
	testCases := []struct {
		name                string
		getVersionError     error
		saveVersionError    error
		getLatestVersionErr error
		storedVersion       string
		newVersion          string
		expectSave          bool
	}{
		{
			name:          "正常: 新しいバージョンがリリースされた場合",
			storedVersion: "1.0.0",
			newVersion:    "1.1.0",
			expectSave:    true,
		},
		{
			name:          "正常: 新しいバージョンがリリースされていない場合",
			storedVersion: "1.0.0",
			newVersion:    "1.0.0",
			expectSave:    false,
		},
		{
			name:            "エラー: GetVersion エラー",
			getVersionError: errors.New("GetVersion error"),
			expectSave:      false,
		},
		{
			name:             "エラー: SaveVersion エラー",
			storedVersion:    "1.0.0",
			newVersion:       "1.1.0",
			saveVersionError: errors.New("SaveVersion error"),
			expectSave:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockVersionRepository)
			mockRepo.On("GetVersion").Return(&entity.Version{Version: tc.storedVersion}, tc.getVersionError)
			mockRepo.On("SaveVersion", &entity.Version{ID: "golang", Version: tc.newVersion}).Return(tc.saveVersionError)

			mockVersionGetter := &mockVersionGetter{
				getLatestVersionFunc: func() (string, error) {
					return tc.newVersion, tc.getLatestVersionErr
				},
			}

			versionInteractor := usecaseVersion.NewVersionInteractor(mockRepo, mockVersionGetter)
			versionInteractor.CheckAndUpdateVersion()

			if tc.expectSave {
				mockRepo.AssertCalled(t, "SaveVersion", &entity.Version{ID: "golang", Version: tc.newVersion})
			} else {
				mockRepo.AssertNotCalled(t, "SaveVersion", &entity.Version{ID: "golang", Version: tc.newVersion})
			}
		})
	}
}
