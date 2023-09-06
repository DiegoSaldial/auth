package viajes

import (
	"database/sql"
	"errors"
	"fmt"
	"taxis/database/auth/roles"
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

func checkViajesActivos(db *sql.DB, pasajero_id string) error {
	viajes := 0
	sql := `
	select count(pasajero_id) from viajes where pasajero_id = ? and estado = 1; 
	`
	db.QueryRow(sql, pasajero_id).Scan(&viajes)
	if viajes > 0 {
		return errors.New("ya tiene un viaje activo, debe ser finalizado o cancelado antes de realizar una nueva solicitud")
	}
	return nil
}

func checkUserHasPerm(db *sql.DB, usuarioid, permiso string) error {
	roles, err := roles.ListarRolesByUsuario(db, usuarioid, true)
	if err != nil {
		return err
	}

	for _, rol := range roles {
		for _, perm := range rol.Permisos {
			if perm.Metodo == permiso {
				return nil
			}
		}
	}

	r := fmt.Sprintf("el usuario con ID: %s, no tiene acceso al permiso: %s", usuarioid, permiso)
	return errors.New(r)
}
