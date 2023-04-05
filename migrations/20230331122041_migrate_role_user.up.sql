CREATE TABLE IF NOT EXISTS role_user (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id),
    
    CONSTRAINT FK_role FOREIGN KEY(role_id)
    REFERENCES roles(id)
);