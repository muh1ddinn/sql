CREATE TABLE IF NOT EXISTS product(
    id uuid ,
    name VARCHAR(50) UNIQUE,
	category_id uuid PRIMARY KEY REFERENCES category(id),
	Price INTEGER ,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS category(
    id uuid PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);
