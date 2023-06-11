CREATE TABLE IF NOT EXISTS monitorings (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    symptom_id INT NOT NULL,
    scale INT NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   
    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id),
    
    CONSTRAINT FK_symptom FOREIGN KEY(symptom_id)
    REFERENCES symptoms(id)
);