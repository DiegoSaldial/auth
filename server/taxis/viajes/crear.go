package viajes

import (
	"database/sql"
	"strconv"
	"taxis/graph/model"
)

func Crear(db *sql.DB, input model.CreateViajes) (*model.ViajesResponse, error) {
	sql := `
	insert into viajes(pasajero_id,descripcion,origen,destino,categoria_id)
	values(?,?,POINT(?,?),POINT(?,?),?)
	`
	res, err := db.Exec(sql, input.PasajeroID, input.Descripcion, input.OrigenLat, input.OrigenLon, input.DestinoLat, input.DestinoLon, input.CategoriaID)
	if err != nil {
		return nil, err
	}

	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)

	return GetById(db, id)
}