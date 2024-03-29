CREATE TABLE IF NOT EXISTS questions (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id INT NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id)
);
