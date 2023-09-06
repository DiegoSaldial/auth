package roles

import (
	"database/sql"
	"inventarios/database/auth/rolpermiso"
	"inventarios/graph/model"
)

func GetById(db *sql.DB, id string) (*model.Rol, error) {
	sql := "select id,nombre,bit from roles where id=?"
	row := db.QueryRow(sql, id)
	rol := model.Rol{}
	er := parseRow(row, &rol)
	if er != nil {
		return nil, er
	}
	return &rol, nil
}

func Listar(db *sql.DB, lite bool) ([]*model.RolResponse, error) {
	/* sql := `
	SELECT r.id,r.nombre, r.bit
	from roles r
	order by r.id desc
	` */
	sql := `
	SELECT r.id, r.nombre, r.bit, COUNT(ru.usuario_id) AS total_usuarios
	FROM roles r
	LEFT JOIN rol_usuario ru ON r.id = ru.rol_id
	GROUP BY r.id, r.nombre, r.bit
	ORDER BY r.id DESC;
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*model.RolResponse{}
	for rows.Next() {
		rol := model.RolResponse{}
		er := parse(rows, &rol)
		if er != nil {
			return nil, er
		}
		if !lite {
			ps, er := rolpermiso.PermisosByRol(db, rol.ID)
			if er != nil {
				return nil, er
			}
			rol.Permisos = len(ps)
		}
		roles = append(roles, &rol)
	}

	return roles, nil
}

func ListarRolesByUsuario(db *sql.DB, userid string, permisos bool) ([]*model.RolesUserResponse, error) {
	sql := `
	select r.id, r.nombre, r.bit 
	from roles r 
	inner join rol_usuario ru on ru.rol_id = r.id 
	where ru.usuario_id = ?
	GROUP BY r.id
	`
	rows, err := db.Query(sql, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*model.RolesUserResponse{}
	for rows.Next() {
		rol := model.RolesUserResponse{}
		er := parseURol(rows, &rol)
		if er != nil {
			return nil, er
		}
		if permisos {
			ps, er := rolpermiso.PermisosByRol(db, rol.ID)
			if er != nil {
				return nil, er
			}
			rol.Permisos = ps
		}

		roles = append(roles, &rol)
	}

	return roles, nil
}
