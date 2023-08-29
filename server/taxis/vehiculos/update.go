package vehiculos

import (
	"database/sql"
	"taxis/graph/model"
)

func actualizar(db *sql.DB, input model.CreateVehiculos) (*model.VehiculosResponse, error) {
	sql := `
	update vehiculos set placa=?,puertas=?,capacidad=?,descripcion=?,color=?,modelo=?,anio=?,categoria_id=?,foto_url=? where id=?
	`
	_, err := db.Exec(sql, input.Placa, input.Puertas, input.Capacidad, input.Descripcion, input.Color, input.Modelo, input.Anio, input.CategoriaID, input.FotoURL, input.ID)
	if err != nil {
		return nil, err
	}
	return GetById(db, *input.ID)
}
