package schemes

type SchemeKecamatan struct {
	CodeKecamatan       string `json:"code_kecamatan" validate:"required"`
	ParentCodeKabupaten string `json:"parent_code_kabupaten" validate:"required"`
	// Input with Uppercase
	Name          string `json:"name" validate:"required,uppercase"`
	Page          int    `json:"page"`
	PerPage       int    `json:"perpage"`
	Sort          string `json:"sort"`
	Search        string `json:"search"`
	CodeNegara    string `json:"code_negara"`
	NameNegara    string `json:"name_negara"`
	CodeProvinsi  string `json:"code_provinsi"`
	NameProvinsi  string `json:"name_provinsi"`
	NameKabupaten string `json:"name_kabupaten"`
}

type SchemeGetDataKecamatan struct {
	CodeNegara    string `json:"code_negara"`
	NameNegara    string `json:"name_negara"`
	CodeProvinsi  string `json:"code_provinsi"`
	NameProvinsi  string `json:"name_provinsi"`
	CodeKabupaten string `json:"code_kabupaten"`
	NameKabupaten string `json:"name_kabupaten"`
	CodeKecamatan string `json:"code_kecamatan"`
	NameKecamatan string `json:"name_kecamatan"`
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
