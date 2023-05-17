CREATE TABLE IF NOT EXISTS reminders (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(100) NOT NULL,
    date TIMESTAMP NOT NULL,
    notification json NOT NULL,
    task json NOT NULL,
    note TEXT DEFAULT NULL,
    is_active BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id)
);
