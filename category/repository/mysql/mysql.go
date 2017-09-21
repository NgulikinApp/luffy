package mysql

import (
	"database/sql"

	"github.com/NgulikinApp/luffy/category"
	log "github.com/Sirupsen/logrus"
	sq "github.com/elgris/sqrl"
)

type MySQLRepository struct {
	Conn *sql.DB
}

func (m *MySQLRepository) Fetch(num int64, cursor int64) ([]*category.Category, error) {
	query := sq.Select(`id, name`).From(`category`).Limit(uint64(num)).Offset(uint64(cursor))

	sql, args, _ := query.ToSql()

	res, err := m.Conn.Query(sql, args...)
	if err != nil {
		log.Error(err, sql)
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

	return result, err
}

func (m *MySQLRepository) unmarshal(rows *sql.Rows) ([]*category.Category, error) {
	defer rows.Close()

	results := []*category.Category{}

	for rows.Next() {
		var cat category.Category

		err := rows.Scan(
			&cat.ID,
			&cat.Name,
		)
		if err != nil {
			log.Error(err, cat)
			return results, err
		}
		results = append(results, &cat)
	}
	return results, nil
}
