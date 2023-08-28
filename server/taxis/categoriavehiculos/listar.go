package categoriavehiculos

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func GetById(db *sql.DB, id string) (*model.CategoriaVehiculos, error) {
	sql := `select id,nombre,descripcion from categoria_vehiculos where id = ?`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("No existe la categoria vehiculo por id")
	}
	cate := model.CategoriaVehiculos{}
	er := parse(row, &cate)
	if er != nil {
		return nil, er
	}
	return &cate, nil
}

func Listar(db *sql.DB) ([]*model.CategoriaVehiculos, error) {
	sql := `select id,nombre,descripcion from categoria_vehiculos`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cates := []*model.CategoriaVehiculos{}
	for rows.Next() {

		cate := model.CategoriaVehiculos{}
		er := parse2(rows, &cate)
		if er != nil {
			return nil, er
		}
		cates = append(cates, &cate)
	}
	return cates, nil
}
