package database

func (db *appdbimpl) CreateUser(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users (username) VALUES (?, ?, ?)", u.Username)

}
