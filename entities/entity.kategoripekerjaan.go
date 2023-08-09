package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityKategoriPekerjaan interface {
	EntityCreate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeKategoriPekerjaan) (*[]models.ModelKategoriPekerjaan, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
}
