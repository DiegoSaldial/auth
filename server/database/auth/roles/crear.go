package roles

import (
	"database/sql"
	"inventarios/graph/model"
	"strconv"
)

func CreateRol(db *sql.DB, input model.NewRol) (*model.Rol, error) {
	if err := verificarNombreExiste(db, input.Nombre); err != nil {
		return nil, err
	}
	if err := verificarPermisosExistentes(db, input.Permisos); err != nil {
		return nil, err
	}
	if err := verificarMaxRegistros(db); err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	bit_serial := 0
	tx.QueryRow("select count(nombre)+1 from roles").Scan(&bit_serial)

	sql := "insert into roles(nombre,bit) values(?,?)"
	res, err := tx.Exec(sql, input.Nombre, bit_serial)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	idrol, _ := res.LastInsertId()
	id_r := strconv.FormatInt(idrol, 10)

	_, err = asignarAPermisos(tx, db, input, bit_serial)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return GetById(db, id_r)
}
