package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityKecamatan interface {
	EntityCreate(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeKecamatan) (*[]schemes.SchemeGetDataKecamatan, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError)
}
