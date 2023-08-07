package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityKategoriPekerjaan interface {
	EntityCreate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
	EntityResults() (*[]models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError)
}
