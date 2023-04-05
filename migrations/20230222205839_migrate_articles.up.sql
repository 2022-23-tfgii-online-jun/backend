CREATE TABLE IF NOT EXISTS articles (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    image TEXT DEFAULT NULL,
    content TEXT DEFAULT NULL,
    is_published BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id)
);