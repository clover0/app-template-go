-- create users table
CREATE TABLE users
(
  id         bigint PRIMARY KEY          not null,
  updated_at timestamp without time zone not null,
  created_at timestamp without time zone not null,
  email      varchar(512)                not null,
  password   varchar(512)                not null
)