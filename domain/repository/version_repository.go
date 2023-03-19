package repository

import "github.com/yourusername/yourproject/domain/entity"

type VersionRepository interface {
	GetVersion() (*entity.Version, error)
	SaveVersion(version *entity.Version) error
}
