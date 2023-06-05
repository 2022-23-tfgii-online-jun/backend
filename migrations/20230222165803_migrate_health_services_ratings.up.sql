CREATE TABLE IF NOT EXISTS health_services_ratings (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    health_service_id INT NOT NULL,
    reminder_id INT NOT NULL,
    rating INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_health_service FOREIGN KEY(health_service_id)
    REFERENCES health_services(id),

    CONSTRAINT FK_reminder FOREIGN KEY(reminder_id)
    REFERENCES reminders(id)
);