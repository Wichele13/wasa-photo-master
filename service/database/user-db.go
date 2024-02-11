package database

import (
	"database/sql"
)

// CreateUser creates a new user
func (db *appdbimpl) CreateUser(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(username) VALUES (?)", u.Username)
	if err != nil {
		var user User
		if err := db.c.QueryRow(`SELECT id, username FROM users WHERE username = ?`, u.Username).Scan(&user.Id, &user.Username); err != nil {
			if err == sql.ErrNoRows {
				return user, ErrUserDoesNotExist
			}
		}
		return user, nil
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.Id = uint64(lastInsertID)
	return u, nil
}

// SetUsernames sets the username of a user
func (db *appdbimpl) SetUsername(u User, username string) (User, error) {
	res, err := db.c.Exec("UPDATE users SET username = ? WHERE id = ? AND Username=?", u.Username, u.Id, username)
	if err != nil {
		return u, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return u, err
	} else if affected == 0 {
		return u, err
	}
	return u, nil
}

// GetUsernames gets the username of a user
func (db *appdbimpl) GetUserId(username string) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE username = ?`, username).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil

}

// CheckUserByUsername checks if a user exists
func (db *appdbimpl) CheckUserByUsername(u User) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE username = ?`, u.Username).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}

// CheckUserById checks if a user exists
func (db *appdbimpl) CheckUserById(u User) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE id = ?`, u.Id).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}

// CheckUser checks if a user exists
func (db *appdbimpl) CheckUser(u User) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE id = ? AND username = ?`, u.Id, u.Username).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}

// GetMyStream gets the photo stream of a user
func (db *appdbimpl) GetMyStream(u User) ([]PhotoStream, error) {
	rows, err := db.c.Query(`SELECT id, userId, file, date, likeCount, commentCount FROM photostream WHERE userId = ?`, u.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photoStreams []PhotoStream
	for rows.Next() {
		var photoStream PhotoStream
		if err := rows.Scan(&photoStream.Id, &photoStream.UserId, &photoStream.File, &photoStream.Date, &photoStream.LikeCount, &photoStream.CommentCount); err != nil {
			return nil, err
		}
		photoStreams = append(photoStreams, photoStream)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return photoStreams, nil
}
