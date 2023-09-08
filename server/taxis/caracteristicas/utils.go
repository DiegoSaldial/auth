package caracteristicas

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func parse(row *sql.Row, t *model.CaracteristicasVehiculosResponse) error {
	return row.Scan(&t.ID, &t.Nombre, &t.Registro, &t.Vehiculos)
}

func parse2(rows *sql.Rows, t *model.CaracteristicasVehiculosResponse) error {
	return rows.Scan(&t.ID, &t.Nombre, &t.Registro, &t.Vehiculos)
}

func validarUnicoNombre(db *sql.DB, nombre string) error {
	tam := 0
	sql := `select count(nombre) from caracteristicas where nombre = ?`
	db.QueryRow(sql, nombre).Scan(&tam)
	if tam > 0 {
		return errors.New("el nombre ya se encuentra registrado")
	}
	return nil
}
