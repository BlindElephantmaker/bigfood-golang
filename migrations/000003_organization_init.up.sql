CREATE TABLE organization
(
    id         uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE role
(
    name VARCHAR(32) PRIMARY KEY
);

CREATE TABLE user_organization
(
    id              uuid PRIMARY KEY,
    user_id         uuid      NOT NULL,
    organization_id uuid      NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (organization_id) REFERENCES organization (id) ON DELETE CASCADE,
    UNIQUE (user_id, organization_id)
);

CREATE TABLE user_organization_role
(
    user_organization_id uuid NOT NULL,
    role_reference VARCHAR(32) NOT NULL,
    FOREIGN KEY (user_organization_id) REFERENCES user_organization (id) ON DELETE CASCADE,
    FOREIGN KEY (role_reference) REFERENCES role (name) ON DELETE CASCADE,
    UNIQUE (user_organization_id, role_reference)
);
