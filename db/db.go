package db

import (
	"encoding/binary"
	"time"

	trie "github.com/ThomasLee94/autosuggest/trie"
	"github.com/boltdb/bolt"
)

// Define our bucket
var trieBucket = []byte("trie")

// Define our db
var db *bolt.DB

// Task struct
type Task struct {
	Key   int
	Value string
}

// Init function for initializing our database and any buckets we'd like to be stored within it
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	// Update the database in the context of a read-write managed transaction
	// If there is an error (the bucket exists) the transaction will get rolled back
	return db.Update(func(tx *bolt.Tx) error {
		// Create a new bucket if it doesn't exist,
		// Throw an error if it does
		_, err := tx.CreateBucketIfNotExists(trieBucket)
		return err
	})
}

// Create a task
func CreateTrie(task string) (int, error) {
	var id int

	// Start our read-write transaction with the database
	err := db.Update(func(tx *bolt.Tx) error {
		// Grab the task bucket
		bucket := tx.Bucket(trieBucket)
		// Get the next incremental id
		id64, _ := bucket.NextSequence()
		// Store our id outside of the closure
		id = int(id64)
		// convert the id into a byte slice to use as a bolt db key
		key := itob(id64)
		// Store both the key and task into our db
		return bucket.Put(key, []byte(task))
	})

	// Check if the transaction threw an error
	if err != nil {
		return -1, err
	}

	// Returns the id and no error, indicating all is good!
	return id, nil
}

func UpdateTrie(trie trie.Trie) trie {

}

// List all tasks that a user has created
func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		// Grab the task bucket
		bucket := tx.Bucket(trieBucket)

		// grab a cursor from the bucket
		// The cursor is only valid for the lifetime of the trasanction
		cursor := bucket.Cursor()

		// using the cursor, iterate over all of our key/value pairs for as long as
		// there aren't any more keys left
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			// convert both the key and value into data types suitable for our
			// Task structure and make sure that we use new variables to make sure
			// we're not overwriting any of the k,v in bolt
			copiedKey := int(btoi(key))
			copiedValue := string(value)

			// Append the task to our tasks slice
			tasks = append(tasks, Task{
				Key:   copiedKey,
				Value: copiedValue,
			})
		}
		return nil
	})

	// If the tx fails, return the error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Conver an integer to a byte slice
// in big endian order so that our oldest keys will be in positions
// where they are the first keys retrieved from
func itob(integer uint64) []byte {
	// Create an with 8 zeroed positions
	byteSlice := make([]byte, 8)
	// Convert the unsigned integer to a big endian byte string that
	// gets stored within our byte slice
	binary.BigEndian.PutUint64(byteSlice, integer)
	return byteSlice
}

// Convert a byte slice into an integer
// Useful for getting the values of the byte slices we're storing into
// the dbs as keys
func btoi(byteSlice []byte) uint64 {
	return binary.BigEndian.Uint64(byteSlice)
}
