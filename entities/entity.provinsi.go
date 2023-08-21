package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityProvinsi interface {
	EntityCreate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeProvinsi) (*[]models.ModelProvinsi, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError)
}
