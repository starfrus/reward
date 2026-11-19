CREATE TABLE users (
    id                      SERIAL       PRIMARY KEY,
    version                 INTEGER      NOT NULL     DEFAULT 1,
    login                   VARCHAR(30)  NOT NULL     UNIQUE      CHECK (login ~ '^[A-Za-z][A-Za-z0-9_]{2,29}$'),
    password_hash           VARCHAR(255) NOT NULL,              
    completed_targets_count INTEGER      NOT NULL     DEFAULT 0   CHECK(completed_targets_count >= 0),
    created_at              TIMESTAMPTZ  NOT NULL     DEFAULT now()
)