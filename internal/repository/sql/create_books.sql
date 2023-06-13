CREATE TABLE IF NOT EXISTS books (
    book_id INTEGER NOT NULL  PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    author_id INTEGER,
    descriptions TEXT,
    FOREIGN KEY(author_id) REFERENCES AUTHORS(author_id))
