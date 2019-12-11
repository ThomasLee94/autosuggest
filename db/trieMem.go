package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
	"github.com/ThomasLee94/autosuggest/trie"
)

var trieBucket = []byte("trie") // create the name key for our bucket
var db *bolt.DB

//trieMem is a trie struct
type Task struct {
	Key   int
	Value string
}

//Init sets up our bolt database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	//	db.Update updates the data base with the callback
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(trieBucket)
		return err //	returns any errors that occur while bucket being created
	})
}

//CreateTask takes in a string as the task and writes a new todo item to db
func CreateTask(task string) (int, error) {
	var id int //	create id variable
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()     //	grabs the next id number generated by nextSequence()
		id = int(id64)                  //	assign id as int of id64 to be passed through itob()
		key := itob(id)                 //	converts the id into a byte slice
		return b.Put(key, []byte(task)) //	stores the key and value(task as byte slice) in the boltdb
	})
	if err != nil {
		return -1, err
	}
	return id, nil //	returns the id value
}

//AllTasks returns a slice of all tasks in list (boltdb)
func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})

		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//DeleteTask removes task from db when completed
func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
	if err != nil {
		return err
	}

	return nil
}

//itob takes in interger and makes in byte string
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//btoi reverts byte slice to interger
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}