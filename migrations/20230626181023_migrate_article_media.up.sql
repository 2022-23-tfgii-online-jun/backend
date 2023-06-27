CREATE TABLE IF NOT EXISTS article_media (
    id BIGSERIAL PRIMARY KEY,
    article_id INT NOT NULL,
    media_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_article_media FOREIGN KEY(article_id)
    REFERENCES articles(id),

    CONSTRAINT FK_media_reminder FOREIGN KEY (media_id)
    REFERENCES media(id)
);