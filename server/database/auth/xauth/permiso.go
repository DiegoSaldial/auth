package xauth

import (
	"database/sql"
	"errors"
	"fmt"
	"inventarios/database/auth/roles"
	"strconv"
)

func CheckPermiso(db *sql.DB, metodo string, rolid int, rol string) error {
	r, er := roles.GetById(db, strconv.Itoa(rolid))
	if er != nil {
		return er
	}

	rol_bit := 1 << (r.Bit - 1)
	rol_bits := 0
	sql := "select rol_bits from rol_permiso where metodo=?"
	db.QueryRow(sql, metodo).Scan(&rol_bits)
	permitido := rol_bits&rol_bit == rol_bit

	if permitido {
		return nil
	}
	texto := fmt.Sprintf("acceso denegado[%s]: solicitado: %d, devuelto: %d, para tu rol: %s", metodo, rol_bit, rol_bits, rol)
	return errors.New(texto)
}
