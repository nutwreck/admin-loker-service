package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityKabupaten interface {
	EntityCreate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeKabupaten) (*[]schemes.SchemeGetDataKabupaten, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError)
}
