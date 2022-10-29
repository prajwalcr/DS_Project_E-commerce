package delivery

import (
	"database/sql"
	"errors"
	"log"
	"github.com/prajwalcr/DS_Project_E-commerce/io"
)

func ReserveAgent() (*Agent, error) {
	// booking the seat
	txn, _ := io.DB.Begin()

	//selecting the first available seat
	row := txn.QueryRow(`
		SELECT id, is_reserved, order_id from agents
		WHERE
			is_reserved is false and order_id is NULL
		LIMIT 1
		FOR UPDATE
	`)
	if row.Err() != nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var agent Agent
	err := row.Scan(&agent.ID, &agent.IsReserved, &agent.OrderID)

	if err != nil {
		txn.Rollback()
		return nil, errors.New("no delivery agent available")
	}

	_, err = txn.Exec(`
		UPDATE agents
		SET
			is_reserved = true
		WHERE id = ?`, agent.ID)
	if err != nil {
		txn.Rollack()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

func BookAgent(orderID string) (*Agent, error) {
	// booking the seat
	txn, _ := io.DB.Begin()

	// selecting the first available seat
	row := txn.QueryRow(`
		SELECT id, is_reserved, order_id from agents
		WHERE
			is_reserved is true and order_id is NULL
		LIMIT 1
		FOR UPDATE
	`)
	if row.Err() != nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var agent Agent
	err := row.Scan(&agent.ID, &agent.IsReserved, &agent.OrderID)

	if err != nil && err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("no delivery agent available")
	}

	_, err = txn.Exec(`
		UPDATE agents
		SET
			is_reserved = false, orderID = ?
		WHERE id = ?`, orderID, agent.ID)
	if err != nil {
		txn.Rollack()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	agent.IsReserved = false
	agent.OrderID = sql.NullString{String: orderID}
	return &agent, nil
}

func Clean() {
	_, err := io.DB.Exec("DROP TABLE IF EXISTS agents;")
	if err != nil {
		panic(err)
	}

	_, err = io.DB.Exec(`
		CREATE TABLE agents (
			id int primary key auto_increment not null,
			is_reserved bool default false,
			order_id varchar(36) default null
		)
	`)
}