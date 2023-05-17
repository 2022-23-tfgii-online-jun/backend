CREATE TABLE IF NOT EXISTS reminder_media (
    id BIGSERIAL PRIMARY KEY,
    reminder_id INT NOT NULL,
    media_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_reminder_media FOREIGN KEY(reminder_id)
    REFERENCES reminders(id),

    CONSTRAINT FK_media_reminder FOREIGN KEY (media_id)
    REFERENCES media(id)
);