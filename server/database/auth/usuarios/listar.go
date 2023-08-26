package usuarios

import (
	"database/sql"
	"errors"
	"inventarios/graph/model"
)

func usuarioByUname(db *sql.DB, username string) (*model.Usuario, error) {
	sql := `
	SELECT u.id, u.nombres, u.apellidos, u.username, u.foto_url, u.telefono, u.correo,u.documento,u.domicilio,u.fecha_nac, u.registrado, u.estado
	FROM usuarios u
	JOIN rol_usuario ru ON u.id = ru.usuario_id
	JOIN roles r ON ru.rol_id = r.id
	WHERE u.username = ? 
	GROUP BY u.id
	`
	row := db.QueryRow(sql, username)
	usuario := model.Usuario{}
	parseRow(row, &usuario)
	if usuario.ID == "" {
		return nil, errors.New("usuario no encontrado")
	}
	return &usuario, nil
}

func UsuarioByUsernameRoles(db *sql.DB, username string) (*model.UsuariosResponse, error) {
	sql := `
	SELECT u.id, u.nombres, u.apellidos, u.username, u.foto_url, u.telefono, u.correo, u.documento,u.domicilio,u.fecha_nac, u.registrado, u.estado,
		IFNULL(r.id, 0) AS rol_id,
		IFNULL(r.nombre, '') AS rol_nombre,
		IFNULL(r.bit, 0) AS rol_bit
	FROM usuarios u
	left JOIN rol_usuario ru ON u.id = ru.usuario_id
	left JOIN roles r ON ru.rol_id = r.id
	WHERE u.username = ?
	order BY ru.rol_id desc
	`
	rows, err := db.Query(sql, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []*model.UsuariosResponse{}
	for rows.Next() {
		usuario := model.UsuariosResponse{Roles: []*model.Rol{}}
		er := parseUL(rows, &usuario)
		if er != nil {
			return nil, er
		}
		existe := findUsuarioByID(usuarios, usuario.ID)
		if existe != nil {
			existe.Roles = append(existe.Roles, usuario.Roles...)
		} else {
			usuarios = append(usuarios, &usuario)

		}
	}
	if len(usuarios) > 1 {
		return nil, errors.New("error interno, hay varios usuarios listados")
	}
	return usuarios[0], nil
}

func Usuarios(db *sql.DB) ([]*model.UsuariosResponse, error) {
	sql := `
	SELECT u.id,u.nombres,u.apellidos,u.username,u.foto_url,u.telefono,u.correo,u.documento,u.domicilio,u.fecha_nac,u.registrado,u.estado,
		IFNULL(r.id, 0) AS rol_id,
		IFNULL(r.nombre, '') AS rol_nombre,
		IFNULL(r.bit, 0) AS rol_bit
	FROM
		usuarios u
	LEFT JOIN
		rol_usuario ru ON u.id = ru.usuario_id
	LEFT JOIN
		roles r ON ru.rol_id = r.id
	ORDER BY
		u.id desc, r.id;
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usuarios := []*model.UsuariosResponse{}
	for rows.Next() {
		usuario := model.UsuariosResponse{Roles: []*model.Rol{}}
		er := parseUL(rows, &usuario)
		if er != nil {
			return nil, er
		}
		existe := findUsuarioByID(usuarios, usuario.ID)
		if existe != nil {
			existe.Roles = append(existe.Roles, usuario.Roles...)
		} else {
			usuarios = append(usuarios, &usuario)

		}
	}
	return usuarios, nil
}
