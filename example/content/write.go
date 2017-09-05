package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	_ "time"
)

func main() {
	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	db, err := leveldb.OpenFile("db_store2", nil)
	if err != nil {
		log.Printf("failed to openfile, %s", err)
		return
	}

	defer db.Close()
	log.Printf("init db ok\n")

	for i := 1000000; i < 2000000; i++ {
		err = db.Put([]byte(fmt.Sprintf("key:%d", i)), []byte(fmt.Sprintf("val:%d", i)), nil)
		if err != nil {
			log.Printf("failed to put, %s", err)
			return
		}
	}

	data, err := db.Get([]byte("key:1999999"), nil)
	if err != nil {
		log.Printf("failed to get, %s", err)
		return
	}

	log.Printf("Got %s", data)
	// time.Sleep(time.Second * 600)

	tr, _ := db.OpenTransaction()

	v, _ := tr.Get([]byte("foo"), nil)
	log.Printf("get foo, v=%s", string(v))

	tr.Put([]byte("foo"), []byte("bar2"), nil)
	v, _ = tr.Get([]byte("foo"), nil)
	log.Printf("get foo, v=%s", string(v))
	tr.Commit()
}
