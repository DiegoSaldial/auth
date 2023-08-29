package viajes

import (
	"database/sql"
	"taxis/graph/model"
)

func parse(row *sql.Row, t *model.ViajesResponse) error {
	return row.Scan(
		&t.ID,
		&t.PasajeroID,
		&t.ConductorID,
		&t.Estado,
		&t.Descripcion,
		&t.OrigenLat,
		&t.OrigenLon,
		&t.DestinoLat,
		&t.DestinoLon,
		&t.CategoriaID,
		&t.Registrado,
		&t.PasajeroUsername,
		&t.ConductorUsername,
	)
}

func parse2(rows *sql.Rows, t *model.ViajesResponse) error {
	return rows.Scan(
		&t.ID,
		&t.PasajeroID,
		&t.ConductorID,
		&t.Estado,
		&t.Descripcion,
		&t.OrigenLat,
		&t.OrigenLon,
		&t.DestinoLat,
		&t.DestinoLon,
		&t.CategoriaID,
		&t.Registrado,
		&t.PasajeroUsername,
		&t.ConductorUsername,
	)
}
