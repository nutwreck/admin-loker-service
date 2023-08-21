package schemes

type SchemeKelurahan struct {
	CodeKelurahan       string `json:"code_kelurahan" validate:"required"`
	ParentCodeKecamatan string `json:"parent_code_kecamatan" validate:"required"`
	// Input with Uppercase
	Name    string `json:"name" validate:"required,uppercase"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
}

type SchemeKelurahanRequest struct {
	CodeKelurahan       string `json:"code_kelurahan" validate:"required" example:"0000001"`
	ParentCodeKecamatan string `json:"parent_code_kecamatan" validate:"required" example:"000001"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"PANJANG"`
}

type SchemeKelurahanRequestUpdate struct {
	ParentCodeKecamatan string `json:"parent_code_kecamatan" validate:"required" example:"000001"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"PANJANG"`
}
