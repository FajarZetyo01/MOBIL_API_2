package repositories

import (
	"MOBIL_API_2/model/domain"
	"context"
	"database/sql"
)

type MobilRepository interface {
	Save(ctx context.Context, tx *sql.Tx, mobil domain.Mobil) domain.Mobil
	Update(ctx context.Context, tx *sql.Tx, mobil domain.Mobil) domain.Mobil
	Delete(ctx context.Context, tx *sql.Tx, mobil domain.Mobil)
	FindById(ctx context.Context, tx *sql.Tx, mobilId int) (domain.Mobil, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Mobil
}
