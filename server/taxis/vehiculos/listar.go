package vehiculos

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func Listar(db *sql.DB) ([]*model.VehiculosResponse, error) {
	sql := getSqlListar()
	res, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	vs := []*model.VehiculosResponse{}
	for res.Next() {
		v := model.VehiculosResponse{}
		er := parse(res, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func GetById(db *sql.DB, id string) (*model.VehiculosResponse, error) {
	sql := getSqlByID()
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("no existe vehiculo")
	}
	v := model.VehiculosResponse{}
	er := parse2(row, &v)
	if er != nil {
		return nil, er
	}
	return &v, nil
}

func ListarByCategoria(db *sql.DB, cate string) ([]*model.VehiculosResponse, error) {
	sql := getSqlByCategoria()
	res, err := db.Query(sql, cate)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	vs := []*model.VehiculosResponse{}
	for res.Next() {
		v := model.VehiculosResponse{}
		er := parse(res, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByConductor(db *sql.DB, userid string) ([]*model.VehiculosResponse, error) {
	sql := getSqlByConductor()
	res, err := db.Query(sql, userid)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	vs := []*model.VehiculosResponse{}
	for res.Next() {
		v := model.VehiculosResponse{}
		er := parse(res, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}
