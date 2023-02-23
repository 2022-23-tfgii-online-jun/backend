CREATE TABLE IF NOT EXISTS health_services_ratings (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    health_services_id INT NOT NULL,
    rating CHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_health_service FOREIGN KEY(health_services_id)
    REFERENCES health_services(id)
);