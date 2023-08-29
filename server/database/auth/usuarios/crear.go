package usuarios

import (
	"database/sql"
	"fmt"
	"strings"
	"taxis/graph/model"
	"taxis/taxis/firestore"
)

func CreateUsuario(db *sql.DB, input model.NewUsuario) (*model.Usuario, error) {
	if err := verificarRolesRepetidos(input.Roles); err != nil {
		return nil, err
	}
	if err := verificarUsernameYaAsignado(db, input.Username, nil); err != nil {
		return nil, err
	}
	if err := verificarRolesvalidos(db, input.Roles); err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := "insert into usuarios(nombres,apellidos,username,password,telefono,correo,documento,domicilio,fecha_nac) values(?,?,?,?,?,?,?,?,?)"
	res, err := tx.Exec(sql,
		input.Nombres, input.Apellidos, input.Username, input.Password, input.Telefono, input.Correo, input.Documento, input.Domicilio, input.FechaNac)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, _ := res.LastInsertId()
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
