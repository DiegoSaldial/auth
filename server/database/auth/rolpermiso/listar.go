package rolpermiso

import (
	"database/sql"
	"errors"
	"fmt"
	"inventarios/graph/model"
	"strings"
)

func Permisos(db *sql.DB) ([]*model.RolPermiso, error) {
	permisos := []*model.RolPermiso{}
	sql := "select metodo,descripcion,rol_bits from rol_permiso"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		perm := model.RolPermiso{}
		er := parse(rows, &perm)
		if er == nil {
			permisos = append(permisos, &perm)
		} else {
			return nil, er
		}
	}
	return permisos, nil
}

func GetByMetodo(db *sql.DB, metodo string) (*model.RolPermiso, error) {
	rp := model.RolPermiso{}
	sql := "select metodo,descripcion,rol_bits from rol_permiso where metodo=?"
	row := db.QueryRow(sql, metodo)
	err := parseRow(row, &rp)
	if err != nil {
		return nil, err
	}
	return &rp, nil
}

func ListarByMetodos(db *sql.DB, metodos []string) ([]*model.RolPermiso, error) {
	permisos := []*model.RolPermiso{}
	sql := fmt.Sprintf("SELECT metodo,descripcion,rol_bits FROM rol_permiso WHERE metodo IN ('%s')", strings.Join(metodos, "','"))
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		perm := model.RolPermiso{}
		er := parse(rows, &perm)
		if er == nil {
			permisos = append(permisos, &perm)
		} else {
			return nil, er
		}
	}
	return permisos, nil
}

func PermisosByRol(db *sql.DB, rol_id string) ([]*model.RolPermiso, error) {
	permisos := []*model.RolPermiso{}
	rls := `select bit from roles where id=?`
	rol_bit := -1
	db.QueryRow(rls, rol_id).Scan(&rol_bit)
	xpermisos, err := Permisos(db)
	if err != nil {
		return nil, err
	}

	if rol_bit == -1 {
		return nil, errors.New("no existe el rol")
	}

	rol_bit = 1 << (rol_bit - 1)

	for _, p := range xpermisos {
		if p.RolBits&rol_bit == rol_bit {
			permisos = append(permisos, p)
		}
	}
	return permisos, nil
}
