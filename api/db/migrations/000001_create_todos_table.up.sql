CREATE TABLE IF NOT EXISTS todos (
  id varchar(26) not null PRIMARY KEY,
  title varchar(255) not null,
  body varchar(255) not null
);