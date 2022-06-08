CREATE TABLE cafe_contact
(
    id            uuid          PRIMARY KEY,
    contact_id    uuid          NOT NULL,
    cafe_id       uuid          NOT NULL,
    created_at    TIMESTAMP     NOT NULL,
    FOREIGN KEY (cafe_id) REFERENCES cafe (id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contact (id) ON DELETE CASCADE
);

CREATE TABLE messages
(
    id                  uuid          PRIMARY KEY,
    cafe_contact_id     uuid          NOT NULL,
    text                text          NOT NULL,
    crated_at           TIMESTAMP     NOT NULL,
    is_read             bool         NOT NULL,
    FOREIGN KEY (cafe_contact_id) references cafe_contact (id) ON DELETE CASCADE
);