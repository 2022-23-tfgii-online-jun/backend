CREATE TABLE IF NOT EXISTS symptoms (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(), 
    user_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL,
    scale CHAR(5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id)
);