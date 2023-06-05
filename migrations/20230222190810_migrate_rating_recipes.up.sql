CREATE TABLE IF NOT EXISTS rating_recipes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    recipe_id INT NOT NULL,
    level INT NOT NULL,
    
    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id),
    
    CONSTRAINT FK_recipe FOREIGN KEY(recipe_id)
    REFERENCES recipes(id)
);