package roles

import (
	"database/sql"
	"errors"
)

func DeleteRol(db *sql.DB, rolid string) (bool, error) {
	if rolid == "1" {
		return false, errors.New("la eliminacion para el rol administrador esta restringida en el sistema (BD)")
	}
	if err := verificarExisteRol(db, rolid); err != nil {
		return false, err
	}
	rol, err := GetById(db, rolid)
	if err != nil {
		return false, err
	}

	tx, err := db.Begin()
	if err != nil {
		return false, err
	}

	sql := "delete from roles where id=?"
	_, err = tx.Exec(sql, rolid)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = quitardePermisos(db, tx, rol)
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}
