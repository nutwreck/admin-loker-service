package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityKeahlian interface {
	EntityCreate(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeKeahlian) (*[]models.ModelKeahlian, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError)
}
