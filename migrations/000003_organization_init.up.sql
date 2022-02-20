CREATE TABLE organization
(
    id         uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE users_organizations
(
    user_id         uuid      NOT NULL,
    organization_id uuid      NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE,
    FOREIGN KEY (organization_id) REFERENCES Organization (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, organization_id)
);