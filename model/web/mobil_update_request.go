package web

type MobilUpdateRequest struct {
	Id        int    `json:"id"`
	NamaMobil string `json:"nama_mobil"`
	Tahun     int    `json:"tahun"`
	Merek     string `json:"merek"`
	Warna     string `json:"warna"`
}
