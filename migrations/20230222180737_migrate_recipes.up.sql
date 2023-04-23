CREATE TABLE IF NOT EXISTS recipes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(200) NOT NULL,
    image TEXT DEFAULT NULL,
    ingredients json NOT NULL,
    elaboration json NOT NULL,
    prep_time INT DEFAULT NULL,
    cook_time INT DEFAULT NULL,
    serving INT DEFAULT NULL,
    difficulty INT DEFAULT NULL,
    nutrition json DEFAULT NULL,
    is_published BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);