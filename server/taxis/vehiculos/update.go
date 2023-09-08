package vehiculos

import (
	"database/sql"
	"taxis/graph/model"
	"taxis/taxis/caracteristicas"
	"taxis/taxis/firestore"
)

func actualizar(db *sql.DB, input model.CreateVehiculos) (*model.VehiculosResponse, error) {
	sql := `
	update vehiculos set placa=?,puertas=?,capacidad=?,descripcion=?,color=?,modelo=?,anio=?,categoria_id=? where id=?
	`
	_, err := db.Exec(sql, input.Placa, input.Puertas, input.Capacidad, input.Descripcion, input.Color, input.Modelo, input.Anio, input.CategoriaID, input.ID)
	if err != nil {
		return nil, err
	}

	sql = `delete from caracteristicas_vehiculo where vehiculo_id=?`
	_, err = db.Exec(sql, input.ID)
	if err != nil {
		return nil, err
	}

	if len(input.Caracteristicas) > 0 {
		caracteristicas.AsignarCaracteristicas(db, input, *input.ID)
	}

	if input.FotoURL != nil {
		foto_url, er := firestore.SubirImagen(*input.FotoURL, input.Placa+".jpg", true)
		if er == nil {
			return actualizarFoto(db, foto_url, *input.ID)
		}
	}
	return GetById(db, *input.ID)
}

func actualizarFoto(db *sql.DB, fotourl, id string) (*model.VehiculosResponse, error) {
	sql := `
	update vehiculos set foto_url=? where id=?
	`
	_, err := db.Exec(sql, fotourl, id)
	if err != nil {
		return nil, err
	}
	return GetById(db, id)
}

func actualizarFotoTx(tx *sql.Tx, fotourl, id string) error {
	sql := `
	update vehiculos set foto_url=? where id=?
	`
	_, err := tx.Exec(sql, fotourl, id)
	if err != nil {
		return err
	}
	return nil
}
