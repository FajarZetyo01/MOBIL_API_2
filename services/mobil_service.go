package services

import (
	"MOBIL_API_2/model/web"
	"context"
)

type MobilService interface {
	Create(ctx context.Context, request web.MobilCreateRequest) web.MobilResponse
	Update(ctx context.Context, request web.MobilUpdateRequest) web.MobilResponse
	Delete(ctx context.Context, mobilId int)
	FindById(ctx context.Context, mobilId int) web.MobilResponse
	FindAll(ctx context.Context) []web.MobilResponse
}
