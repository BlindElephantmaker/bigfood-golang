CREATE TABLE tables
(
    id         uuid PRIMARY KEY,
    cafe_id    uuid        NOT NULL,
    title      VARCHAR(32) NOT NULL,
    comment    VARCHAR(32) NOT NULL,
    seats      INTEGER     NOT NULL,
    created_at TIMESTAMP   NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (cafe_id) REFERENCES cafe (id) ON DELETE CASCADE
);