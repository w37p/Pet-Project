-- +goose Up
CREATE TABLE basket (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    menu_item_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_basket_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_basket_menu FOREIGN KEY (menu_item_id) REFERENCES menu(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE basket;
