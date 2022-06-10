CREATE TABLE chat
(
    id         UUID PRIMARY KEY,
    cafe_id    uuid      NOT NULL,
    contact_id uuid      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (cafe_id) REFERENCES cafe (id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contact (id) ON DELETE CASCADE
);

CREATE TABLE message
(
    id        uuid PRIMARY KEY,
    chat_id   uuid      NOT NULL,
    text      text      NOT NULL,
    crated_at TIMESTAMP NOT NULL,
    is_read   bool      NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES chat (id) ON DELETE CASCADE
);