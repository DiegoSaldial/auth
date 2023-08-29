package usuarios

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"taxis/graph/model"
	"taxis/taxis/firestore"
)

func UpdateUsuario(db *sql.DB, input model.UpdateUsuario) (*model.Usuario, error) {
	if input.ID == "1" {
		return nil, errors.New("la modificacion para el usuario administrador esta restringida en el sistema (BD)")
	}
	if err := verificarRolesRepetidos(input.Roles); err != nil {
		return nil, err
	}
	if err := verificarUsernameYaAsignado(db, input.Username, &input.ID); err != nil {
		return nil, err
	}
	if err := verificarUserIDValido(db, input.ID); err != nil {
		return nil, err
	}
	if err := verificarRolesvalidos(db, input.Roles); err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	/* sql := "update usuarios set nombres=?,apellidos=?,username=?,password=?,foto_url=?,telefono=?,correo=? where id=?"
	_, err = tx.Exec(sql,
		input.Nombres, input.Apellidos, input.Username, input.Password, input.FotoURL, input.Telefono, input.Correo, input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	} */

	sql := "update usuarios set nombres=?,apellidos=?,username=?,telefono=?,correo=?,documento=?,domicilio=?,fecha_nac=?"
	params := []interface{}{input.Nombres, input.Apellidos, input.Username, input.Telefono, input.Correo, input.Documento, input.Domicilio, input.FechaNac}

	if input.Password != nil {
		sql += ",password=?"
		params = append(params, input.Password)
	}

	sql += " where id=?"

	params = append(params, input.ID)

	_, err = tx.Exec(sql, params...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("delete from rol_usuario where usuario_id=?", input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id := input.ID

	sql = "insert into rol_usuario(usuario_id,rol_id) values %s"
	places := make([]string, len(input.Roles))
	args := make([]interface{}, len(input.Roles)*2)
	for i, v := range input.Roles {
		places[i] = "(?,?)"
		args[i*2] = id
		args[i*2+1] = v
	}
	sql = fmt.Sprintf(sql, strings.Join(places, ", "))

	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if input.FotoURL != nil {
		foto_url, er := firestore.SubirImagen(*input.FotoURL, input.Username+".jpg")
		if er == nil {
			return actualizarFoto(db, foto_url, input.Username)
		}
	}

	return usuarioByUname(db, input.Username)
}

func actualizarFoto(db *sql.DB, fotourl, username string) (*model.Usuario, error) {
	sql := `
	update usuarios set foto_url=? where username=?
	`
	_, err := db.Exec(sql, fotourl, username)
	if err != nil {
		return nil, err
	}
	return usuarioByUname(db, username)
}
