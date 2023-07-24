package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityUser interface {
	EntityRegister(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
	EntityLogin(input *schemes.SchemeUser) (*models.ModelUser, schemes.SchemeDatabaseError)
}
