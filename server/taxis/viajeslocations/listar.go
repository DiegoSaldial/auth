package viajeslocations

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func GetById(db *sql.DB, id string) (*model.ViajesLocations, error) {
	sql := `select id,ST_X(origen) AS lat, ST_Y(origen) AS lon,viaje_id,registrado where id=?`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("no existe el registro")
	}
	vl := model.ViajesLocations{}
	er := parse(row, &vl)
	if er != nil {
		return nil, er
	}
	return &vl, nil
}
