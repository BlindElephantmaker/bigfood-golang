CREATE TABLE user_token
(
    refresh_token uuid      PRIMARY KEY,
    user_id       uuid      NOT NULL,
    expires_at    TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE
);