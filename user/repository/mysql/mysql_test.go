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

func (m *MySQLTest) seedSimpleUser(u user.User) {
	query := "INSERT INTO `user` (`username`, `fullname`, `dob`, `gender`, `source`, `activated`) " +
		"VALUES (?, ?, ?, ?, ?, ?)"

	preparedQuery, err := m.DBConn.Prepare(query)
	assert.NoError(m.T(), err)
	defer preparedQuery.Close()

	_, insertError := preparedQuery.Exec(
		u.Username,
		u.Fullname,
		u.DOB,
		u.Gender,
		u.Source,
		u.Activated,
	)
	assert.NoError(m.T(), insertError)
}

func (m *MySQLTest) TestGetByID() {
	mockUser := userData
	m.seedSimpleUser(mockUser)
	repo := mysql.MySQLRepository{m.DBConn}
	res, err := repo.GetByID(int64(1))

	assert.NotNil(m.T(), res)
	assert.NoError(m.T(), err)
}

func (m *MySQLTest) TestGetByIDNotFound() {
	mockUser := userData
	m.seedSimpleUser(mockUser)
	repo := mysql.MySQLRepository{m.DBConn}
	res, err := repo.GetByID(int64(2))

	assert.Nil(m.T(), res)
	assert.NoError(m.T(), err)
}
