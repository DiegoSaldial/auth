package xauth

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"strconv"
	"taxis/database/auth/tokens"
	"taxis/graph/model"
)

func Login(ctx context.Context, db *sql.DB, u string, p string) ([]*model.Rol, error) {
	roles := []*model.Rol{}
	sql := `
	SELECT DISTINCT r.id, r.nombre, r.bit
	FROM usuarios u
	JOIN rol_usuario ru ON u.id = ru.usuario_id
	JOIN roles r ON ru.rol_id = r.id
	WHERE u.username = ? AND u.password = ? and u.estado=1
	`
	rows, err := db.Query(sql, u, p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rol := model.Rol{}
		er := parseRoles(rows, &rol)
		if er == nil {
			roles = append(roles, &rol)
		}
	}
	if len(roles) == 0 {
		return nil, errors.New("datos incorrectos o usuario inactivo")
	}
	return roles, nil
}

func UseRol(ctx context.Context, db *sql.DB, u string, p string, rol_id string) (*model.UsuarioLogin, error) {
	sql := `
	SELECT u.id,u.nombres,u.apellidos,u.username,u.foto_url,u.telefono,u.correo,u.registrado,u.estado,r.nombre as rol, r.id as rolid
	FROM usuarios u
	JOIN rol_usuario ru ON u.id = ru.usuario_id
	JOIN roles r ON ru.rol_id = r.id and r.id=?
	WHERE u.username = ? and u.password = ? and r.id= ?
	`
	rows, err := db.Query(sql, rol_id, u, p, rol_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usuario := model.UsuarioLogin{}
	for rows.Next() {
		er := parse(rows, &usuario)
		if er != nil {
			return nil, er
		}
	}
	if usuario.ID == "" {
		return nil, errors.New("usuario, clave o rol incorrectos")
	}
	if !usuario.Estado {
		return nil, errors.New("usuario tiene el estado inactivo")
	}

	idrol, _ := strconv.Atoi(usuario.RolID)
	duration := os.Getenv("TOKEN_DURATION")
	tokenduration, err := strconv.Atoi(duration)
	if err != nil {
		tokenduration = 60
	}

	token, er2 := GenerateToken(ctx, usuario.ID, usuario.Rol, idrol, tokenduration)
	if er2 != nil {
		return nil, er2
	}

	err = tokens.Guardar(db, token, u)
	if err != nil {
		return nil, err
	}

	usuario.Exp = strconv.Itoa(tokenduration) + " minutos."
	dn, er := encrypt(u)
	if er != nil {
		return nil, er
	}
	usuario.Dataname = dn
	return &usuario, nil
}
