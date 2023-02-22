CREATE TABLE IF NOT EXISTS reminders (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    file_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(100) NOT NULL,
    date TIMESTAMP NOT NULL,
    notifications json NOT NULL,
    tasks json NOT NULL,
    notes TEXT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id),
    
    CONSTRAINT FK_file FOREIGN KEY(file_id)
    REFERENCES files(id)
);