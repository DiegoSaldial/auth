package conductorvehiculos

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func GetByUsuVehiculo(db *sql.DB, usuario, vehiculo string) (*model.ConductorVehiculos, error) {
	sql := `select usuario_id,vehiculo_id,estado,registrado from conductor_vehiculos where usuario_id = ? and vehiculo_id = ?`
	row := db.QueryRow(sql, usuario, vehiculo)
	if row == nil {
		return nil, errors.New("no existe vehiculo con el conductor")
	}
	cv := model.ConductorVehiculos{}
	er := parse(row, &cv)
	if er != nil {
		return nil, er
	}
	return &cv, nil
}

func UsuariosByVehiculo(db *sql.DB, vehiculoid string) ([]*model.UsuariosResponse, error) {
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
		conductor_vehiculos cv ON u.id = cv.usuario_id
	LEFT JOIN
		roles r ON ru.rol_id = r.id
	WHERE 
		cv.vehiculo_id = ?
	ORDER BY
		u.id desc, r.id;
	`
	rows, err := db.Query(sql, vehiculoid)
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

func Conductores(db *sql.DB) ([]*model.UsuariosResponse, error) {
	sql := `
	SELECT u.id,u.nombres,u.apellidos,u.username,u.foto_url,u.telefono,u.correo,u.documento,u.domicilio,u.fecha_nac,u.registrado,u.estado,
		IFNULL(r.id, 0) AS rol_id,
		IFNULL(r.nombre, '') AS rol_nombre,
		IFNULL(r.bit, 0) AS rol_bit
	FROM
		usuarios u
	INNER JOIN 
		conductor_vehiculos cv ON u.id = cv.usuario_id
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

func Clientes(db *sql.DB) ([]*model.UsuariosResponse, error) {
	sql := `
	SELECT u.id,u.nombres,u.apellidos,u.username,u.foto_url,u.telefono,u.correo,u.documento,u.domicilio,u.fecha_nac,u.registrado,u.estado,
		IFNULL(r.id, 0) AS rol_id,
		IFNULL(r.nombre, '') AS rol_nombre,
		IFNULL(r.bit, 0) AS rol_bit
	FROM
		usuarios u
	LEFT JOIN
		conductor_vehiculos cv ON u.id = cv.usuario_id
	LEFT JOIN
		rol_usuario ru ON u.id = ru.usuario_id
	LEFT JOIN
		roles r ON ru.rol_id = r.id 
	INNER JOIN
		viajes vi ON u.id= vi.pasajero_id 
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
