CREATE TABLE IF NOT EXISTS roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  display_name VARCHAR(255) NOT NULL
);

INSERT INTO roles (name, display_name) VALUES ("admin", "Admin");
INSERT INTO roles (name, display_name) VALUES ("user", "User");
