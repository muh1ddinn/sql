CREATE TABLE IF NOT EXISTS countryy(
    id uuid PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    code int
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);