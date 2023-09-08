package caracteristicas

import (
	"database/sql"
	"taxis/graph/model"
)

func Actualizar(db *sql.DB, input model.NewCaracteristica) (*model.CaracteristicasVehiculosResponse, error) {
	sql := `update caracteristicas set nombre =? where id = ?`
	_, err := db.Exec(sql, input.Nombre, input.ID)
	if err != nil {
		return nil, err
	}
	return GetById(db, *input.ID)
}
