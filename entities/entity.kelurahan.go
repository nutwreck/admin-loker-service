package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityKelurahan interface {
	EntityCreate(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeKelurahan) (*[]schemes.SchemeGetDataKelurahan, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError)
}
