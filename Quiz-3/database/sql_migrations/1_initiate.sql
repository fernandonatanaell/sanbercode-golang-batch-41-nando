-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    image_url VARCHAR(256) NOT NULL,
    release_year INT NOT NULL,
    price VARCHAR(256) NOT NULL,
    total_page INT NOT NULL,
    thickness VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    category_id INTEGER NOT NULL,
    FOREIGN KEY (category_id) REFERENCES category(id)
);

-- +migrate StatementEnd