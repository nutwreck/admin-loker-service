package schemes

type SchemeKabupaten struct {
	CodeKabupaten      string `json:"code_kabupaten" validate:"required"`
	ParentCodeProvinsi string `json:"parent_code_provinsi" validate:"required"`
	// Input with Uppercase
	Name    string `json:"name" validate:"required,uppercase"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
}

type SchemeKabupatenRequest struct {
	CodeKabupaten      string `json:"code_kabupaten" validate:"required" example:"00001"`
	ParentCodeProvinsi string `json:"parent_code_provinsi" validate:"required" example:"001"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"KAB. SEMARANG"`
}

type SchemeKabupatenRequestUpdate struct {
	ParentCodeProvinsi string `json:"parent_code_provinsi" validate:"required" example:"001"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"KAB. SEMARANG"`
}
