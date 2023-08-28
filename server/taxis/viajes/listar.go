package viajes

import (
	"database/sql"
	"errors"
	"inventarios/graph/model"
)

func ListarByRadio(db *sql.DB, lat, lon float64, radio int) ([]*model.Viajes, error) {
	sql := `
	select id,pasajero_id,conductor_id,estado,descripcion,ST_X(origen), ST_Y(origen),ST_X(destino), ST_Y(destino),categoria_id,registrado
	from viajes
	where ST_Distance_Sphere(Point(?, ?), origen) <= ?
	`
	rows, err := db.Query(sql, lat, lon, radio)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.Viajes{}
	for rows.Next() {
		v := model.Viajes{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func GetById(db *sql.DB, id string) (*model.Viajes, error) {
	sql := `
	select id,pasajero_id,conductor_id,estado,descripcion,ST_X(origen), ST_Y(origen),ST_X(destino), ST_Y(destino),categoria_id,registrado
	from viajes
	where id=?
	`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("no se encontraron el viaje")
	}

	v := model.Viajes{}
	er := parse(row, &v)
	if er != nil {
		return nil, er
	}
	return &v, nil
}

func ListarByUsuario(db *sql.DB, usuario_id string) ([]*model.Viajes, error) {
	sql := `
	select id,pasajero_id,conductor_id,estado,descripcion,ST_X(origen), ST_Y(origen),ST_X(destino), ST_Y(destino),categoria_id,registrado
	from viajes
	where pasajero_id=? or conductor_id=?
	`
	rows, err := db.Query(sql, usuario_id, usuario_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.Viajes{}
	for rows.Next() {
		v := model.Viajes{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByPasajero(db *sql.DB, usuario_id string) ([]*model.Viajes, error) {
	sql := `
	select id,pasajero_id,conductor_id,estado,descripcion,ST_X(origen), ST_Y(origen),ST_X(destino), ST_Y(destino),categoria_id,registrado
	from viajes
	where pasajero_id=?
	`
	rows, err := db.Query(sql, usuario_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.Viajes{}
	for rows.Next() {
		v := model.Viajes{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByConductor(db *sql.DB, usuario_id string) ([]*model.Viajes, error) {
	sql := `
	select id,pasajero_id,conductor_id,estado,descripcion,ST_X(origen), ST_Y(origen),ST_X(destino), ST_Y(destino),categoria_id,registrado
	from viajes
	where conductor_id=?
	`
	rows, err := db.Query(sql, usuario_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.Viajes{}
	for rows.Next() {
		v := model.Viajes{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByCategoria(db *sql.DB, categoria_id string) ([]*model.Viajes, error) {
	sql := `
	select id,pasajero_id,conductor_id,estado,descripcion,ST_X(origen), ST_Y(origen),ST_X(destino), ST_Y(destino),categoria_id,registrado
	from viajes
	where categoria_id=?
	`
	rows, err := db.Query(sql, categoria_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.Viajes{}
	for rows.Next() {
		v := model.Viajes{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}
