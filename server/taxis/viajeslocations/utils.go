package viajeslocations

import (
	"database/sql"
	"taxis/graph/model"
)

func parse(row *sql.Row, t *model.ViajesLocations) error {
	return row.Scan(
		&t.ID,
		&t.Latitud,
		&t.Longitud,
		&t.ViajeID,
		&t.Registrado,
	)
}
