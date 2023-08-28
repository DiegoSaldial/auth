package direcciones

import (
	"database/sql"
	"inventarios/graph/model"
)

func parse(row *sql.Row, t *model.Direcciones) error {
	return row.Scan(
		&t.ID,
		&t.Nombre,
		&t.Descripcion,
		&t.Latitud,
		&t.Longitud,
		&t.UsuarioID,
		&t.Registrado,
	)
}

func parse2(rows *sql.Rows, t *model.Direcciones) error {
	return rows.Scan(
		&t.ID,
		&t.Nombre,
		&t.Descripcion,
		&t.Latitud,
		&t.Longitud,
		&t.UsuarioID,
		&t.Registrado,
	)
}
