package schemes

type SchemeKelurahan struct {
	CodeKelurahan       string `json:"code_kelurahan" validate:"required"`
	ParentCodeKecamatan string `json:"parent_code_kecamatan" validate:"required"`
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
	CodeKabupaten string `json:"code_kabupaten"`
	NameKabupaten string `json:"name_kabupaten"`
	NameKecamatan string `json:"name_kecamatan"`
}

type SchemeGetDataKelurahan struct {
	CodeNegara    string `json:"code_negara"`
	NameNegara    string `json:"name_negara"`
	CodeProvinsi  string `json:"code_provinsi"`
	NameProvinsi  string `json:"name_provinsi"`
	CodeKabupaten string `json:"code_kabupaten"`
	NameKabupaten string `json:"name_kabupaten"`
	CodeKecamatan string `json:"code_kecamatan"`
	NameKecamatan string `json:"name_kecamatan"`
	CodeKelurahan string `json:"code_kelurahan"`
	NameKelurahan string `json:"name_kelurahan"`
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
