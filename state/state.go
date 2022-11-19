package state

import (
	"github.com/boltdb/bolt"
	log "github.com/sirupsen/logrus"
)
func openDB() *bolt.DB {
	db, err := bolt.Open("state.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Write(bucket []byte, key []byte, value []byte) error {
	db := openDB()
	defer db.Close()
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("state/state.go:Balance(): Failed to write")
			return err
		}
		return b.Put(key, value)
	})
	return err
}

func Read(bucket []byte, key []byte) ([]byte, error) {
	db := openDB()
	var val []byte
	var ret []byte
	defer db.Close()
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			log.Error("state/state.go:Balance(): Failed open bucket")
			return nil
		}
		ret = b.Get(key)
		return nil
	})
	copy(ret, val)
	return val, err
}
