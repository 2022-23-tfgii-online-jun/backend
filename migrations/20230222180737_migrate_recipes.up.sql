CREATE TABLE IF NOT EXISTS recipes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    category_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    ingredients TEXT DEFAULT NULL,
    elaboration TEXT DEFAULT NULL,
    time INT DEFAULT NULL,
    is_published BOOLEAN DEFAULT false,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
