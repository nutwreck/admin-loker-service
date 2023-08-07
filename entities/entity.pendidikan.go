package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityPendidikan interface {
	EntityCreate(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelPendidikan, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError)
}
