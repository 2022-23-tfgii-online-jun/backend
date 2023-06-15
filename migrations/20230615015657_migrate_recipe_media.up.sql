CREATE TABLE IF NOT EXISTS recipe_media (
    id BIGSERIAL PRIMARY KEY,
    recipe_id INT NOT NULL,
    media_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_recipe_media FOREIGN KEY(recipe_id)
    REFERENCES recipes(id),

    CONSTRAINT FK_media_reminder FOREIGN KEY (media_id)
    REFERENCES media(id)
);