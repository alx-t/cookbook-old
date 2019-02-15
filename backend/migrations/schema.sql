// psql -h localhost -p 5432 -U postgres
// \c cookbook

CREATE DATABASE cookbook;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  nick_name VARCHAR (50) NOT NULL,
  email VARCHAR (50) UNIQUE NOT NULL,
  data JSONB DEFAULT '{}'::jsonb,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE recipes (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  title VARCHAR (50) DEFAULT NULL,
  description TEXT DEFAULT NULL,
  data JSONB DEFAULT '{}'::jsonb,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO users (nick_name, email) VALUES
  ('test user 1', 'test@test.com');

INSERT INTO recipes (user_id, title, description) VALUES
  (1, 'test recipe 01', 'test recipe 01 description'),
  (1, 'test recipe 02', 'test recipe 02 description'),
  (1, 'test recipe 03', 'test recipe 03 description'),
  (1, 'test recipe 04', 'test recipe 04 description'),
  (1, 'test recipe 05', 'test recipe 05 description');
