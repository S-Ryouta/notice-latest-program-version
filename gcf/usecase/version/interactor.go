package version

import (
	"fmt"
	"notice-latest-program-version/domain/entity"
	"notice-latest-program-version/domain/repository"
	"notice-latest-program-version/infrastructure/notification"
	"notice-latest-program-version/infrastructure/version"
)

type VersionInteractor struct {
	VersionRepo repository.VersionRepository
}

func NewVersionInteractor(client repository.VersionRepository) *VersionInteractor {
	return &VersionInteractor{
		VersionRepo: client,
	}
}

func (interactor *VersionInteractor) CheckAndUpdateVersion() {
	newVersion, err := version.GetLatestVersion()
	if err != nil {
		fmt.Println("Error getting the latest Golang version:", err)
		return
	}

	storedVersionEntity, err := interactor.VersionRepo.GetVersion()
	if err != nil {
		fmt.Println("Error getting the stored Golang version:", err)
		return
	}

	if storedVersionEntity.Version != newVersion {
		err = interactor.VersionRepo.SaveVersion(&entity.Version{
			ID:      "golang",
			Version: newVersion,
		})
		if err != nil {
			fmt.Println("Error saving the new Golang version:", err)
			return
		}

		notification.SendNotificationToLine(newVersion)
		notification.SendNotificationToDiscord(newVersion)
	}
}
