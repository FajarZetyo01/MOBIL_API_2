package repositories

import (
	"MOBIL_API_2/helpers"
	"MOBIL_API_2/model/domain"
	"context"
	"database/sql"
	"errors"
)

type MobilRepositoryImpl struct {
}

func NewMobilRepository() *MobilRepositoryImpl {
	return &MobilRepositoryImpl{}
}

func (repository *MobilRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, mobil domain.Mobil) domain.Mobil {
	SQL := "INSERT INTO mobil(nama_mobil,tahun,merek,warna) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, mobil.NamaMobil, mobil.Tahun, mobil.Merek, mobil.Warna)
	helpers.PanicIfError(err)
	id, err := result.LastInsertId()
	helpers.PanicIfError(err)
	mobil.Id = int(id)
	return mobil
}

func (repository MobilRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, mobil domain.Mobil) domain.Mobil {
	SQL := "UPDATE mobil SET nama_mobil = ?, tahun = ?, merek = ?, warna = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, mobil.NamaMobil, mobil.Tahun, mobil.Merek, mobil.Warna, mobil.Id)
	helpers.PanicIfError(err)
	return mobil
}

func (repository MobilRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, mobil domain.Mobil) {
	SQL := "DELETE FROM mobil WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, mobil.Id)
	helpers.PanicIfError(err)
}

func (repository MobilRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, mobilId int) (domain.Mobil, error) {
	SQL := "SELECT id, nama_mobil,tahun,merek,warna FROM mobil WHERE id =?"
	rows, err := tx.QueryContext(ctx, SQL, mobilId)
	helpers.PanicIfError(err)
	defer rows.Close()

	mobil := domain.Mobil{}
	if rows.Next() {
		err := rows.Scan(&mobil.Id, &mobil.NamaMobil, &mobil.Tahun, &mobil.Merek, &mobil.Warna)
		helpers.PanicIfError(err)
		return mobil, nil
	} else {
		return mobil, errors.New("Mobil is not found")
	}
}

func (repository MobilRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Mobil {
	SQL := "SELECT id,nama_mobil,tahun,merek,warna FROM mobil"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var mobils []domain.Mobil
	for rows.Next() {
		mobil := domain.Mobil{}
		err := rows.Scan(&mobil.Id, &mobil.NamaMobil, &mobil.Tahun, &mobil.Merek, &mobil.Warna)
		helpers.PanicIfError(err)
		mobils = append(mobils, mobil)
	}
	return mobils
}
