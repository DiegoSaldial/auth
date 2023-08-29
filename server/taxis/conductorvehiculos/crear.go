package conductorvehiculos

import (
	"database/sql"
	"errors"
	"fmt"
	"taxis/graph/model"
)

func Crear(db *sql.DB, input model.CreateConductorVehiculos) (*model.ConductorVehiculos, error) {
	hayOtros := hayOtrosActivos(db, input.UsuarioID)
	if hayOtros > 0 {
		text := fmt.Sprintf("no puede ser conductor en mas de 1 vehiculo a la ves, quite de los demas [%d]", hayOtros)
		return nil, errors.New(text)
	}

	existe := checkExiste(db, input.UsuarioID, input.VehiculoID)
	if existe {
		return reAsignar(db, input)
	}

	sql := `insert into conductor_vehiculos(usuario_id,vehiculo_id) values(?,?)`
	_, err := db.Exec(sql, input.UsuarioID, input.VehiculoID)
	if err != nil {
		return nil, err
	}

	return GetByUsuVehiculo(db, input.UsuarioID, input.VehiculoID)

}
