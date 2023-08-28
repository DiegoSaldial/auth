package conductorvehiculos

import (
	"database/sql"
	"taxis/graph/model"
)

func Crear(db *sql.DB, input model.CreateConductorVehiculos) (*model.ConductorVehiculos, error) {
	sql := `insert into conductor_vehiculos(usuario_id,vehiculo_id) values(?,?)`
	_, err := db.Exec(sql, input.UsuarioID, input.VehiculoID)
	if err != nil {
		return nil, err
	}

	return GetByUsuVehiculo(db, input.UsuarioID, input.VehiculoID)

}
