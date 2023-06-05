CREATE TABLE media (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    recipe_id INTEGER,
    media_type VARCHAR(255) NOT NULL,
    media_url TEXT NOT NULL,
    media_thumb TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_recipe FOREIGN KEY(recipe_id)
    REFERENCES recipes(id)
);
