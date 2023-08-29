package conductorvehiculos

import (
	"database/sql"
	"taxis/graph/model"
)

func parse(row *sql.Row, t *model.ConductorVehiculos) error {
	return row.Scan(
		&t.UsuarioID,
		&t.VehiculoID,
		&t.Estado,
		&t.Registrado,
	)
}

// from auth usuarios utils
func parseUL(rows *sql.Rows, t *model.UsuariosResponse) error {
	rol := model.Rol{}
	err := rows.Scan(
		&t.ID,
		&t.Nombres,
		&t.Apellidos,
		&t.Username,
		&t.FotoURL,
		&t.Telefono,
		&t.Correo,
		&t.Documento,
		&t.Domicilio,
		&t.FechaNac,
		&t.Registrado,
		&t.Estado,
		&rol.ID,
		&rol.Nombre,
		&rol.Bit,
	)
	t.Roles = []*model.Rol{&rol}
	return err
}

func findUsuarioByID(usuarios []*model.UsuariosResponse, targetID string) *model.UsuariosResponse {
	for _, usuario := range usuarios {
		if usuario.ID == targetID {
			return usuario
		}
	}
	return nil
}

func checkExiste(db *sql.DB, usuario, vehiculo string) bool {
	existe := 0
	sql := `select count(usuario_id) from conductor_vehiculos where usuario_id = ? and vehiculo_id = ?`
	db.QueryRow(sql, usuario, vehiculo).Scan(&existe)
	return existe > 0
}

func hayOtrosActivos(db *sql.DB, usuario string) int {
	existe := 0
	sql := `select count(usuario_id) from conductor_vehiculos where estado=1 && usuario_id = ?`
	db.QueryRow(sql, usuario).Scan(&existe)
	return existe
}
