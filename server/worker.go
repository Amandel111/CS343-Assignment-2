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

// type CounterMap map[string]int // may not need

// how does the leader access the map information stored in local files output
// fortm this RPC?
// RPC func?

// map function

// var m CounterMap

func (t *WordCount) Map(content string, reply *map[string]int) error {
	//for 
	fmt.Println("Hello world")
	fmt.Println("Mapper", content[:20])
	fmt.Println(reply)
	dict := make(map[string]int)
	
	dict["hello"] = 1
	*reply = dict //this is going to be key value, pair of a word and its count 1
	
	return nil
}

// reducer looks at what information is stored in files and returns to leader?
func (t *WordCount) Reduce(args *[]map[string]int, reply *map[string]int) error {
	//originally had args *Maps
	//return nil
	fmt.Println("Hello world")
	fmt.Println("Reducer", args)
	fmt.Println(reply)
	dict := make(map[string]int)
	*reply = dict
	return nil
}


	
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
