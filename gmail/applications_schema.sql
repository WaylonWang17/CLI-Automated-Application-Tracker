CREATE TABLE IF NOT EXISTS applications(
    id SERIAL PRIMARY KEY,
    subject TEXT,
    sender TEXT,
    date TEXT,
    snippet TEXT
);