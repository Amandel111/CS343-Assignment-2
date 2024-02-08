package main

//package server
import (
	//"errors"
	"fmt"
	//"client"
	"log"
	"net"
	"net/rpc"

	"net/http"

)
// need to be part of the same package to acces eachother's code

type WordCount int
// Question: is parsing the doc done outside a function?

//InputChunk in this case is the chunk of input we plan to pass to a mapepr
type InputChunk struct{
	//string chunkID
	ChunkContent string
}

// type CounterMap struct {
// 	dict map[string]int //= make(map[string]int)
// }

// how does the leader access the map information stored in local files output
// fortm this RPC?
// RPC func?

// map function
func (t *WordCount) Map(content string, reply *map[string]int) error {
	//for 
	fmt.Println("Hello world")
	dict := make(map[string]int)
	*reply = dict //this is going to be key value, pair of a word and its count 1
	//this reply is going to be the location of the key value pair, and this location gets sent to reduce
	
	//should the leader be accessing intermediate info via reply from RPC
	//or should leader have an rpc that gets called by mapper periodically int order to
	//pass location info?

	//what does leader pass to reduce? Does it just pass the lcoation of the info stored on mapper pc
	//or also pass the key value pair of word type and count?

	//how does the mapper store its data locally? How do we code this? Is it all stored in a file
	//or some other contiguous chunk? 

	// who/what reads the files themselves - prolly the leader?
	return nil
}
/*
Step 1: leader calls RPC map and passes InputChunk w file info to parse
Step 2: map parces InputChunk into key-value pairs and stores those locally
Step 3: Once map has parced all of InputChunk, it "replies" to the leader
with the location address of the information it stored
Step 4: Leader calls RPC reduce and passes location information from mapper 
to reducer so that the reducer 
Find all the values of the word through reduce and then a new reduce process ocucrs
*/
// reducer looks at what information is stored in files and returns to leader?
func (t *WordCount) Reduce(args *map[string]int, reply *map[string]int) {
	//originally had args *Maps
	//return nil
}

// leader reads in; make a struct of a map of values and store it there
// main function in worker depends, leader should have a main function
// worker is the processt that does the work

// when the reduction is done, outputs to local files and then files are combined
// and stored in the DFS

	
func main() {
	// Worker publishes the rpc
	words := new(WordCount)
	rpc.Register(words)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)

}
