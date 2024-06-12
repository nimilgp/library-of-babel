// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package dbLayer

import (
	"context"
	"database/sql"
)

const createNewBook = `-- name: CreateNewBook :exec
INSERT INTO books (
  title, author, year, genre, isbn, rating, readers, quantity
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateNewBookParams struct {
	Title    string
	Author   string
	Year     int64
	Genre    string
	Isbn     string
	Rating   float64
	Readers  int64
	Quantity int64
}

func (q *Queries) CreateNewBook(ctx context.Context, arg CreateNewBookParams) error {
	_, err := q.db.ExecContext(ctx, createNewBook,
		arg.Title,
		arg.Author,
		arg.Year,
		arg.Genre,
		arg.Isbn,
		arg.Rating,
		arg.Readers,
		arg.Quantity,
	)
	return err
}

const createNewUser = `-- name: CreateNewUser :exec
INSERT INTO users (
  uname, passwd_hash, email, first_name, last_name, user_type
) VALUES (
  ?, ?, ?, ?, ?, ?
)
`

type CreateNewUserParams struct {
	Uname      string
	PasswdHash string
	Email      string
	FirstName  string
	LastName   string
	UserType   string
}

func (q *Queries) CreateNewUser(ctx context.Context, arg CreateNewUserParams) error {
	_, err := q.db.ExecContext(ctx, createNewUser,
		arg.Uname,
		arg.PasswdHash,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.UserType,
	)
	return err
}

const createTableBooks = `-- name: CreateTableBooks :exec
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
)
`

func (q *Queries) CreateTableBooks(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createTableBooks)
	return err
}

const createTableUsers = `-- name: CreateTableUsers :exec
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
)
`

func (q *Queries) CreateTableUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createTableUsers)
	return err
}

const retrieveAllBooks = `-- name: RetrieveAllBooks :many
SELECT book_id, title, author, year, genre, isbn, rating, readers, quantity, sqltime, validity FROM books
`

func (q *Queries) RetrieveAllBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, retrieveAllBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Genre,
			&i.Isbn,
			&i.Rating,
			&i.Readers,
			&i.Quantity,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveBooksByAuthorValue = `-- name: RetrieveBooksByAuthorValue :many
SELECT book_id, title, author, year, genre, isbn, rating, readers, quantity, sqltime, validity FROM books
WHERE author LIKE '%'|| ? ||'%'
`

func (q *Queries) RetrieveBooksByAuthorValue(ctx context.Context, dollar_1 sql.NullString) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, retrieveBooksByAuthorValue, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Genre,
			&i.Isbn,
			&i.Rating,
			&i.Readers,
			&i.Quantity,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveBooksByISBN = `-- name: RetrieveBooksByISBN :many
SELECT book_id, title, author, year, genre, isbn, rating, readers, quantity, sqltime, validity FROM books
WHERE isbn = ?
`

func (q *Queries) RetrieveBooksByISBN(ctx context.Context, isbn string) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, retrieveBooksByISBN, isbn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Genre,
			&i.Isbn,
			&i.Rating,
			&i.Readers,
			&i.Quantity,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveBooksByPopularity = `-- name: RetrieveBooksByPopularity :many
SELECT book_id, title, author, year, genre, isbn, rating, readers, quantity, sqltime, validity FROM books 
ORDER BY readers DESC
`

func (q *Queries) RetrieveBooksByPopularity(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, retrieveBooksByPopularity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Genre,
			&i.Isbn,
			&i.Rating,
			&i.Readers,
			&i.Quantity,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveBooksByRating = `-- name: RetrieveBooksByRating :many
SELECT book_id, title, author, year, genre, isbn, rating, readers, quantity, sqltime, validity FROM books
ORDER BY rating DESC
`

func (q *Queries) RetrieveBooksByRating(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, retrieveBooksByRating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Genre,
			&i.Isbn,
			&i.Rating,
			&i.Readers,
			&i.Quantity,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveBooksByTitleValue = `-- name: RetrieveBooksByTitleValue :many
SELECT book_id, title, author, year, genre, isbn, rating, readers, quantity, sqltime, validity FROM books
WHERE title LIKE '%'|| ? ||'%'
`

func (q *Queries) RetrieveBooksByTitleValue(ctx context.Context, dollar_1 sql.NullString) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, retrieveBooksByTitleValue, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Genre,
			&i.Isbn,
			&i.Rating,
			&i.Readers,
			&i.Quantity,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveBooksOfGenre = `-- name: RetrieveBooksOfGenre :many
SELECT book_id, title, author, year, genre, isbn, rating, readers, quantity, sqltime, validity FROM books
WHERE genre = ?
ORDER BY rating DESC
`

func (q *Queries) RetrieveBooksOfGenre(ctx context.Context, genre string) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, retrieveBooksOfGenre, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Author,
			&i.Year,
			&i.Genre,
			&i.Isbn,
			&i.Rating,
			&i.Readers,
			&i.Quantity,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveUserByUID = `-- name: RetrieveUserByUID :one
SELECT user_id, uname, passwd_hash, email, first_name, last_name, user_type, actions_left, sqltime, validity FROM users
WHERE user_id = ? AND validity = 'valid'
`

func (q *Queries) RetrieveUserByUID(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRowContext(ctx, retrieveUserByUID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Uname,
		&i.PasswdHash,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.UserType,
		&i.ActionsLeft,
		&i.Sqltime,
		&i.Validity,
	)
	return i, err
}

const retrieveUsersByUType = `-- name: RetrieveUsersByUType :many
SELECT user_id, uname, passwd_hash, email, first_name, last_name, user_type, actions_left, sqltime, validity FROM users
WHERE user_type = ? AND validity = 'valid'
`

func (q *Queries) RetrieveUsersByUType(ctx context.Context, userType string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, retrieveUsersByUType, userType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Uname,
			&i.PasswdHash,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.UserType,
			&i.ActionsLeft,
			&i.Sqltime,
			&i.Validity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserType = `-- name: UpdateUserType :exec
UPDATE users
SET user_type = ?
WHERE user_id = ? AND validity = 'valid'
`

type UpdateUserTypeParams struct {
	UserType string
	UserID   int64
}

func (q *Queries) UpdateUserType(ctx context.Context, arg UpdateUserTypeParams) error {
	_, err := q.db.ExecContext(ctx, updateUserType, arg.UserType, arg.UserID)
	return err
}

const updateUserValidity = `-- name: UpdateUserValidity :exec
UPDATE users
SET validity = 'invalid'
WHERE user_id = ?
`

func (q *Queries) UpdateUserValidity(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, updateUserValidity, userID)
	return err
}
