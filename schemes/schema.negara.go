package schemes

type SchemeNegara struct {
	CodeNegara string `json:"code_negara" validate:"required"`
	ParentCode string `json:"parent_code" validate:"required"`
	// Input with Uppercase
	Name    string `json:"name" validate:"required,uppercase"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	SortBy  string `json:"sortby"`
	OrderBy string `json:"orderby"`
}

type SchemeNegaraRequest struct {
	CodeNegara string `json:"code_negara" validate:"required" example:"1"`
	ParentCode string `json:"parent_code" validate:"required" example:"0"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"INDONESIA"`
}

type SchemeNegaraRequestUpdate struct {
	ParentCode string `json:"parent_code" validate:"required" example:"0"`
	// Input with Uppercase
	Name string `json:"name" validate:"required,uppercase" example:"INDONESIA"`
}
