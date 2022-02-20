CREATE TABLE organization
(
    id         uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE organization_role
(
    name VARCHAR(32) PRIMARY KEY
);

CREATE TABLE organization_user
(
    id              uuid PRIMARY KEY,
    organization_id uuid      NOT NULL,
    user_id         uuid      NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (organization_id) REFERENCES organization (id) ON DELETE CASCADE,
    UNIQUE (organization_id, user_id)
);

CREATE TABLE organization_user_role
(
    organization_user_id uuid        NOT NULL,
    role                 VARCHAR(32) NOT NULL,
    FOREIGN KEY (organization_user_id) REFERENCES organization_user (id) ON DELETE CASCADE,
    FOREIGN KEY (role) REFERENCES organization_role (name) ON DELETE CASCADE,
    UNIQUE (organization_user_id, role)
);

INSERT INTO organization_role (name)
VALUES ('owner'),
       ('admin'),
       ('hostess');