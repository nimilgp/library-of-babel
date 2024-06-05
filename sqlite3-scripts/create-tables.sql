CREATE TABLE users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		uname TEXT UNIQUE NOT NULL,
		passwd_hash TEXT NOT NULL,
		email TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		user_type TEXT 	DEFAULT 'approvalreq',
		actions_left INTEGER DEFAULT '5',
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		validity TEXT DEFAULT 'valid'
);


CREATE TABLE books (
		book_id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT UNIQUE NOT NULL,
		author TEXT NOT NULL,
		year INTEGER NOT NULL,
		genre TEXT NOT NULL,
		isbn INTEGER NOT NULL,
		rating REAL DEFAULT '0',
		readers INTEGER DEFAULT '0',
		quantity INTEGER DEFAULT '1',
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		validity TEXT DEFAULT 'valid'
);

CREATE TABLE transactions (
		transaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		book_id INTEGER NOT NULL,
		transaction_type TEXT NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		validity TEXT DEFAULT 'valid'
);

CREATE TABLE reservations (
		reservation_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		book_id INTEGER NOT NULL,
		sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		validity TEXT DEFAULT 'valid'
);
