CREATE TABLE IF NOT EXISTS verses (
    id SERIAL PRIMARY KEY,
    song_id INT NOT NULL,
    num INT NOT NULL,
    lyrics TEXT NOT NULL,

    FOREIGN KEY (song_id) REFERENCES songs(id) ON DELETE CASCADE
);
