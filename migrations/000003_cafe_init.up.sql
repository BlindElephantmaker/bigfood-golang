CREATE TABLE cafe
(
    id         uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE cafe_user
(
    id         uuid PRIMARY KEY,
    cafe_id    uuid        NOT NULL,
    user_id    uuid        NOT NULL,
    comment    VARCHAR(32) NOT NULL,
    created_at TIMESTAMP   NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (cafe_id) REFERENCES cafe (id) ON DELETE CASCADE,
    UNIQUE (cafe_id, user_id)
);

CREATE TABLE cafe_user_role
(
    cafe_user_id uuid        NOT NULL,
    role         VARCHAR(32) NOT NULL,
    FOREIGN KEY (cafe_user_id) REFERENCES cafe_user (id) ON DELETE CASCADE,
    UNIQUE (cafe_user_id, role)
);