CREATE TABLE IF NOT EXISTS article_categories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    category_id INT NOT NULL,
    article_id INT NOT NULL,
    image TEXT DEFAULT NULL,
    content TEXT DEFAULT NULL,
    is_published BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT FK_categorie FOREIGN KEY(category_id)
    REFERENCES categories(id),
    
    CONSTRAINT FK_article FOREIGN KEY(article_id)
    REFERENCES articles(id)
);




