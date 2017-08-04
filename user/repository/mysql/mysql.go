package mysql

import (
	"database/sql"

	"github.com/NgulikinApp/luffy/user"
	log "github.com/Sirupsen/logrus"
	sq "github.com/elgris/sqrl"
)

type UserRepository struct {
	Conn *sql.DB
}

func (self *UserRepository) GetByID(id int64) (*user.User, error) {
	query := sq.Select(`id, username, fullname, DATE_FORMAT(dob, "%Y-%m-%d"), gender, source, activated`).From("user")
	query.Where("id = ?", id)

	sql, args, _ := query.ToSql()

	u := new(user.User)
	res, err := self.Conn.Query(sql, args...)
	if err != nil {
		log.Error(err, sql, id)
		return nil, err
	}
	defer res.Close()

	result, err := self.unmarshal(res)
	if err != nil {
		return nil, err
	}

	if len(result) < 1 {
		return nil, err
	}
	u = result[0]

	return u, err
}

func (self *UserRepository) unmarshal(rows *sql.Rows) ([]*user.User, error) {
	defer rows.Close()

	results := []*user.User{}

	for rows.Next() {
		var u user.User

		err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Fullname,
			&u.DOB,
			&u.Gender,
			&u.Source,
			&u.Activated,
		)
		if err != nil {
			log.Error(err, u)
			return results, err
		}
		results = append(results, &u)
	}
	return results, nil
}
