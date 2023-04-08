package version

import (
	"fmt"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/domain/entity"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/domain/repository"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/infrastructure/notification"
	"github.com/S-Ryouta/notice-latest-program-version/gcf/infrastructure/version"
)

type VersionInteractor struct {
	VersionRepo   repository.VersionRepository
	VersionGetter version.VersionGetter
}

func NewVersionInteractor(client repository.VersionRepository, getter version.VersionGetter) *VersionInteractor {
	return &VersionInteractor{
		VersionRepo:   client,
		VersionGetter: getter,
	}
}

func (interactor *VersionInteractor) CheckAndUpdateVersion() {
	newVersion, err := interactor.VersionGetter.GetLatestVersion()
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

		err = notification.SendNotificationToLine(newVersion)
		if err != nil {
			fmt.Println("Error notification to line:", err)
		}
		err = notification.SendNotificationToSlack(newVersion)
		if err != nil {
			fmt.Println("Error notification to slack:", err)
		}
		err = notification.SendNotificationToDiscord(newVersion)
		if err != nil {
			fmt.Println("Error notification to discord:", err)
		}
	}
}
