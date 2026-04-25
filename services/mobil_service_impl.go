package services

import (
	"MOBIL_API_2/exception"
	"MOBIL_API_2/helpers"
	"MOBIL_API_2/model/domain"
	"MOBIL_API_2/model/web"
	"MOBIL_API_2/repositories"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type MobilServiceImpl struct {
	MobilRepository repositories.MobilRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewMobilService(mobilRepository repositories.MobilRepository, DB *sql.DB, validate *validator.Validate) MobilService {
	return &MobilServiceImpl{
		MobilRepository: mobilRepository,
		DB:              DB,
		Validate:        validate}
}

func (service *MobilServiceImpl) Create(ctx context.Context, request web.MobilCreateRequest) web.MobilResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	var mobil = domain.Mobil{
		NamaMobil: request.NamaMobil,
		Tahun:     request.Tahun,
		Merek:     request.Merek,
		Warna:     request.Warna,
	}
	mobil = service.MobilRepository.Save(ctx, tx, mobil)
	return helpers.ToMobilresponse(mobil)
}

func (service *MobilServiceImpl) Update(ctx context.Context, request web.MobilUpdateRequest) web.MobilResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	mobil, err := service.MobilRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	mobil.NamaMobil = request.NamaMobil
	mobil.Tahun = request.Tahun
	mobil.Merek = request.Merek
	mobil.Warna = request.Warna
	mobil = service.MobilRepository.Update(ctx, tx, mobil)
	return helpers.ToMobilresponse(mobil)
}

func (service *MobilServiceImpl) Delete(ctx context.Context, mobilId int) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	mobil, err := service.MobilRepository.FindById(ctx, tx, mobilId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.MobilRepository.Delete(ctx, tx, mobil)
}

func (service *MobilServiceImpl) FindById(ctx context.Context, mobilId int) web.MobilResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	mobil, err := service.MobilRepository.FindById(ctx, tx, mobilId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helpers.ToMobilresponse(mobil)
}

func (service *MobilServiceImpl) FindAll(ctx context.Context) []web.MobilResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	mobils := service.MobilRepository.FindAll(ctx, tx)
	return helpers.ToMobilResponses(mobils)
}
