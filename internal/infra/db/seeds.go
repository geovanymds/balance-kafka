package database

import "database/sql"

func ExecSeeds(db *sql.DB) error {

	statements := []string{
		"INSERT INTO clients (id, name, email, created_at) VALUES ('6d44a72f-5b09-4b4d-a2f8-7e14e6b50e1d', 'João Silva', 'joao@example.com', '2023-05-10'),('7c89a0f7-18c5-4f0b-a1c1-415b4e9ae51f', 'Maria Souza', 'maria@example.com', '2023-04-22') ON DUPLICATE KEY UPDATE name=name;",

		"INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('d2b896d1-4b63-4b77-b2f2-1cbb9c17b7e0', '6d44a72f-5b09-4b4d-a2f8-7e14e6b50e1d', 5000, '2023-05-10'),('8e2ec8a1-394d-4c87-83b1-8f68e4a1028f', '7c89a0f7-18c5-4f0b-a1c1-415b4e9ae51f', 8000, '2023-04-22') ON DUPLICATE KEY UPDATE balance=balance;",

		"INSERT INTO balances (account_id, value, created_at, updated_at) VALUES ('d2b896d1-4b63-4b77-b2f2-1cbb9c17b7e0', 5000, '2023-05-10', '2023-05-10'),('8e2ec8a1-394d-4c87-83b1-8f68e4a1028f', 8000, '2023-04-22', '2023-05-10') ON DUPLICATE KEY UPDATE value=value;",
	}

	for _, statement := range statements {

		stmt, err := db.Prepare(statement)

		if err != nil {
			return err
		}

		_, err = stmt.Exec()

		defer stmt.Close()

		if err != nil {
			return err
		}

	}

	return nil

}
