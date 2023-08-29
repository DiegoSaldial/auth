package conductorvehiculos

import (
	"database/sql"
	"taxis/graph/model"
)

func QuitarVehiculo(db *sql.DB, input model.CreateConductorVehiculos) (bool, error) {
	sql := `update conductor_vehiculos set estado=0 where usuario_id=? and vehiculo_id=?`
	res, err := db.Exec(sql, input.UsuarioID, input.VehiculoID)
	if err != nil {
		return false, err
	}

	idd, _ := res.RowsAffected()
	return idd > 0, nil
}

func reAsignar(db *sql.DB, input model.CreateConductorVehiculos) (*model.ConductorVehiculos, error) {
	sql := `update conductor_vehiculos set estado=1 where usuario_id=? and vehiculo_id=?`
	_, err := db.Exec(sql, input.UsuarioID, input.VehiculoID)
	if err != nil {
		return nil, err
	}

	return GetByUsuVehiculo(db, input.UsuarioID, input.VehiculoID)
}
