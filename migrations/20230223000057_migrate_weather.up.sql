CREATE TABLE IF NOT EXISTS weather (
    country CHAR(3) NOT NULL,
    state VARCHAR(100) NOT NULL,
    temp_c VARCHAR(20) NOT NULL,
    description VARCHAR(100) NOT NULL,
    humidity VARCHAR(5) NOT NULL,
    code VARCHAR(5) NOT NULL, 
    wind VARCHAR(5) NOT NULL,     
    uv VARCHAR(5) NOT NULL,     
    alert VARCHAR(100) NOT NULL,     
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
