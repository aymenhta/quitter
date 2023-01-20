DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;


CREATE TABLE users(
	id 			SERIAL 			PRIMARY KEY,
	username 	VARCHAR(255) 	UNIQUE NOT NULL,
	email 		VARCHAR(255) 	UNIQUE NOT NULL,
	password 	CHAR(60) 		NOT NULL,
	joined_at 	DATE
);

CREATE INDEX idx_users_id ON users(id);

CREATE TABLE posts (
    id          SERIAL  PRIMARY KEY,
    content     TEXT    NOT NULL,
    posted_at   TIME,	
    user_id 	SERIAL,
	CONSTRAINT fk_user
      	FOREIGN KEY(user_id) 
		REFERENCES users(id)
		ON DELETE CASCADE
);

CREATE INDEX idx_posts_id ON posts(id);

INSERT iNTO users (username, email, password, joined_at) 
	VALUES ('aymengk94', 'aymengk94@gmail.com', 'test123', CURRENT_DATE);
	
	
SELECT * FROM users;

INSERT INTO posts (content, posted_at, user_id) 
	VALUES('This is my first post here!!', CURRENT_TIME, 1);
INSERT INTO posts (content, posted_at, user_id) 
	VALUES('This is my second post here!!', CURRENT_TIME, 1);
INSERT INTO posts (content, posted_at, user_id) 
	VALUES('This is my third post here!!', CURRENT_TIME, 1);
	
SELECT * FROM posts;