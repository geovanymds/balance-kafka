CREATE TABLE clients (
    id VARCHAR(255) PRIMARY KEY, 
    name VARCHAR(255) NOT NULL, 
    email VARCHAR(255) NOT NULL, 
    created_at DATE NOT NULL
);

CREATE TABLE accounts (
    id VARCHAR(255) PRIMARY KEY, 
    client_id VARCHAR(255) NOT NULL, 
    balance int NOT NULL, 
    created_at DATE NOT NULL
);

CREATE TABLE transactions (
    id VARCHAR(255) PRIMARY KEY, 
    account_id_from VARCHAR(255) NOT NULL, 
    account_id_to VARCHAR(255) NOT NULL, 
    amount int NOT NULL, 
    created_at DATE NOT NULL
);

CREATE TABLE balances (
    id VARCHAR(255) PRIMARY KEY, 
    account_id VARCHAR(255) NOT NULL, 
    value FLOAT NOT NULL, 
    created_at DATE NOT NULL, 
    updated_at DATE NOT NULL
);  