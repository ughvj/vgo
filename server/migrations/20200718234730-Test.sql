-- +migrate Up
CREATE TABLE IF NOT EXISTS test (
  id INTEGER        AUTO_INCREMENT NOT NULL,
  name VARCHAR(255)                NOT NULL,
  PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE IF EXISTS test;