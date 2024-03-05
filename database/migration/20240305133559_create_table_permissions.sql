-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  users_id int(10) COLLATE utf8_unicode_ci NOT NULL,
  activity_id varchar(36) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permissions;
-- +goose StatementEnd
