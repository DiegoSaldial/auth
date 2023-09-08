package vehiculos

import (
	"database/sql"
	"errors"
	"taxis/graph/model"
)

func parse(rows *sql.Rows, t *model.VehiculosResponse) error {
	return rows.Scan(
		&t.ID,
		&t.Placa,
		&t.Puertas,
		&t.Capacidad,
		&t.Descripcion,
		&t.Color,
		&t.Modelo,
		&t.Anio,
		&t.CategoriaID,
		&t.FotoURL,
		&t.Estado,
		&t.Registrado,
		&t.ConductorID,
		&t.Conductor,
		&t.Categoria,
		&t.EstadoConductorVehiculo,
	)
}

func parse2(row *sql.Row, t *model.VehiculosResponse) error {
	return row.Scan(
		&t.ID,
		&t.Placa,
		&t.Puertas,
		&t.Capacidad,
		&t.Descripcion,
		&t.Color,
		&t.Modelo,
		&t.Anio,
		&t.CategoriaID,
		&t.FotoURL,
		&t.Estado,
		&t.Registrado,
		&t.ConductorID,
		&t.Conductor,
		&t.Categoria,
		&t.EstadoConductorVehiculo,
	)
}

func verificarPlacaUnica(db *sql.DB, placa string, id *string) error {
	tam := 0
	sql := `select count(id) from vehiculos where placa=?`
	if id != nil {
		sql += ` and id != ?`
		db.QueryRow(sql, placa, id).Scan(&tam)
	} else {
		db.QueryRow(sql, placa).Scan(&tam)
	}
	if tam > 0 {
		return errors.New("Placa ya registrada para otro vehiculo")
	}
	return nil
}

func getSqlListar() string {
	sql := `
	SELECT
		v.id,
		v.placa,
		v.puertas,
		v.capacidad,
		v.descripcion,
		v.color,
		v.modelo,
		v.anio,
		v.categoria_id,
		v.foto_url,
		v.estado,
		v.registrado,
		u.id AS conductor_id,
		CONCAT(u.nombres,' ',u.apellidos) AS ultimo_conductor,
		cvf.nombre AS categoria_vehiculo,
		cv.estado AS estado_conductor_vehiculo
	FROM
		vehiculos v
	LEFT JOIN (
		SELECT
			cv1.vehiculo_id,
			cv1.usuario_id,
			cv1.registrado,
			cv1.estado
		FROM
			conductor_vehiculos cv1
		WHERE
			cv1.registrado = (
				SELECT MAX(cv2.registrado)
				FROM conductor_vehiculos cv2
				WHERE cv2.vehiculo_id = cv1.vehiculo_id and cv2.estado=1
			) and cv1.estado = 1
	) cv ON v.id = cv.vehiculo_id
	LEFT JOIN usuarios u ON cv.usuario_id = u.id
	JOIN categoria_vehiculos cvf ON v.categoria_id = cvf.id
	ORDER BY v.id desc
	`
	return sql
}

func getSqlByConductor() string {
	sql := `
	SELECT
		v.id,
		v.placa,
		v.puertas,
		v.capacidad,
		v.descripcion,
		v.color,
		v.modelo,
		v.anio,
		v.categoria_id,
		v.foto_url,
		v.estado,
		v.registrado,
		u.id AS conductor_id,
		CONCAT(u.nombres,' ',u.apellidos) AS ultimo_conductor,
		cvf.nombre AS categoria_vehiculo,
		cv.estado AS estado_conductor_vehiculo
	FROM
		vehiculos v
	LEFT JOIN (
		SELECT
			cv1.vehiculo_id,
			cv1.usuario_id,
			cv1.registrado,
			cv1.estado
		FROM
			conductor_vehiculos cv1
		
	) cv ON v.id = cv.vehiculo_id
	LEFT JOIN usuarios u ON cv.usuario_id = u.id
	JOIN categoria_vehiculos cvf ON v.categoria_id = cvf.id
	where cv.usuario_id = ?
	ORDER BY v.id desc
	`
	return sql
}

func getSqlByCategoria() string {
	sql := `
	SELECT
		v.id,
		v.placa,
		v.puertas,
		v.capacidad,
		v.descripcion,
		v.color,
		v.modelo,
		v.anio,
		v.categoria_id,
		v.foto_url,
		v.estado,
		v.registrado,
		u.id AS conductor_id,
		CONCAT(u.nombres,' ',u.apellidos) AS ultimo_conductor,
		cvf.nombre AS categoria_vehiculo,
		cv.estado AS estado_conductor_vehiculo
	FROM
		vehiculos v
	LEFT JOIN (
		SELECT
			cv1.vehiculo_id,
			cv1.usuario_id,
			cv1.registrado,
			cv1.estado
		FROM
			conductor_vehiculos cv1
		
	) cv ON v.id = cv.vehiculo_id
	LEFT JOIN usuarios u ON cv.usuario_id = u.id
	JOIN categoria_vehiculos cvf ON v.categoria_id = cvf.id
	where cvf.id = ?
	GROUP BY v.id
	ORDER BY v.id desc
	`
	return sql
}

func getSqlByID() string {
	sql := `
	SELECT
		v.id,
		v.placa,
		v.puertas,
		v.capacidad,
		v.descripcion,
		v.color,
		v.modelo,
		v.anio,
		v.categoria_id,
		v.foto_url,
		v.estado,
		v.registrado,
		u.id AS conductor_id,
		CONCAT(u.nombres,' ',u.apellidos) AS ultimo_conductor,
		cvf.nombre AS categoria_vehiculo,
		cv.estado AS estado_conductor_vehiculo
	FROM
		vehiculos v
	LEFT JOIN (
		SELECT
			cv1.vehiculo_id,
			cv1.usuario_id,
			cv1.registrado,
			cv1.estado
		FROM
			conductor_vehiculos cv1
		WHERE
			cv1.registrado = (
				SELECT MAX(cv2.registrado)
				FROM conductor_vehiculos cv2
				WHERE cv2.vehiculo_id = cv1.vehiculo_id
			)
	) cv ON v.id = cv.vehiculo_id
	LEFT JOIN usuarios u ON cv.usuario_id = u.id
	JOIN categoria_vehiculos cvf ON v.categoria_id = cvf.id
	where v.id = ?
	`
	return sql
}

func getSqlByCaracteristica() string {
	sql := `
	SELECT
		v.id,
		v.placa,
		v.puertas,
		v.capacidad,
		v.descripcion,
		v.color,
		v.modelo,
		v.anio,
		v.categoria_id,
		v.foto_url,
		v.estado,
		v.registrado,
		u.id AS conductor_id,
		CONCAT(u.nombres,' ',u.apellidos) AS ultimo_conductor,
		cvf.nombre AS categoria_vehiculo,
		cv.estado AS estado_conductor_vehiculo
	FROM
		vehiculos v
	LEFT JOIN (
		SELECT
			cv1.vehiculo_id,
			cv1.usuario_id,
			cv1.registrado,
			cv1.estado
		FROM
			conductor_vehiculos cv1
		
	) cv ON v.id = cv.vehiculo_id
	LEFT JOIN usuarios u ON cv.usuario_id = u.id
	JOIN categoria_vehiculos cvf ON v.categoria_id = cvf.id
	LEFT JOIN (
		SELECT
			cav.vehiculo_id,
			cav.caracteristica_id,
			cav.registrado
		FROM
			caracteristicas_vehiculo cav
		
	) cav ON v.id = cav.vehiculo_id
	where cav.caracteristica_id = ? 
	GROUP BY v.id
	ORDER BY v.id desc
	`
	return sql
}
