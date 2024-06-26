-- name: CreateTableUsers :exec
CREATE TABLE users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		uname TEXT UNIQUE NOT NULL,
		passwd_hash TEXT NOT NULL,
		email TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		user_type TEXT NOT NULL,
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

-- name: CreateTableTransactions :exec
CREATE TABLE transactions (
		transaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
		uname TEXT NOT NULL,
		title TEXT NOT NULL,
		transaction_type TEXT NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		validity TEXT DEFAULT 'valid' NOT NULL
);

-- name: CreateTableReservation :exec
CREATE TABLE reservations (
		reservation_id INTEGER PRIMARY KEY AUTOINCREMENT,
		uname TEXT NOT NULL,
		title TEXT NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		validity TEXT DEFAULT 'valid' NOT NULL
);

-- name: CreateNewUser :exec
INSERT INTO users (
  uname, passwd_hash, email, first_name, last_name, user_type
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: RetrieveUserByUName :one
SELECT * FROM users
WHERE uname = ? AND validity = 'valid';

-- name: RetrievePsswdHash :one
SELECT passwd_hash FROM users
WHERE uname = ? AND validity = 'valid' AND user_type != 'approvalreq';

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

-- name: RetrieveBooksByISBN :many
SELECT * FROM books
WHERE isbn = ?;

-- name: RetrieveBooksByPopularity :many
SELECT * FROM books 
ORDER BY readers DESC;

-- name: RetrieveBooksByRating :many
SELECT * FROM books
ORDER BY rating DESC;

-- name: RetrieveBooksOfGenre :many
SELECT * FROM books
WHERE genre = ?
ORDER BY rating DESC;

-- name: CreateReservationForBook :exec
INSERT INTO reservations (
	uname, title
) VALUES (
	?, ?
);

-- name: UpdateBookQuantityDecrease :exec
UPDATE books
SET quantity = quantity - 1
WHERE book_id = ?;

-- name: RetrieveReservedBooks :many
SELECT reservation_id, reservations.title, author, rating, book_id
FROM reservations,books
WHERE reservations.uname = ? 
AND reservations.validity = 'valid' 
AND reservations.title = books.title ;

-- name: UpdateReservationValidity :exec
UPDATE reservations
SET validity = 'invalid'
WHERE reservation_id = ?;

-- name: UpdateBookQuantityIncrease :exec
UPDATE books
SET quantity = quantity + 1
WHERE title = ?;

-- name: RetrieveBookByBID :one
SELECT * FROM books
WHERE book_id = ?;

-- name: RetrieveLibrarian :one
SELECT * FROM users
WHERE uname = ? and user_type='librarian';

-- name: RetrieveUsersThatReqApprovalLike :many
SELECT * FROM users
WHERE validity='valid' and user_type='approvalreq' AND uname LIKE '%'|| ? || '%';

-- name: RetrieveMembersToRevokeLike :many
SELECT * FROM users
WHERE validity='valid' and user_type='member' AND uname LIKE '%'|| ? || '%';

-- name: RetrieveBookFromBID :one
SELECT * FROM books
WHERE book_id = ?;

-- name: RetrieveAvailibleBooks :many
SELECT * FROM books
WHERE quantity > 0 AND validity = 'valid' AND title LIKE '%'|| ? || '%';

-- name: CreateTransaction :exec
INSERT INTO transactions (
	uname, title, transaction_type
) VALUES (
	?, ?, ?
);

-- name: RetrieveReturnableBooksOfUser :many
SELECT * FROM transactions
WHERE validity = 'valid' AND transaction_type = 'issue' AND uname = ?;

-- name: UpdateTransactionValidity :exec
UPDATE transactions
SET validity = 'invalid'
WHERE transaction_id = ?;

-- name: InvalidateUser :exec
UPDATE users
SET validity = 'invalid'
WHERE uname = ?;