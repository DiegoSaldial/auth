package viajes

import (
	"database/sql"
	"taxis/graph/model"
)

func AceptarViaje(db *sql.DB, viaje_id, usuario_id string) (*model.ViajesResponse, error) {
	sql := `update viajes set conductor_id=? where id=?`
	_, err := db.Exec(sql, usuario_id, viaje_id)
	if err != nil {
		return nil, err
	}
	return GetById(db, viaje_id)
}

func CancelarViaje(db *sql.DB, viaje_id string) (*model.ViajesResponse, error) {
	sql := `update viajes set conductor_id=NULL where id=?`
	_, err := db.Exec(sql, viaje_id)
	if err != nil {
		return nil, err
	}
	return GetById(db, viaje_id)
}

func FinalizarViaje(db *sql.DB, viaje_id string) (*model.ViajesResponse, error) {
	sql := `update viajes set estado=0 where id=?`
	_, err := db.Exec(sql, viaje_id)
	if err != nil {
		return nil, err
	}
	return GetById(db, viaje_id)
}
