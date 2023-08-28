package categoriavehiculos

import (
	"database/sql"
	"taxis/graph/model"
)

func actualizar(db *sql.DB, input model.CreateCategoriaVehiculos) (*model.CategoriaVehiculos, error) {
	sql := `update categoria_vehiculos set nombre = ?, descripcion = ? where id = ?`
	_, err := db.Exec(sql, input.Nombre, input.Descripcion, input.ID)
	if err != nil {
		return nil, err
	}
	return GetById(db, *input.ID)
}
