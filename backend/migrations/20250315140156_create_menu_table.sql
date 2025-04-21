-- +goose Up
CREATE TABLE menu (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    category_id INTEGER NOT NULL,
    image_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_menu_category
      FOREIGN KEY (category_id)
      REFERENCES categories(id)
      ON DELETE RESTRICT
      ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE menu;
