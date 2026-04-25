package web

type MobilCreateRequest struct {
	NamaMobil string `validate:"required,max=100,min=1" json:"nama_mobil"`
	Tahun     int    `validate:"required,max=100000,min=2" json:"tahun"`
	Merek     string `validate:"required,max=100,min=1" json:"merek"`
	Warna     string `validate:"required,max=100,min=1" json:"warna"`
}
