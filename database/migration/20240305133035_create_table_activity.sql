-- +goose Up
-- +goose StatementBegin
CREATE TABLE activities (
  id varchar(36) COLLATE utf8_unicode_ci NOT NULL UNIQUE,
  parent varchar(36) COLLATE utf8_unicode_ci NULL,
  label varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  path_url varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE activities;
-- +goose StatementEnd
