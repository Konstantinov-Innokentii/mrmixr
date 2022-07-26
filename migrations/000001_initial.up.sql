create table github_installation
(
    id                SERIAL PRIMARY KEY,
    installation_id   INTEGER                  NOT NULL,
    installed_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    installation_type INTEGER                  NOT NULL
);

create table github_repository
(
    id                     SERIAL PRIMARY KEY,
    github_id              INTEGER                                     NOT NULL,
    name                   VARCHAR(128)                                NOT NULL,
    github_installation_id INTEGER REFERENCES github_installation (id) NOT NULL
);

create table pr_check
(
    id     SERIAL PRIMARY KEY,
    hour   INTEGER      NOT NULL,
    minute INTEGER      NOT NULL,
    days   INTEGER[]    NOT NULL,
    tz     VARCHAR(128) NOT NULL,
    name   VARCHAR(128) NOT NULL
);


create table check_has_repository
(
    id                   SERIAL PRIMARY KEY,
    check_id             INTEGER REFERENCES pr_check (id)          NOT NULL,
    github_repository_id INTEGER REFERENCES github_repository (id) NOT NULL,
    UNIQUE (check_id, github_repository_id)
);
