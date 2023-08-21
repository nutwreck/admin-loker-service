package schemes

type SchemeProvinsi struct {
	CodeProvinsi     string `json:"code_provinsi" validate:"required"`
	ParentCodeNegara string `json:"parent_code_negara" validate:"required"`
	// Input with Uppercase
	Name    string `json:"name" validate:"required,uppercase"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
}

type SchemeProvinsiRequest struct {
	CodeProvinsi     string `json:"code_provinsi" validate:"required" example:"001"`
	ParentCodeNegara string `json:"parent_code_negara" validate:"required" example:"1"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"JAWA TENGAH"`
}

type SchemeProvinsiRequestUpdate struct {
	ParentCodeNegara string `json:"parent_code_negara" validate:"required" example:"1"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"JAWA TENGAH"`
}
