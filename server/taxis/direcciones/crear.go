package direcciones

import (
	"database/sql"
	"inventarios/graph/model"
	"strconv"
)

func Crear(db *sql.DB, input model.CreateDirecciones) (*model.Direcciones, error) {
	if input.ID != nil {
		return actualizar(db, input)
	}
	sql := `
	insert into direcciones(nombre,descripcion,ubicacion,usuario_id)
	values(?,?,POINT(?,?),?)
	`
	res, err := db.Exec(sql, input.Nombre, input.Descripcion, input.Latitud, input.Longitud, input.UsuarioID)
	if err != nil {
		return nil, err
	}
	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)
	return GetById(db, id)
}
