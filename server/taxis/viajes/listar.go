package viajes

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func ListarByRadio(db *sql.DB, lat, lon float64, radio int) ([]*model.ViajesResponse, error) {
	sql := `
	SELECT v.id, 
       v.pasajero_id, 
       v.conductor_id, 
       v.estado, 
       v.descripcion, 
       ST_X(v.origen) AS origen_x, 
       ST_Y(v.origen) AS origen_y, 
       ST_X(v.destino) AS destino_x, 
       ST_Y(v.destino) AS destino_y, 
       v.categoria_id, 
       v.registrado,
       u1.username AS pasajero_username, 
       u2.username AS conductor_username
	FROM viajes v
	INNER JOIN usuarios u1 ON v.pasajero_id = u1.id
	LEFT JOIN usuarios u2 ON v.conductor_id = u2.id
	where ST_Distance_Sphere(Point(?, ?), v.origen) <= ?
	`
	rows, err := db.Query(sql, lat, lon, radio)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.ViajesResponse{}
	for rows.Next() {
		v := model.ViajesResponse{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func GetById(db *sql.DB, id string) (*model.ViajesResponse, error) {
	sql := `
	SELECT v.id, 
       v.pasajero_id, 
       v.conductor_id, 
       v.estado, 
       v.descripcion, 
       ST_X(v.origen) AS origen_x, 
       ST_Y(v.origen) AS origen_y, 
       ST_X(v.destino) AS destino_x, 
       ST_Y(v.destino) AS destino_y, 
       v.categoria_id, 
       v.registrado,
       u1.username AS pasajero_username, 
       u2.username AS conductor_username
	FROM viajes v
	INNER JOIN usuarios u1 ON v.pasajero_id = u1.id
	LEFT JOIN usuarios u2 ON v.conductor_id = u2.id
	where v.id=?
	`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("no se encontraron el viaje")
	}

	v := model.ViajesResponse{}
	er := parse(row, &v)
	if er != nil {
		return nil, er
	}
	return &v, nil
}

func ListarByUsuario(db *sql.DB, usuario_id string) ([]*model.ViajesResponse, error) {
	sql := `
	SELECT v.id, 
       v.pasajero_id, 
       v.conductor_id, 
       v.estado, 
       v.descripcion, 
       ST_X(v.origen) AS origen_x, 
       ST_Y(v.origen) AS origen_y, 
       ST_X(v.destino) AS destino_x, 
       ST_Y(v.destino) AS destino_y, 
       v.categoria_id, 
       v.registrado,
       u1.username AS pasajero_username, 
       u2.username AS conductor_username
	FROM viajes v
	INNER JOIN usuarios u1 ON v.pasajero_id = u1.id
	LEFT JOIN usuarios u2 ON v.conductor_id = u2.id
	where v.pasajero_id=? or v.conductor_id=?
	`
	rows, err := db.Query(sql, usuario_id, usuario_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.ViajesResponse{}
	for rows.Next() {
		v := model.ViajesResponse{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByPasajero(db *sql.DB, usuario_id string) ([]*model.ViajesResponse, error) {
	sql := `
	SELECT v.id, 
       v.pasajero_id, 
       v.conductor_id, 
       v.estado, 
       v.descripcion, 
       ST_X(v.origen) AS origen_x, 
       ST_Y(v.origen) AS origen_y, 
       ST_X(v.destino) AS destino_x, 
       ST_Y(v.destino) AS destino_y, 
       v.categoria_id, 
       v.registrado,
       u1.username AS pasajero_username, 
       u2.username AS conductor_username
	FROM viajes v
	INNER JOIN usuarios u1 ON v.pasajero_id = u1.id
	LEFT JOIN usuarios u2 ON v.conductor_id = u2.id
	where v.pasajero_id=?
	`
	rows, err := db.Query(sql, usuario_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.ViajesResponse{}
	for rows.Next() {
		v := model.ViajesResponse{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByConductor(db *sql.DB, usuario_id string) ([]*model.ViajesResponse, error) {
	sql := `
	SELECT v.id, 
       v.pasajero_id, 
       v.conductor_id, 
       v.estado, 
       v.descripcion, 
       ST_X(v.origen) AS origen_x, 
       ST_Y(v.origen) AS origen_y, 
       ST_X(v.destino) AS destino_x, 
       ST_Y(v.destino) AS destino_y, 
       v.categoria_id, 
       v.registrado,
       u1.username AS pasajero_username, 
       u2.username AS conductor_username
	FROM viajes v
	INNER JOIN usuarios u1 ON v.pasajero_id = u1.id
	LEFT JOIN usuarios u2 ON v.conductor_id = u2.id
	where v.conductor_id=?
	`
	rows, err := db.Query(sql, usuario_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.ViajesResponse{}
	for rows.Next() {
		v := model.ViajesResponse{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}

func ListarByCategoria(db *sql.DB, categoria_id string) ([]*model.ViajesResponse, error) {
	sql := `
	SELECT v.id, 
       v.pasajero_id, 
       v.conductor_id, 
       v.estado, 
       v.descripcion, 
       ST_X(v.origen) AS origen_x, 
       ST_Y(v.origen) AS origen_y, 
       ST_X(v.destino) AS destino_x, 
       ST_Y(v.destino) AS destino_y, 
       v.categoria_id, 
       v.registrado,
       u1.username AS pasajero_username, 
       u2.username AS conductor_username
	FROM viajes v
	INNER JOIN usuarios u1 ON v.pasajero_id = u1.id
	LEFT JOIN usuarios u2 ON v.conductor_id = u2.id
	where v.categoria_id=?
	`
	rows, err := db.Query(sql, categoria_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := []*model.ViajesResponse{}
	for rows.Next() {
		v := model.ViajesResponse{}
		er := parse2(rows, &v)
		if er != nil {
			return nil, er
		}
		vs = append(vs, &v)
	}
	return vs, nil
}
