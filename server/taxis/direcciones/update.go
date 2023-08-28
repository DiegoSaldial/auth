package direcciones

import (
	"database/sql"
	"inventarios/graph/model"
)

func actualizar(db *sql.DB, input model.CreateDirecciones) (*model.Direcciones, error) {
	sql := `update direcciones set nombre=?, descripcion=?, ubicacion=POINT(?,?) where id=?`
	_, err := db.Exec(sql, input.Nombre, input.Descripcion, input.Latitud, input.Longitud, input.ID)
	if err != nil {
		return nil, err
	}
	return GetById(db, *input.ID)
}
