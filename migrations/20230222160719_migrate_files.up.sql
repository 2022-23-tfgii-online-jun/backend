CREATE TABLE IF NOT EXISTS reminders (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    reminder_id INT NOT NULL,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id),

    CONSTRAINT FK_reminder FOREIGN KEY(reminder_id)
    REFERENCES reminders(id)
);