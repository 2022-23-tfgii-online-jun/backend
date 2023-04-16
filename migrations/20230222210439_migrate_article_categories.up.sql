CREATE TABLE IF NOT EXISTS article_category (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    article_id INT NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT FK_categorie FOREIGN KEY(category_id)
    REFERENCES categories(id),
    
    CONSTRAINT FK_article FOREIGN KEY(article_id)
    REFERENCES articles(id)
);




