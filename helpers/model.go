package helpers

import (
	"MOBIL_API_2/model/domain"
	"MOBIL_API_2/model/web"
)

func ToMobilresponse(mobil domain.Mobil) web.MobilResponse {
	return web.MobilResponse{
		Id:        mobil.Id,
		NamaMobil: mobil.NamaMobil,
		Tahun:     mobil.Tahun,
		Merek:     mobil.Merek,
		Warna:     mobil.Warna,
	}
}

func ToMobilResponses(mobils []domain.Mobil) []web.MobilResponse {
	var mobilResponses []web.MobilResponse
	for _, mobil := range mobils {
		mobilResponses = append(mobilResponses, ToMobilresponse(mobil))
	}
	return mobilResponses
}
