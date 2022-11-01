package transport

import (
	"database/sql"
	"errors"
	"log"

	"github.com/prajwalcr/DS_Project_E-commerce/io"
)

func ReserveVehicle() (*Vehicle, error) {
	// booking the vehicle
	txn, _ := io.DB.Begin()

	//selecting the first available vehicle
	row := txn.QueryRow(`
		SELECT id, is_reserved, order_id from vehicles
		WHERE
			is_reserved is false and order_id is NULL
		LIMIT 1
		FOR UPDATE
	`)
	if row.Err() != nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var vehicle Vehicle
	err := row.Scan(&vehicle.ID, &vehicle.IsReserved, &vehicle.OrderID)

	if err != nil && err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("no transport vehicle available")
	}
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE vehicles
		SET
			is_reserved = true
		WHERE id = $1`, vehicle.ID)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func BookVehicle(orderID string) (*Vehicle, error) {

	log.Println(orderID)
	txn, _ := io.DB.Begin()


	row := txn.QueryRow(`
		SELECT id, is_reserved, order_id from vehicles
		WHERE
			is_reserved is true and order_id is NULL
		LIMIT 1
		FOR UPDATE
	`)
	if row.Err() != nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var vehicle Vehicle
	err := row.Scan(&vehicle.ID, &vehicle.IsReserved, &vehicle.OrderID)

	if err != nil && err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("no transport vehicle available")
	}
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE vehicles
		SET
			is_reserved = false, order_id = $1
		WHERE id = $2`, orderID, vehicle.ID)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	vehicle.IsReserved = false
	vehicle.OrderID = sql.NullString{String: orderID}
	return &vehicle, nil
}

func Clean() {
	_, err := io.DB.Exec("DROP TABLE IF EXISTS vehicles;")
	if err != nil {
		panic(err)
	}

	_, err = io.DB.Exec(`
		CREATE TABLE vehicles (
			id serial primary key,
			is_reserved bool default false,
			order_id varchar(36) default null
		);
	`)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 11; i++ {
		_, err := io.DB.Exec("insert into vehicles default values;")
		if err != nil {
			panic(err)
		}
	}
}