CREATE TABLE users
(
    id    uuid        PRIMARY KEY,
    name  VARCHAR(64) NOT NULL,
    phone VARCHAR(12) NOT NULL UNIQUE
);