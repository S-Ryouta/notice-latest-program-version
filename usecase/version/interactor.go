package version

import (
	"fmt"
	"github.com/S-Ryouta/notice-latest-program-version/domain/entity"
	"github.com/S-Ryouta/notice-latest-program-version/domain/repository"
	"github.com/S-Ryouta/notice-latest-program-version/infrastructure/notification"
	"github.com/S-Ryouta/notice-latest-program-version/infrastructure/version"
)

type VersionInteractor struct {
	VersionRepo   repository.VersionRepository
	VersionGetter version.VersionGetter
	language      string
}

func NewVersionInteractor(client repository.VersionRepository, getter version.VersionGetter, language string) *VersionInteractor {
	return &VersionInteractor{
		VersionRepo:   client,
		VersionGetter: getter,
		language:      language,
	}
}

func (interactor *VersionInteractor) CheckAndUpdateVersion() {
	newVersion, err := interactor.VersionGetter.GetLatestVersion(interactor.language)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error getting the latest %s version:", interactor.language), err)
		return
	}

	storedVersionEntity, err := interactor.VersionRepo.GetVersion(interactor.language)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error getting the stored %s version:", interactor.language), err)
		return
	}

	fmt.Printf("new %s version detail: %s \n", interactor.language, newVersion)
	fmt.Printf("current %s version detail: %+v \n", interactor.language, storedVersionEntity)
	if storedVersionEntity.Version != newVersion {
		err = interactor.VersionRepo.SaveVersion(&entity.Version{
			ID:      interactor.language,
			Version: newVersion,
		})
		if err != nil {
			fmt.Println("Error saving the new Golang version:", err)
			return
		}

		err = notification.SendNotificationToLine(interactor.language, newVersion)
		if err != nil {
			fmt.Println("Error notification to line:", err)
		}
		err = notification.SendNotificationToSlack(interactor.language, newVersion)
		if err != nil {
			fmt.Println("Error notification to slack:", err)
		}
		err = notification.SendNotificationToDiscord(interactor.language, newVersion)
		if err != nil {
			fmt.Println("Error notification to discord:", err)
		}
	}
}
