CREATE TABLE IF NOT EXISTS user_has_roles (
  user_id BIGINT UNSIGNED,
  role_id BIGINT UNSIGNED,

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (role_id) REFERENCES roles(id)
)
