CREATE TABLE IF NOT EXISTS activity_users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    activity_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id),

    CONSTRAINT FK_activity FOREIGN KEY(activity_id)
    REFERENCES activities(id)
);




