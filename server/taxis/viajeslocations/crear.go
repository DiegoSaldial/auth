package viajeslocations

import (
	"database/sql"
	"inventarios/graph/model"
	"strconv"
)

func SetUbicacion(db *sql.DB, input model.CreateViajesLocations) (*model.ViajesLocations, error) {
	sql := `insert into viajes_locations(origen,viaje_id) values(POINT(?,?),?)`
	res, err := db.Exec(sql, input.Latitud, input.Longitud, input.ViajeID)
	if err != nil {
		return nil, err
	}
	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)
	return GetById(db, id)
}
