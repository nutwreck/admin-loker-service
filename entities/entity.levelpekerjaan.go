package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityLevelPekerjaan interface {
	EntityCreate(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeLevelPekerjaan) (*[]models.ModelLevelPekerjaan, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError)
}
