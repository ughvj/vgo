-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id INTEGER        AUTO_INCREMENT NOT NULL,
  name VARCHAR(255)                NOT NULL,
  pass VARCHAR(255)                NOT NULL,
  administrator BOOLEAN            NOT NULL,
  token VARCHAR(1024)                      ,
  token_expire_date DATETIME               ,
  PRIMARY KEY(id)
);

INSERT INTO users(
  name, pass, administrator, token, token_expire_date
) VALUES(
  'user1', 'password', 0, NULL, NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;