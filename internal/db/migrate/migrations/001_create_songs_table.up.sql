CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    group_performer VARCHAR(100) NOT NULL,
    link VARCHAR(300),
    release_date DATE,
    verses_count INT NOT NULL
);
