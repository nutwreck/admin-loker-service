package schemes

type SchemeKecamatan struct {
	CodeKecamatan       string `json:"code_kecamatan" validate:"required"`
	ParentCodeKabupaten string `json:"parent_code_kabupaten" validate:"required"`
	// Input with Uppercase
	Name    string `json:"name" validate:"required,uppercase"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
}

type SchemeKecamatanRequest struct {
	CodeKecamatan       string `json:"code_kecamatan" validate:"required" example:"000001"`
	ParentCodeKabupaten string `json:"parent_code_kabupaten" validate:"required" example:"00001"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"AMBARAWA"`
}

type SchemeKecamatanRequestUpdate struct {
	ParentCodeKabupaten string `json:"parent_code_kabupaten" validate:"required" example:"00001"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"AMBARAWA"`
}
