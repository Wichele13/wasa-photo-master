package database

// Comment is the struct for the comment
func (db *appdbimpl) AddComment(c Comment) (Comment, error) {
	_, err := db.c.Exec(`INSERT INTO comments (id, userId, photoId, comment, date) VALUES (?, ?, ?, ?, ?)`, c.Id, c.UserId, c.PhotoId, c.Comment, c.Date)
	if err != nil {
		return c, err
	}
	return c, nil
}

// RemoveComment removes a comment from the database
func (db *appdbimpl) RemoveComment(c Comment) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE id=?`, c.Id)
	if err != nil {
		return err
	}
	return nil
}

// GetComments returns all comments for a photo
func (db *appdbimpl) GetComments(photoId uint64) ([]Comment, error) {
	var ret []Comment
	rows, err := db.c.Query(`SELECT id, userId, photoId, comment, date FROM comments WHERE photoId = ?`, photoId)
	if err != nil {
		return ret, ErrCommentDoesNotExist
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var b Comment
		err = rows.Scan(&b.Id, &b.UserId, &b.PhotoId, &b.Comment, &b.Date)
		if err != nil {
			return nil, err
		}
		ret = append(ret, b)
	}
	return ret, nil
}

// GetCommentsCount returns the number of comments for a photo
func (db *appdbimpl) GetCommentsCount(photoId uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM comments WHERE photoId = ?`, photoId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// GetCommentById returns a comment by its ID
func (db *appdbimpl) GetCommentById(c Comment) (Comment, error) {
	err := db.c.QueryRow(`SELECT id, userId, photoId, comment, date FROM comments WHERE id = ?`, c.Id).Scan(&c.Id, &c.UserId, &c.PhotoId, &c.Comment, &c.Date)
	if err != nil {
		return c, err
	}
	return c, nil
}
