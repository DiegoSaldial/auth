package usuarios

import (
	"database/sql"
	"errors"
	"fmt"
	"inventarios/graph/model"
	"strings"
)

func parseRow(row *sql.Row, t *model.Usuario) error {
	return row.Scan(
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
	)
}

/*
	 func parse(rows *sql.Rows, t *model.Usuario) error {
		return rows.Scan(
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
		)
	}
*/
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

func verificarRolesRepetidos(arr []string) error {
	ocurrencias := make(map[string]bool)
	repetidos := []string{}

	for _, elemento := range arr {
		if ocurrencias[elemento] {
			repetidos = append(repetidos, elemento)
		}
		ocurrencias[elemento] = true
	}

	if len(repetidos) == 0 {
		return nil
	}

	rep := fmt.Sprintf("%+v", repetidos)
	return errors.New("roles repetidos: " + rep)
}

func verificarUserIDValido(db *sql.DB, userid string) error {
	usernameexis := 0
	db.QueryRow("select count(id) from usuarios where id=?", userid).Scan(&usernameexis)
	if usernameexis == 0 {
		return errors.New("usuario no encontrado")
	}
	return nil
}

func verificarUsernameYaAsignado(db *sql.DB, username string, userid *string) error {
	usernameexis := 0
	if userid == nil {
		cero := "0"
		userid = &cero
	}
	sql := `
	SELECT COUNT(u.username) 
	FROM usuarios u
	JOIN rol_usuario ru ON u.id = ru.usuario_id
	JOIN roles r ON ru.rol_id = r.id
	WHERE u.username = ? AND u.id != ?
	GROUP by(u.id)
	`
	db.QueryRow(sql, username, userid).Scan(&usernameexis)
	if usernameexis > 0 {
		return errors.New("el username ya lo tiene asignado otro usuario")
	}
	return nil
}

func verificarRolesvalidos(db *sql.DB, valores []string) error {
	query := `SELECT id FROM roles WHERE id IN (%s)`
	placeholders := make([]string, len(valores))
	args := make([]interface{}, len(valores))
	for i, v := range valores {
		placeholders[i] = "?"
		args[i] = v
	}
	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Crear un mapa para almacenar los valores encontrados
	existentes := make(map[string]bool)
	for rows.Next() {
		var rolID string
		if err := rows.Scan(&rolID); err != nil {
			return err
		}
		existentes[rolID] = true
	}

	// Verificar si todos los valores existen
	ausentes := []string{}
	for _, valor := range valores {
		if !existentes[valor] {
			ausentes = append(ausentes, valor)
		}
	}

	if len(ausentes) == 0 {
		return nil
	}

	rep := fmt.Sprintf("%+v", ausentes)
	return errors.New("roles no existentes: " + rep)
}
