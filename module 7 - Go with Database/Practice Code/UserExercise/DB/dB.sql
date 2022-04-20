CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL,
  zipCode TEXT,
  city TEXT,
  addressU TEXT,
  stateU TEXT,
  passwordU TEXT
)