package store

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/prajwalcr/DS_Project_E-commerce/io"
)

func ReserveProduct(productID int) (*Packet, error) {
	log.Println("reserving product", productID)

	txn, _ := io.DB.Begin()

	row := txn.QueryRow(`
		SELECT id, product_id, reserved_timestamp, order_id
		FROM packets
		WHERE
			reserved_timestamp < current_timestamp - (10 * interval '1 second') and product_id = $1 and order_id is NULL
		LIMIT 1
		FOR UPDATE;
	`, productID)
	if row.Err() != nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var packet Packet
	err := row.Scan(&packet.ID, &packet.ProductID, &packet.IsReserved, &packet.OrderID)

	if err != nil && err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("No product packet available")
	}
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE packets
		SET
			reserved_timestamp = current_timestamp
		WHERE id = $1
	`, packet.ID)

	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return &packet, nil
}

func BookProduct(orderID string, productID int) (*Packet, error) {
	txn, _ := io.DB.Begin()
	log.Println(orderID, productID)
	row := txn.QueryRow(`
		SELECT id, product_id, reserved_timestamp, order_id from packets
		WHERE
			reserved_timestamp >= current_timestamp - (10 * interval '1 second') and order_id is NULL and product_id = $1
		LIMIT 1
		FOR UPDATE
	`, productID)
	if row.Err() != nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var packet Packet
	err := row.Scan(&packet.ID, &packet.ProductID, &packet.IsReserved, &packet.OrderID)

	if err != nil && err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("no product packet available")
	}
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE 	packets
		SET
			reserved_timestamp = current_timestamp, order_id = $1
		WHERE
			id = $2
	`, orderID, packet.ID)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	packet.IsReserved = sql.NullTime{Time: time.Now()}
	packet.OrderID = sql.NullString{String: orderID}
	return &packet, nil
}

func Clean() {
	_, err := io.DB.Exec("DROP TABLE IF EXISTS packets;")
	if err != nil {
		panic(err)
	}

	_, err = io.DB.Exec(`
		CREATE TABLE packets (
			id serial primary key,
			reserved_timestamp timestamp default '2000-01-01 00:00:00',
			order_id varchar(36) default null,
			product_id int default 1
		);
	`)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 11; i++ {
		_, err := io.DB.Exec("insert into packets default values;")
		if err != nil {
			panic(err)
		}
	}
}
