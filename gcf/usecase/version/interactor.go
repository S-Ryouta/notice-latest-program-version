package version

import (
	"fmt"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/domain/entity"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/domain/repository"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/infrastructure/notification"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/infrastructure/version"
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

	fmt.Printf("new version detail: %s \n", newVersion)
	fmt.Printf("current version detail: %+v \n", storedVersionEntity)
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
