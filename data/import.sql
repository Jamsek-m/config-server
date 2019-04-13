INSERT INTO roles(code, label) VALUES ('ADMIN', 'Admin');
INSERT INTO roles(code, label) VALUES ('CONF_READ', 'Read only');
INSERT INTO roles(code, label) VALUES ('CONF_WRITE', 'Read and write');

-- geslo123, token
INSERT INTO users(username, password, access_token) VALUES ('miha', '$2a$10$DgFGdDX31W9QocuGhyXWqOdY.aVzaXjU9aw71jv4KAGqOjSs7xqiS', 'token');
INSERT INTO users(username, password, access_token) VALUES ('readonly', '$2a$10$DgFGdDX31W9QocuGhyXWqOdY.aVzaXjU9aw71jv4KAGqOjSs7xqiS', 'token');
INSERT INTO users(username, password, access_token) VALUES ('readandwrite', '$2a$10$DgFGdDX31W9QocuGhyXWqOdY.aVzaXjU9aw71jv4KAGqOjSs7xqiS', 'token');

INSERT INTO user_roles(user_id, role_id) VALUES (1, 1);
INSERT INTO user_roles(user_id, role_id) VALUES (1, 2);
INSERT INTO user_roles(user_id, role_id) VALUES (1, 3);
INSERT INTO user_roles(user_id, role_id) VALUES (2, 2);
INSERT INTO user_roles(user_id, role_id) VALUES (3, 2);
INSERT INTO user_roles(user_id, role_id) VALUES (3, 3);

INSERT INTO configuration(config_key, config_value) VALUES ('/service1/1.0.0/dev/email', 'test@mail.com');
INSERT INTO configuration(config_key, config_value) VALUES ('/service1/2.0.0/dev/email', 'test@mail.com');
INSERT INTO configuration(config_key, config_value) VALUES ('/service1/2.0.0/prod/email', 'test@mail.com');
INSERT INTO configuration(config_key, config_value) VALUES ('/service2/1.0.0/dev/email', 'test@mail.com');
INSERT INTO configuration(config_key, config_value) VALUES ('/service2/instances/df32443hgdsfvb3q', 'http://localhost:8080/v3');
INSERT INTO configuration(config_key, config_value) VALUES ('/service2/instances/po6kjdjgffgas1324', 'http://api.test.com/v3');
INSERT INTO configuration(config_key, config_value) VALUES ('/service3/val1', 'value one');
INSERT INTO configuration(config_key, config_value) VALUES ('/service3/val2/val2.1', 'value two point one');
INSERT INTO configuration(config_key, config_value) VALUES ('/service3/val2/val2.2', 'value two point two');
INSERT INTO configuration(config_key, config_value) VALUES ('/service3/very/very/very/very/very/ok/neda/se/mi/vec/se/enkrat/very/long/key', 'vrednost');

