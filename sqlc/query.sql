-- name: CreateTableUsers :exec
CREATE TABLE users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		uname TEXT UNIQUE NOT NULL,
		passwd_hash TEXT NOT NULL,
		email TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		user_type TEXT  NOT NULL,
		actions_left INTEGER DEFAULT '5' NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		validity TEXT DEFAULT 'valid' NOT NULL
);

-- name: CreateTableBooks :exec
CREATE TABLE books (
		book_id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT UNIQUE NOT NULL,
		author TEXT NOT NULL,
		year INTEGER NOT NULL,
		genre TEXT NOT NULL,
		isbn TEXT NOT NULL,
		rating REAL DEFAULT '0' NOT NULL,
		readers INTEGER DEFAULT '0' NOT NULL,
		quantity INTEGER DEFAULT '1' NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		validity TEXT DEFAULT 'valid' NOT NULL
);

-- name: CreateNewUser :exec
INSERT INTO users (
  uname, passwd_hash, email, first_name, last_name, user_type
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: RetrieveUserByUID :one
SELECT * FROM users
WHERE user_id = ? AND validity = 'valid';

-- name: RetrieveUsersByUType :many
SELECT * FROM users
WHERE user_type = ? AND validity = 'valid';

-- name: UpdateUserType :exec
UPDATE users
SET user_type = ?
WHERE user_id = ? AND validity = 'valid';

-- name: UpdateUserValidity :exec
UPDATE users
SET validity = 'invalid'
WHERE user_id = ?;

-- name: CreateNewBook :exec
INSERT INTO books (
  title, author, year, genre, isbn, rating, readers, quantity
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: RetrieveAllBooks :many
SELECT * FROM books;

-- name: RetrieveBooksByTitleValue :many
SELECT * FROM books
WHERE title LIKE '%'|| ? ||'%';

-- name: RetrieveBooksByAuthorValue :many
SELECT * FROM books
WHERE author LIKE '%'|| ? ||'%';

-- name: RetrieveBooksByISBN :one
SELECT * FROM books
WHERE isbn = ?;