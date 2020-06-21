CREATE DATABASE IF NOT EXISTS sample;

use sample;

DROP TABLE IF EXISTS todos;

CREATE TABLE todos (
    id INTEGER NOT NULL AUTO_INCREMENT,
    title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO
    todos (title)
VALUES
    ('hoge');

INSERT INTO
    todos (title)
VALUES
    ('fuga');

INSERT INTO
    todos (title)
VALUES
    ('foo');

INSERT INTO
    todos (title)
VALUES
    ('bar');