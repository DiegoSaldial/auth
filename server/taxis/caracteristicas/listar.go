package caracteristicas

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func Listar(db *sql.DB) ([]*model.CaracteristicasVehiculosResponse, error) {
	sql := `
	SELECT
		c.id,
		c.nombre,
		c.registrado,
		COUNT(cv.caracteristica_id) AS cantidad_vehiculos
	FROM
		caracteristicas c
	LEFT JOIN
		caracteristicas_vehiculo cv ON c.id = cv.caracteristica_id
	GROUP BY
		c.id, c.nombre, c.registrado;
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cs := []*model.CaracteristicasVehiculosResponse{}
	for rows.Next() {
		ca := model.CaracteristicasVehiculosResponse{}
		err := parse2(rows, &ca)
		if err != nil {
			return nil, err
		}
		cs = append(cs, &ca)
	}

	return cs, nil
}

func ListarByVehiculo(db *sql.DB, id string) ([]*model.CaracteristicasVehiculosResponse, error) {
	sql := `
	select c.id,c.nombre, c.registrado, -1 
	from caracteristicas c 
	inner join caracteristicas_vehiculo cv on cv.caracteristica_id = c.id
	where cv.vehiculo_id = ?
	`
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cs := []*model.CaracteristicasVehiculosResponse{}
	for rows.Next() {
		ca := model.CaracteristicasVehiculosResponse{}
		err := parse2(rows, &ca)
		if err != nil {
			return nil, err
		}
		cs = append(cs, &ca)
	}

	return cs, nil
}

func GetById(db *sql.DB, id string) (*model.CaracteristicasVehiculosResponse, error) {
	sql := `
	SELECT
		c.id,
		c.nombre,
		c.registrado,
		COUNT(cv.caracteristica_id) AS cantidad_vehiculos
	FROM
		caracteristicas c
	LEFT JOIN
		caracteristicas_vehiculo cv ON c.id = cv.caracteristica_id
	WHERE c.id = ?
	GROUP BY
		c.id, c.nombre, c.registrado;
	`
	row := db.QueryRow(sql, id)
	if row == nil {
		return nil, errors.New("no existe el registro")
	}
	ca := model.CaracteristicasVehiculosResponse{}
	err := parse(row, &ca)
	if err != nil {
		return nil, err
	}
	return &ca, nil
}

// POR HACER:
// VALIDAR QUE AL MOMENTO DE ASIGNAR UN CONDUCTOR AL VEHICULO,
// ESTE TENGA EL PERMISO 'aceptarViaje'
