CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    group_performer VARCHAR(100) NOT NULL,
    link VARCHAR(300),
    lyrics TEXT,
    release_date DATE
);
