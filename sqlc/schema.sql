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

CREATE TABLE transactions (
		transaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
		uname TEXT NOT NULL,
		title TEXT NOT NULL,
		transaction_type TEXT NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		validity TEXT DEFAULT 'valid' NOT NULL
);

CREATE TABLE reservations (
		reservation_id INTEGER PRIMARY KEY AUTOINCREMENT,
		uname TEXT NOT NULL,
		title TEXT NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		validity TEXT DEFAULT 'valid' NOT NULL
);
