CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) NOT NULL,
    salt VARCHAR(100) NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE messages (
    messageID SERIAL PRIMARY KEY,
    author VARCHAR(100) REFERENCES users(username),
    text TEXT,
    createdAt TIMESTAMP,
    flagged BOOLEAN
);

CREATE TABLE follows (
    follower INTEGER REFERENCES users(id),
    following INTEGER REFERENCES users(id),
    PRIMARY KEY (follower, following)
);
