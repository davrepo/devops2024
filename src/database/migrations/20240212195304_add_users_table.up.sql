CREATE TABLE "user" (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(100) UNIQUE NOT NULL,
  email VARCHAR(100) NOT NULL,
  salt VARCHAR(100) NOT NULL,
  pw_hash TEXT NOT NULL
);

CREATE TABLE message (
  message_id SERIAL PRIMARY KEY,
  author_id INTEGER REFERENCES "user"(user_id),
  text TEXT NOT NULL,
  pub_date INTEGER NOT NULL,
  flagged INTEGER DEFAULT 0
);

CREATE TABLE follower (
  who_id INTEGER REFERENCES "user"(user_id),
  whom_id INTEGER REFERENCES "user"(user_id),
  PRIMARY KEY (who_id, whom_id)
);