
-- +migrate Up
CREATE TABLE IF NOT EXISTS questions (
    id serial PRIMARY KEY,
    content varchar(100) NOT NULL,
    is_answerd boolean DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS answers (
    id serial PRIMARY KEY,
    content varchar(100) NOT NULL,
    question_id serial REFERENCES questions(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS question;
DROP TABLE IF EXISTS answers;