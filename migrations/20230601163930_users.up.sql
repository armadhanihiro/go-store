CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
	phone_number VARCHAR(50),
	address TEXT,
	city VARCHAR(50),
	province VARCHAR(50),
	country VARCHAR(50),
	postal_code VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
