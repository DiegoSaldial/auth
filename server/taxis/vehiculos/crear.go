package vehiculos

import (
	"database/sql"
	"inventarios/graph/model"
	"strconv"
)

func Crear(db *sql.DB, input model.CreateVehiculos) (*model.Vehiculos, error) {
	if input.ID != nil {
		return actualizar(db, input)
	}
	sql := `
	insert into vehiculos (placa,puertas,capacidad,descripcion,color,modelo,anio,categoria_id,foto_url)
	values(?,?,?,?,?,?,?,?,?)
	`
	res, err := db.Exec(sql, input.Placa, input.Puertas, input.Capacidad, input.Descripcion, input.Color, input.Modelo, input.Anio, input.CategoriaID, input.FotoURL)
	if err != nil {
		return nil, err
	}
	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)
	return GetById(db, id)
}
