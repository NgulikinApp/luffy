package mysql_test

import (
	"database/sql"
	"testing"

	"github.com/NgulikinApp/luffy/db"
	"github.com/NgulikinApp/luffy/user"
	"github.com/NgulikinApp/luffy/user/repository/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MySQLTest struct {
	db.MySQLSuite
	conn *sql.DB
}

var (
	userData = user.User{
		Username:  `gama`,
		Fullname:  `andhika gama`,
		DOB:       `1992-03-23`,
		Gender:    `male`,
		Source:    `web`,
		Activated: true,
	}

	falseUser = user.User{
		Username:  `this kind of username is far too long from the maximum username length`,
		Fullname:  `andhika gama`,
		DOB:       `false-date`,
		Gender:    `male`,
		Source:    `web`,
		Activated: true,
	}
)

func TestMySQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip mysql repository test")
	}

	println("Running admin test")
	suite.Run(t, new(MySQLTest))
}

func (self *MySQLTest) seedSimpleUser(u user.User) {
	query := "INSERT INTO `user` (`username`, `fullname`, `dob`, `gender`, `source`, `activated`) " +
		"VALUES (?, ?, ?, ?, ?, ?)"

	preparedQuery, err := self.DBConn.Prepare(query)
	assert.NoError(self.T(), err)
	defer preparedQuery.Close()

	_, insertError := preparedQuery.Exec(
		u.Username,
		u.Fullname,
		u.DOB,
		u.Gender,
		u.Source,
		u.Activated,
	)
	assert.NoError(self.T(), insertError)
}

func (self *MySQLTest) TestGetByID() {
	mockUser := userData
	self.seedSimpleUser(mockUser)
	repo := mysql.UserRepository{self.DBConn}
	res, err := repo.GetByID(int64(1))

	assert.NotNil(self.T(), res)
	assert.NoError(self.T(), err)
}

func (self *MySQLTest) TestGetByIDNotFound() {
	mockUser := userData
	self.seedSimpleUser(mockUser)
	repo := mysql.UserRepository{self.DBConn}
	res, err := repo.GetByID(int64(2))

	assert.Nil(self.T(), res)
	assert.NoError(self.T(), err)
}
