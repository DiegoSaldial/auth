package vehiculos

import (
	"database/sql"
	"taxis/graph/model"
)

func parse(rows *sql.Rows, t *model.Vehiculos) error {
	return rows.Scan(
		&t.ID,
		&t.Placa,
		&t.Puertas,
		&t.Capacidad,
		&t.Descripcion,
		&t.Color,
		&t.Modelo,
		&t.Anio,
		&t.CategoriaID,
		&t.FotoURL,
		&t.Estado,
		&t.Registrado,
	)
}

func parse2(row *sql.Row, t *model.Vehiculos) error {
	return row.Scan(
		&t.ID,
		&t.Placa,
		&t.Puertas,
		&t.Capacidad,
		&t.Descripcion,
		&t.Color,
		&t.Modelo,
		&t.Anio,
		&t.CategoriaID,
		&t.FotoURL,
		&t.Estado,
		&t.Registrado,
	)
}
