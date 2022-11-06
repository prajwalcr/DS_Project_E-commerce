package delivery

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/prajwalcr/DS_Project_E-commerce/io"
)

func ReserveAgent() (*Agent, error) {
	// booking the agent
	txn, _ := io.DB.Begin()

	//selecting the first available agent
	row := txn.QueryRow(`
		SELECT id, is_reserved, order_id from agents
		WHERE
			is_reserved < current_timestamp - (10 * interval '1 second') and order_id is NULL
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
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE agents
		SET
			is_reserved = current_timestamp
		WHERE id = $1`, agent.ID)
	if err != nil {
		txn.Rollback()
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
	log.Println(orderID)
	txn, _ := io.DB.Begin()

	// selecting the first available seat
	row := txn.QueryRow(`
		SELECT id, is_reserved, order_id from agents
		WHERE
			is_reserved >= current_timestamp - (10 * interval '1 second') and order_id is NULL
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
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE agents
		SET
			is_reserved = current_timestamp, order_id = $1
		WHERE id = $2`, orderID, agent.ID)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	agent.IsReserved = sql.NullTime{Time: time.Now()}
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
			id serial primary key,
			is_reserved timestamp default '2000-01-01 00:00:00',
			order_id varchar(36) default null
		);
	`)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 11; i++ {
		_, err := io.DB.Exec("insert into agents default values;")
		if err != nil {
			panic(err)
		}
	}
}
