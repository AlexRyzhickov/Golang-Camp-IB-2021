CREATE TABLE notes (
    service VARCHAR,
    mode BOOLEAN,
    PRIMARY KEY(service)
);

INSERT INTO notes (service, mode) VALUES ('portal', FALSE);
INSERT INTO notes (service, mode) VALUES ('responder', TRUE);
INSERT INTO notes (service, mode) VALUES ('storage', TRUE);
