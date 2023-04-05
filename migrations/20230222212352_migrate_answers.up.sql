CREATE TABLE IF NOT EXISTS answers (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id INT NOT NULL,
    question_id INT NOT NULL,
    response TEXT DEFAULT NULL,
    is_public BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id),
    
    CONSTRAINT FK_question FOREIGN KEY(question_id)
    REFERENCES questions(id)
);
