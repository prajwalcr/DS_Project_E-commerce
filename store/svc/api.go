package store

import (
	"database/sql"
	"errors"
	"log"

	"github.com/prajwalcr/DS_Project_E-commerce/io"
)

func ReserveProduct(productID int) (*Packet, error) {
	log.Println("reserving product", productID)

	txn, _ := io.DB.Begin()

	row := txn.QueryRow(`
		SELECT id, product_id, is_reserved, order_id
		FROM packets
		WHERE
			is_reserved is false and product_id = $1 and order_id is NULL
		LIMIT 1
		FOR UPDATE;
	`, productID)
	if row.Err() != nil {
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
			is_reserved = True
		WHERE id = $1
	`, packet.ID)

	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	return &packet, nil
}

func BookProduct(orderID string, productID int) (*Packet, error) {
	txn, _ := io.DB.Begin()
	log.Println(orderID, productID)
	row := txn.QueryRow(`
		SELECT id, product_id, is_reserved, order_id from packets
		WHERE
			is_reserved is true and order_id is NULL and product_id = $1
		LIMIT 1
		FOR UPDATE
	`, productID)
	if row.Err() != nil {
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
			is_reserved = false, order_id = $1
		WHERE
			id = $2
	`, orderID, packet.ID)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	packet.IsReserved = false
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
			is_reserved bool default false,
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
