package vehiculos

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func Listar(db *sql.DB) ([]*model.Vehiculos, error) {
	sql := `select id,placa,puertas,capacidad,descripcion,color,modelo,anio,categoria_id,foto_url,estado,registrado from vehiculos`
	res, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	vs := []*model.Vehiculos{}
	for res.Next() {
		v := model.Vehiculos{}
		er := parse(res, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func GetById(db *sql.DB, id string) (*model.Vehiculos, error) {
	sql := `select id,placa,puertas,capacidad,descripcion,color,modelo,anio,categoria_id,foto_url,estado,registrado from vehiculos where id=?`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("no existe vehiculo")
	}
	v := model.Vehiculos{}
	er := parse2(row, &v)
	if er != nil {
		return nil, er
	}
	return &v, nil
}

func ListarByCategoria(db *sql.DB, cate string) ([]*model.Vehiculos, error) {
	sql := `select id,placa,puertas,capacidad,descripcion,color,modelo,anio,categoria_id,foto_url,estado,registrado from vehiculos where categoria_id=?`
	res, err := db.Query(sql, cate)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	vs := []*model.Vehiculos{}
	for res.Next() {
		v := model.Vehiculos{}
		er := parse(res, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByConductor(db *sql.DB, userid string) ([]*model.Vehiculos, error) {
	sql := `
	select id,placa,puertas,capacidad,descripcion,color,modelo,anio,categoria_id,foto_url,estado,registrado 
	from vehiculos 
	where id in(select vehiculo_id from conductor_vehiculos where usuario_id=?)`
	res, err := db.Query(sql, userid)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	vs := []*model.Vehiculos{}
	for res.Next() {
		v := model.Vehiculos{}
		er := parse(res, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}
