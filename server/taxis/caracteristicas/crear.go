package caracteristicas

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"taxis/graph/model"
)

func Crear(db *sql.DB, input model.NewCaracteristica) (*model.CaracteristicasVehiculosResponse, error) {
	if input.ID != nil {
		return Actualizar(db, input)
	}
	if err := validarUnicoNombre(db, input.Nombre); err != nil {
		return nil, err
	}
	sql := `insert into caracteristicas (nombre) values (?)`
	res, err := db.Exec(sql, input.Nombre)
	if err != nil {
		return nil, err
	}
	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)
	return GetById(db, id)
}

func AsignarCaracteristicasTx(tx *sql.Tx, input model.CreateVehiculos, idvehiculo string) error {
	if err := verifExistenTx(tx, input); err != nil {
		return err
	}
	sql := `insert into caracteristicas_vehiculo(caracteristica_id,vehiculo_id) values `
	caras := input.Caracteristicas

	placeholders := make([]string, 0, len(caras)*2)
	values := make([]interface{}, 0, len(caras)*2)

	for _, caracteristica := range caras {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, caracteristica, idvehiculo)
	}

	sql += strings.Join(placeholders, ",")

	_, err := tx.Exec(sql, values...)
	if err != nil {
		return err
	}

	return nil
}

func verifExistenTx(tx *sql.Tx, input model.CreateVehiculos) error {
	sql := `select id from caracteristicas where id in (`
	caras := input.Caracteristicas

	placeholders := make([]string, 0, len(caras)*2)
	values := make([]interface{}, 0, len(caras)*2)

	for _, caracteristica := range caras {
		placeholders = append(placeholders, "?")
		values = append(values, caracteristica)
	}

	sql += strings.Join(placeholders, ",")
	sql += ")"

	rows, err := tx.Query(sql, values...)
	if err != nil {
		return err
	}
	defer rows.Close()
	noexistentes := []string{}
	for rows.Next() {
		d := ""
		rows.Scan(&d)
		existe := false
		for _, idc := range caras {
			if d == idc {
				existe = true
				break
			}
		}
		if !existe {
			noexistentes = append(noexistentes, d)
		}
	}

	if len(noexistentes) > 0 {
		r := fmt.Sprintf("no existen las caracteristicas: %s", noexistentes)
		return errors.New(r)
	}

	return nil
}

func AsignarCaracteristicas(db *sql.DB, input model.CreateVehiculos, idvehiculo string) error {
	if err := verifExisten(db, input); err != nil {
		return err
	}
	sql := `insert into caracteristicas_vehiculo(caracteristica_id,vehiculo_id) values `
	caras := input.Caracteristicas

	placeholders := make([]string, 0, len(caras)*2)
	values := make([]interface{}, 0, len(caras)*2)

	for _, caracteristica := range caras {
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, caracteristica, idvehiculo)
	}

	sql += strings.Join(placeholders, ",")

	_, err := db.Exec(sql, values...)
	if err != nil {
		return err
	}

	return nil
}

func verifExisten(db *sql.DB, input model.CreateVehiculos) error {
	sql := `select id from caracteristicas where id in (`
	caras := input.Caracteristicas

	placeholders := make([]string, 0, len(caras)*2)
	values := make([]interface{}, 0, len(caras)*2)

	for _, caracteristica := range caras {
		placeholders = append(placeholders, "?")
		values = append(values, caracteristica)
	}

	sql += strings.Join(placeholders, ",")
	sql += ")"

	rows, err := db.Query(sql, values...)
	if err != nil {
		return err
	}
	defer rows.Close()
	noexistentes := []string{}
	for rows.Next() {
		d := ""
		rows.Scan(&d)
		existe := false
		for _, idc := range caras {
			if d == idc {
				existe = true
				break
			}
		}
		if !existe {
			noexistentes = append(noexistentes, d)
		}
	}

	if len(noexistentes) > 0 {
		r := fmt.Sprintf("no existen las caracteristicas: %s", noexistentes)
		return errors.New(r)
	}

	return nil
}
