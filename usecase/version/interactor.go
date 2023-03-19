package version

import (
	"context"
	"fmt"
	"github.com/yourusername/yourproject/domain/entity"
	"github.com/yourusername/yourproject/domain/repository"
	"github.com/yourusername/yourproject/infrastructure/notification"
	"github.com/yourusername/yourproject/infrastructure/version"
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
	ctx := context.Background()

	newVersion, err := version.GetLatestGolangVersion()
	if err != nil {
		fmt.Println("Error getting the latest Golang version:", err)
		return
	}

	storedVersionEntity, err := interactor.VersionRepo.GetVersion(ctx)
	if err != nil {
		fmt.Println("Error getting the stored Golang version:", err)
		return
	}

	if storedVersionEntity.Version != newVersion {
		err = interactor.VersionRepo.SaveVersion(ctx, &entity.Version{
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
