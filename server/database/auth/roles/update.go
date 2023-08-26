package roles

import (
	"database/sql"
	"errors"
	"inventarios/graph/model"
)

func ModificarRol(db *sql.DB, input model.UpdateRol) (*model.Rol, error) {
	if input.ID == "1" {
		return nil, errors.New("la modificacion para el rol administrador esta restringida en el sistema (BD)")
	}
	if err := verificarMismoNombre(db, input.Nombre, input.ID); err != nil {
		return nil, err
	}
	if err := verificarExisteRol(db, input.ID); err != nil {
		return nil, err
	}
	rol, err := GetById(db, input.ID)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = quitardePermisos(db, tx, rol)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	tx, err = db.Begin()
	if err != nil {
		return nil, err
	}

	// idrol, _ := strconv.ParseInt(rol.ID, 10, 32)

	xrol := model.NewRol{
		Nombre:   input.Nombre,
		Permisos: input.Permisos,
	}

	_, err = asignarAPermisos(tx, db, xrol, rol.Bit)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	upd := `
	update roles set nombre=? where id=?
	`
	_, errx := tx.Exec(upd, input.Nombre, input.ID)
	if errx != nil {
		return nil, errx
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return GetById(db, input.ID)
}
