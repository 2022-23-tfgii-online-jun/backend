CREATE TABLE IF NOT EXISTS recipes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    category_id INT NOT NULL,
    user_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    image TEXT DEFAULT NULL,
    ingredients TEXT DEFAULT NULL,
    elaboration TEXT DEFAULT NULL,
    time TIMESTAMP DEFAULT NULL,
    rating CHAR(1) NOT NULL,
    is_published BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);