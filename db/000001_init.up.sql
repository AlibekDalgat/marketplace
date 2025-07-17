CREATE TABLE users
(
    login VARCHAR(20) PRIMARY KEY,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE advertisements
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    text VARCHAR(1000),
    image VARCHAR(2000),
    price REAL NOT NULL,
    posting_date TIMESTAMP NOT NULL ,
    owner VARCHAR(20) REFERENCES users(login) ON DELETE CASCADE NOT NULL
);