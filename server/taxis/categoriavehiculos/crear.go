package categoriavehiculos

import (
	"database/sql"
	"strconv"
	"taxis/graph/model"
)

func Crear(db *sql.DB, input model.CreateCategoriaVehiculos) (*model.CategoriaVehiculos, error) {
	if input.ID != nil {
		return actualizar(db, input)
	}

	sql := `insert into categoria_vehiculos(nombre,descripcion) values(?,?)`
	res, err := db.Exec(sql, input.Nombre, input.Descripcion)
	if err != nil {
		return nil, err
	}
	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)
	return GetById(db, id)

}
