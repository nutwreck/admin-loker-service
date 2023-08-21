package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityNegara interface {
	EntityCreate(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeNegara) (*[]models.ModelNegara, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError)
}
