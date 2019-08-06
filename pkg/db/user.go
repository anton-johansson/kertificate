package db

import (
	"database/sql"
	"fmt"

	"kertificate.io/kertificate/pkg/model"
)

const getUser = `
select  "user"."username"
,       "user"."firstName"
,       "user"."lastName"
,       "user"."emailAddress"
from    "User" "user"
where   "user"."userId" = $1`

const updateUserByUsername = `
update  "User" "user"
set     "username" = $1
,       "firstName" = $2
,       "lastName" = $3
,       "emailAddress" = $4
,       "loggedInAt" = now()
where   lower("user"."username") = lower($1)
returning "userId";`

const createUserAndGetId = `
insert into "User"
(
        "username"
,       "firstName"
,       "lastName"
,       "emailAddress"
)
values
(
        $1
,       $2
,       $3
,       $4
)
returning "userId";`

const deactivateUserIfExists = `
update	"User" as "user"
set		"active" = false
where	lower("user"."username") = lower($1)
and		"user"."active" = true;`

const isActive = `
select	"user"."active"
from	"User" "user"
where	"user"."userId" = $1`

type UserDAO struct {
	database *Database
}

func NewUserDAO(database *Database) *UserDAO {
	return &UserDAO{database}
}

func (dao *UserDAO) GetUser(userId int) (model.User, error) {
	row := dao.database.db.QueryRow(getUser, userId)
	var user model.User
	if err := row.Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.EmailAddress); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (dao *UserDAO) GetOrCreateUser(user model.User) (model.User, error) {
	row := dao.database.db.QueryRow(updateUserByUsername,
		user.Username,
		user.FirstName,
		user.LastName,
		user.EmailAddress)
	err := row.Scan(&user.UserId)
	if err == nil {
		return user, nil
	} else if err != sql.ErrNoRows {
		fmt.Println("error scanning 1:", err)
		return model.User{}, err
	}

	if err := dao.database.db.QueryRow(createUserAndGetId,
		user.Username,
		user.FirstName,
		user.LastName,
		user.EmailAddress).Scan(&user.UserId); err != nil {
		fmt.Println("error creating user:", err)
		return model.User{}, err
	}
	return user, nil
}

func (dao *UserDAO) DeactivateIfExists(username string) {
	_, err := dao.database.db.Exec(deactivateUserIfExists, username)
	if err != nil {
		fmt.Println("error deactivating:", err)
	}
}

// IsActive checks if a user is active
func (dao *UserDAO) IsActive(userId int) bool {
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
