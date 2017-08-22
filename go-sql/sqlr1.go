package main

import "time"

// begin null OMIT
type User struct {
	ID       int64 `sql:"primary key"`
	Name     string
	Nickname string `sql:"null"` // empty string is stored as NULL // HL
}

// end null OMIT

// begin json OMIT
type DocumentRow struct {
	ID         int64 `sql:"primary key"`
	Name       string
	ModifiedAt time.Time
	Contents   *Document `sql:"json"` // struct serialized in DB as JSON // HL
}

// end json OMIT

// begin wherein OMIT
func doWhereIn(ids []int64) ([]User, error) {
	var users []User
	err := schema.Select(db, &users, "select {} from users where id in (?)", ids) // HL
	if err != nil {
		return nil, err
	}
	return users, nil
}

// end wherein OMIT
