CREATE TABLE IF NOT EXISTS treatments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),  
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    frequency json NOT NULL,
    shots json NOT NULL,
    notes VARCHAR(255) NOT NULL,
    date_start TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id)
);