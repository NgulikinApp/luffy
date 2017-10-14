package mongodb

import (
	"github.com/NgulikinApp/luffy/product"
	"github.com/NgulikinApp/luffy/product/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDBRepository struct {
	Session *mgo.Session
	DBName  string
}

func (r *MongoDBRepository) Store(p *product.Product) error {
	c := r.Session.DB(r.DBName).C(`product`)
	err := c.Insert(p)
	return err
}

func (r *MongoDBRepository) Fetch(filter product.Filter, num int64, cursor int64) ([]*product.Product, error) {
	c := r.Session.DB(r.DBName).C(`product`)

	q := bson.M{}
	p := make([]*product.Product, 0)
	err := c.Find(q).Skip(int(cursor)).Limit(int(num)).All(&p)

	return p, err
}

func NewRepository(s *mgo.Session, name string) repository.Repository {
	return &MongoDBRepository{
		Session: s,
		DBName:  name,
	}
}
