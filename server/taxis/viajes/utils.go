package viajes

import (
	"database/sql"
	"inventarios/graph/model"
)

func parse(row *sql.Row, t *model.Viajes) error {
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
	)
}

func parse2(rows *sql.Rows, t *model.Viajes) error {
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
	)
}
