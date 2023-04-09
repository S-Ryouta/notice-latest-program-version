package repository

import (
	"github.com/S-Ryouta/notice-latest-program-version/domain/entity"
)

type VersionRepository interface {
	GetVersion() (*entity.Version, error)
	SaveVersion(version *entity.Version) error
}
