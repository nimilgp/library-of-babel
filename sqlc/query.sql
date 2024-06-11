-- name: CreateNewUser :exec
INSERT INTO users (
  uname, passwd_hash, email, first_name, last_name
) VALUES (
  ?, ?, ?, ?, ?
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