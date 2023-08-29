package categoriavehiculos

import (
	"database/sql"
	"taxis/graph/model"
)

func parse(row *sql.Row, t *model.CategoriaVehiculosResponse) error {
	return row.Scan(
		&t.ID,
		&t.Nombre,
		&t.Descripcion,
		&t.Vehiculos,
	)
}

func parse2(rows *sql.Rows, t *model.CategoriaVehiculosResponse) error {
	return rows.Scan(
		&t.ID,
		&t.Nombre,
		&t.Descripcion,
		&t.Vehiculos,
	)
}
