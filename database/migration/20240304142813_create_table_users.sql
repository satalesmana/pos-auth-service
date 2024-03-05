-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  email varchar(255) COLLATE utf8_unicode_ci NOT NULL UNIQUE,
  password varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
