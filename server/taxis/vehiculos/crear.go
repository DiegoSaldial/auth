package vehiculos

import (
	"database/sql"
	"strconv"
	"taxis/graph/model"
	"taxis/taxis/caracteristicas"
	"taxis/taxis/firestore"
)

func Crear(db *sql.DB, input model.CreateVehiculos) (*model.VehiculosResponse, error) {
	if err := verificarPlacaUnica(db, input.Placa, input.ID); err != nil {
		return nil, err
	}
	if input.ID != nil {
		return actualizar(db, input)
	}
	sql := `
	insert into vehiculos (placa,puertas,capacidad,descripcion,color,modelo,anio,categoria_id)
	values(?,?,?,?,?,?,?,?)
	`
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	res, err := tx.Exec(sql, input.Placa, input.Puertas, input.Capacidad, input.Descripcion, input.Color, input.Modelo, input.Anio, input.CategoriaID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)

	if len(input.Caracteristicas) > 0 {
		err := caracteristicas.AsignarCaracteristicasTx(tx, input, id)
		if err != nil {
			return nil, err
		}
	}

	if input.FotoURL != nil {
		foto_url, er := firestore.SubirImagen(*input.FotoURL, input.Placa+".jpg", true)
		if er == nil {
			actualizarFotoTx(tx, foto_url, id)
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return GetById(db, id)
}
