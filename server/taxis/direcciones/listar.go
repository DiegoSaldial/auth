package direcciones

import (
	"database/sql"
	"errors"
	"inventarios/graph/model"
)

func GetById(db *sql.DB, id string) (*model.Direcciones, error) {
	// sql := `select id, nombre,descripcion,ubicacion,usuarios_id,registrado from direcciones where id=?`
	sql := `select id,nombre,descripcion,ST_X(ubicacion) AS lat, ST_Y(ubicacion) AS lon,usuario_id,registrado from direcciones where id=?`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("No se encontro el registro con id")
	}
	dir := model.Direcciones{}
	er := parse(row, &dir)
	if er != nil {
		return nil, er
	}
	return &dir, nil
}

func GetByUsuario(db *sql.DB, id string) ([]*model.Direcciones, error) {
	// sql := `select id, nombre,descripcion,ubicacion,usuarios_id,registrado from direcciones where id=?`
	sql := `select id,nombre,descripcion,ST_X(ubicacion) AS lat, ST_Y(ubicacion) AS lon,usuario_id,registrado from direcciones where usuario_id=?`
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dirs := []*model.Direcciones{}
	for rows.Next() {
		dir := model.Direcciones{}
		er := parse2(rows, &dir)
		if er != nil {
			return nil, er
		}
		dirs = append(dirs, &dir)
	}

	return dirs, nil
}
