package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/ThomasLee94/autosuggest/cmd"
	"github.com/ThomasLee94/autosuggest/trie"
)

func main() {
	trie := trie.NewTrie()

	var buffer bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buffer) // Will write to network.
	dec := gob.NewDecoder(&buffer) // Will read from network.

	// Encode (send) some values.
	err := enc.Encode(trie)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// // Decode (receive) and print the values.
	// var q Q
	// err = dec.Decode(&q)
	// if err != nil {
	// 	log.Fatal("decode error 1:", err)
	// }
	// fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	// err = dec.Decode(&q)
	// if err != nil {
	// 	log.Fatal("decode error 2:", err)
	// }
	// fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

	// create trie file
	f, err := os.Create("trie")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err := f.Write(err)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	cmd.RootCmd.Execute()
}
