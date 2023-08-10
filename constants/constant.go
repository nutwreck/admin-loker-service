package constants

// GENERAL
var (
	TRUE_VALUE  = true
	FALSE_VALUE = false

	EMPTY_VALUE = ""
)

// JENIS KELAMIN
type JenisKelamin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var JenisKelamins = []JenisKelamin{
	{ID: 1, Name: "laki-laki"},
	{ID: 2, Name: "perempuan"},
}

// STATUS PERNIKAHAN
type StatusPernikahan struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var StatusPernikahans = []StatusPernikahan{
	{ID: 1, Name: "belum menikah"},
	{ID: 2, Name: "menikah"},
}
