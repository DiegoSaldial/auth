package vehiculos

import (
	"database/sql"
	"strconv"
	"taxis/graph/model"
	"taxis/taxis/firestore"
)

func Crear(db *sql.DB, input model.CreateVehiculos) (*model.VehiculosResponse, error) {
	if input.ID != nil {
		return actualizar(db, input)
	}
	sql := `
	insert into vehiculos (placa,puertas,capacidad,descripcion,color,modelo,anio,categoria_id)
	values(?,?,?,?,?,?,?,?)
	`
	res, err := db.Exec(sql, input.Placa, input.Puertas, input.Capacidad, input.Descripcion, input.Color, input.Modelo, input.Anio, input.CategoriaID)
	if err != nil {
		return nil, err
	}
	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)
	if input.FotoURL != nil {
		foto_url, er := firestore.SubirImagen(*input.FotoURL, input.Placa+".jpg")
		if er == nil {
			return actualizarFoto(db, foto_url, id)
		}
	}
	return GetById(db, id)
}
