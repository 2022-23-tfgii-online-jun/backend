CREATE TABLE IF NOT EXISTS activities (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    content TEXT DEFAULT NULL,
    image TEXT DEFAULT NULL,
    place VARCHAR(255) NOT NULL,
    host VARCHAR(255) NOT NULL,
    date TIMESTAMP NOT NULL,
    duration INTEGER NOT NULL,
    is_published BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 );




