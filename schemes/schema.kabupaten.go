package schemes

type SchemeKabupaten struct {
	CodeKabupaten      string `json:"code_kabupaten" validate:"required"`
	ParentCodeProvinsi string `json:"parent_code_provinsi" validate:"required"`
	// Input with Uppercase
	Name         string `json:"name" validate:"required,uppercase"`
	Page         int    `json:"page"`
	PerPage      int    `json:"perpage"`
	Sort         string `json:"sort"`
	Search       string `json:"search"`
	CodeNegara   string `json:"code_negara"`
	NameNegara   string `json:"name_negara"`
	NameProvinsi string `json:"name_provinsi"`
}

type SchemeGetDataKabupaten struct {
	CodeNegara    string `json:"code_negara"`
	NameNegara    string `json:"name_negara"`
	CodeProvinsi  string `json:"code_provinsi"`
	NameProvinsi  string `json:"name_provinsi"`
	CodeKabupaten string `json:"code_kabupaten"`
	NameKabupaten string `json:"name_kabupaten"`
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
