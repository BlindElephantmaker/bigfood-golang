CREATE TABLE contact
(
    id         uuid PRIMARY KEY,
    type       VARCHAR(32) NOT NULL,
    created_at TIMESTAMP   NOT NULL
);

CREATE TABLE contact_phone
(
    contact_id uuid PRIMARY KEY,
    phone      VARCHAR(12) NOT NULL UNIQUE,
    FOREIGN KEY (contact_id) REFERENCES contact (id) ON DELETE CASCADE
);

CREATE TABLE reserve
(
    id          uuid PRIMARY KEY,
    table_id    uuid      NOT NULL,
    contact_id  uuid      NOT NULL,
    comment     TEXT      NOT NULL,
    guest_count INTEGER   NOT NULL,
    from_date   TIMESTAMP NOT NULL, -- todo: index it?
    until_date  TIMESTAMP NOT NULL, -- todo: index it?
    created_at  TIMESTAMP NOT NULL,
    deleted_at  TIMESTAMP,
    FOREIGN KEY (table_id) REFERENCES tables (id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contact (id) ON DELETE CASCADE
);
