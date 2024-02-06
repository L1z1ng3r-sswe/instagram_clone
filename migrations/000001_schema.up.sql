CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  avatar TEXT NOT NULL
);

CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  description TEXT NOT NULL,
  main_image TEXT NOT NULL,
  images TEXT[] NOT NULL,
  created_by BIGINT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);
