package db

import (
	"database/sql"
	"fmt"
)

const getUser = `
select	"username"
from	"User" "user"
where	"user"."userId" = $1`

const getUserIdByUsername = `
select	"userId"
from	"User" "user"
where	lower("user"."username") = lower($1);
`

const createUserAndGetId = `
insert into "User" ("username")
values ($1)
returning "userId";
`

const deactivateUserIfExists = `
update	"User" as "user"
set		"active" = false
where	lower("user"."username") = lower($1)
and		"user"."active" = true;
`

const isActive = `
select	"active"
from	"User" "user"
where	"user"."userId" = $1
`

type User struct {
	Username  string
	FirstName string
	LastName  string
}

type UserDAO struct {
	database *Database
}

func NewUserDAO(database *Database) *UserDAO {
	return &UserDAO{
		database,
	}
}

func (dao *UserDAO) GetUser(userId int) (User, error) {
	row := dao.database.db.QueryRow(getUser, userId)
	var user User
	err := row.Scan(&user.Username)
	if err != nil {
		return User{}, err
	}
	// TODO: move these to the table
	user.FirstName = "Anton"
	user.LastName = "Johansson"
	return user, nil
}

func (dao *UserDAO) GetOrCreateId(username string) int64 {
	row := dao.database.db.QueryRow(getUserIdByUsername, username)
	var userId int64
	err := row.Scan(&userId)
	if err == nil {
		return userId
	} else if err != sql.ErrNoRows {
		fmt.Println("error scanning 1:", err)
		return -1
	}

	if err := dao.database.db.QueryRow(createUserAndGetId, username).Scan(&userId); err != nil {
		fmt.Println("error creating user:", err)
		return -1
	}
	return userId
}

func (dao *UserDAO) DeactivateIfExists(username string) {
	_, err := dao.database.db.Exec(deactivateUserIfExists, username)
	if err != nil {
		fmt.Println("error deactivating:", err)
	}
}

// IsActive checks if a user is active
func (dao *UserDAO) IsActive(userId int64) bool {
	row := dao.database.db.QueryRow(isActive, userId)
	var active bool
	err := row.Scan(&active)
	if err == nil {
		return active
	} else if err != sql.ErrNoRows {
		fmt.Println("Error checking active status:", err)
	}
	return false
}
