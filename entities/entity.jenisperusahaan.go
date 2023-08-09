package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityJenisPerusahaan interface {
	EntityCreate(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeJenisPerusahaan) (*[]models.ModelJenisPerusahaan, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError)
}
