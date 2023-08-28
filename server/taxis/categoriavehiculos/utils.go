package categoriavehiculos

import (
	"database/sql"
	"inventarios/graph/model"
)

func parse(row *sql.Row, t *model.CategoriaVehiculos) error {
	return row.Scan(
		&t.ID,
		&t.Nombre,
		&t.Descripcion,
	)
}

func parse2(rows *sql.Rows, t *model.CategoriaVehiculos) error {
	return rows.Scan(
		&t.ID,
		&t.Nombre,
		&t.Descripcion,
	)
}
