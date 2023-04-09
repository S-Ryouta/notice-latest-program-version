package repository

import (
	"github.com/S-Ryouta/notice-latest-program-version/domain/entity"
)

type VersionRepository interface {
	GetVersion(language string) (*entity.Version, error)
	SaveVersion(version *entity.Version) error
}
