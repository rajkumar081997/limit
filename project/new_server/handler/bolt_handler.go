package handler

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type Bolt_handler struct {
	root       string
	dbname     string
	bucketname string
	db         *bolt.DB
}

func (b *Bolt_handler) Close() error {
	b.db.Close()
	return nil
}

func Newbolt_handler(s *Bolt_handler) (*Bolt_handler, error) {
	s.root = "../"
	s.dbname = "my.bolt"
	s.bucketname = "world"
	db, err := bolt.Open(s.root+s.dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	return &Bolt_handler{
		root:       s.root,
		dbname:     s.dbname,
		bucketname: s.bucketname,
		db:         db,
	}, err

}

type Item struct {
	Descript string
}

func (i *Item) Marshal() ([]byte, error) {
	return []byte(i.Descript), nil
}

func (i *Item) Unmarshal(key []byte) error {
	i.Descript = string(key)
	return nil
}

func (s *Bolt_handler) Item(input string) (string, error) {

	var i Item
	key := strconv.Itoa(rand.Int())

	i.Descript = input
	value, err := i.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	s.db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte(s.bucketname))
		if err != nil {
			log.Fatal(err)
		}
		i.Descript = string(key)
		var key []byte
		key, err = i.Marshal()

		bucket.Put(key, value)
		return nil
	})
	var returnstring string = "Data stored Successfully , Id is : "
	returnstring = returnstring + key
	return returnstring, err
}

func (b *Bolt_handler) GetId(input string) (string, error) {

	var i Item
	var value []byte
	i.Descript = input

	key, err := i.Marshal()
	if err != nil {
		log.Fatal()
	}

	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.bucketname))

		value = bucket.Get(key)
		return nil
	})
	err = i.Unmarshal(value)
	if err != nil {
		log.Fatal(err)
	}

	return i.Descript, err
}

func (b *Bolt_handler) List(input string) ([]string, error) {

	numberofItem, _ := strconv.Atoi(input)
	var stringvalues []string

	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.bucketname))

		iterator := bucket.Cursor()
		var i int = 0
		for k, v := iterator.First(); k != nil; k, v = iterator.Next() {
			stringvalues = append(stringvalues, string(v))
			i++
			if i == numberofItem {
				break
			}
		}

		return nil
	})

	return stringvalues, nil
}

func (b *Bolt_handler) Remove(input string) (string, error) {

	var i Item
	i.Descript = input
	key, err := i.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.bucketname))

		bucket.Delete(key)
		return nil
	})

	return "Data Deleted Sucessfully", err
}
