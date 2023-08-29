package categoriavehiculos

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func GetById(db *sql.DB, id string) (*model.CategoriaVehiculosResponse, error) {
	sql := `
	SELECT
		cv.id,
		cv.nombre,
		cv.descripcion,
		COUNT(v.id) AS cantidad_vehiculos
	FROM
		categoria_vehiculos cv
	LEFT JOIN
		vehiculos v ON cv.id = v.categoria_id
	where cv.id = ?
	GROUP BY
		cv.id, cv.nombre, cv.descripcion;
	`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("No existe la categoria vehiculo por id")
	}
	cate := model.CategoriaVehiculosResponse{}
	er := parse(row, &cate)
	if er != nil {
		return nil, er
	}
	return &cate, nil
}

func Listar(db *sql.DB) ([]*model.CategoriaVehiculosResponse, error) {
	// sql := `select id,nombre,descripcion from categoria_vehiculos`
	sql := `
	SELECT
		cv.id,
		cv.nombre,
		cv.descripcion,
		COUNT(v.id) AS cantidad_vehiculos
	FROM
		categoria_vehiculos cv
	LEFT JOIN
		vehiculos v ON cv.id = v.categoria_id
	GROUP BY
		cv.id, cv.nombre, cv.descripcion
	ORDER BY cv.id desc
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cates := []*model.CategoriaVehiculosResponse{}
	for rows.Next() {

		cate := model.CategoriaVehiculosResponse{}
		er := parse2(rows, &cate)
		if er != nil {
			return nil, er
		}
		cates = append(cates, &cate)
	}
	return cates, nil
}
