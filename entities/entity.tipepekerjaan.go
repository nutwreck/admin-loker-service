package entities

import (
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type EntityTipePekerjaan interface {
	EntityCreate(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError)
	EntityResult(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError)
	EntityResults(input *schemes.SchemeTipePekerjaan) (*[]models.ModelTipePekerjaan, int64, schemes.SchemeDatabaseError)
	EntityDelete(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError)
	EntityUpdate(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError)
}
