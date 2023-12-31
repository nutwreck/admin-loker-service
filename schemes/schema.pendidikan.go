package schemes

type SchemePendidikan struct {
	ID string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=75" example:"SMA"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type SchemePendidikanRequest struct {
	// Input with Lowercase
	Name   string `json:"name" validate:"required,lowercase,max=75" example:"SMA"`
	Active *bool  `json:"active" validate:"boolean" example:"true"`
}
