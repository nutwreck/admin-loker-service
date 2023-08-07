package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityTahunPengalaman interface {
	EntityCreate(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelTahunPengalaman, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError)
}
