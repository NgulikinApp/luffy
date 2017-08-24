package mysql

import (
	"database/sql"
	"time"

	"github.com/NgulikinApp/luffy/user"
	log "github.com/Sirupsen/logrus"
	sq "github.com/elgris/sqrl"
)

type MySQLRepository struct {
	Conn *sql.DB
}

func (m *MySQLRepository) GetByID(id int64) (*user.User, error) {
	query := sq.Select(`id, username, fullname, DATE_FORMAT(dob, "%Y-%m-%d"), gender, source, activated`).From("user")
	query.Where("id = ?", id)

	sql, args, _ := query.ToSql()

	u := new(user.User)
	res, err := m.Conn.Query(sql, args...)
	if err != nil {
		log.Error(err, sql, id)
		return nil, err
	}
	defer res.Close()

	result, err := m.unmarshal(res)
	if err != nil {
		return nil, err
	}

	if len(result) < 1 {
		return nil, err
	}
	u = result[0]

	return u, err
}

func (m *MySQLRepository) Store(usr *user.User) error {
	trx, err := m.Conn.Begin()
	if err != nil {
		return err
	}

	query := sq.Insert("user").
		Columns(`username`, `fullname`, `dob`, `gender`, `time_signup`).
		Values(usr.Username, usr.Fullname, usr.DOB, usr.Gender, time.Now())

	queryStr, args, _ := query.ToSql()

	r, err := trx.Exec(queryStr, args...)

	if err != nil {
		trx.Rollback()
		return err
	}

	usr.ID, err = r.LastInsertId()
	if err != nil {
		trx.Rollback()
		return err
	}

	trx.Commit()

	return nil
}

func (m *MySQLRepository) unmarshal(rows *sql.Rows) ([]*user.User, error) {
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
