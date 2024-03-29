CREATE TABLE IF NOT EXISTS medical_ratings (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    medical_id INT NOT NULL,
    reminder_id INT NOT NULL,
    rating INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_medical FOREIGN KEY(medical_id)
    REFERENCES medicals(id),

    CONSTRAINT FK_reminder FOREIGN KEY(reminder_id)
    REFERENCES reminders(id)
);