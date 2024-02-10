package main
//package server
import (
	"net/rpc"
	"fmt"
	//"client"
	"log"

	//"errors"

	
)
// leader and worker are servers
// the leader is supposed to call the worker's rpcs
// a worker 
type Leader struct {
	files [] string // files
	// address of worker computer?
}


type InputChunk struct{
	//string chunkID
	ChunkContent string
}


//maybe the leader has an RPC that map can call in order to notify  the leader where it has stored info

// func rpcCall(chunk []string) {
// 	//call map function
// 	//FIX
// 	// args := &server.Args{chunk}
// 	args := &Args{chunk}
// 	fmt.Printf("Reached");
// 	var reply int
// 	err = client.Call("WordCount.Map", args, &reply)
// 	if err != nil {
// 		log.Fatal("map error:", err)
// 	}
// 	fmt.Printf("Map: %d", reply)
// 	}
	// map := new(Map)
	// divCall := client.Go("Arith.Divide", args, quotient, nil)
	// replyCall := <-divCall.Done	// will be equal to divCall
	
func main() {
	
	// directory := os.Args[1]
    // files, err := os.ReadDir(directory)
    // if err != nil {
    //     fmt.Println("Error:", err)
    //     return
    // }
    // var fileNames []string
    // for _, file := range files {
    //     // fmt.Println(file.Name())
    //     fileNames = append(fileNames, directory+"/"+file.Name())
    // }

	// var lines []string
	// for _, fileName := range fileNames {
	// 	file, err := os.Open(fileName)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer file.Close()

	// 	fileScanner := bufio.NewScanner(file)

	// 	fileScanner.Split(bufio.ScanLines)

	// 	// scan each file line by line
	// 	for fileScanner.Scan() {
	// 		line := fileScanner.Text()
	// 		lines = append(lines, line)
	// 	}
	// }
	// numLinesPerChunk := len(lines) / numChunks


    // for index := 0; index < numChunks; index++ {

    //     lowerBound := index * numLinesPerChunk
    //     upperBound := lowerBound + numLinesPerChunk
    //     slice := lines[lowerBound:upperBound]
    //     //wg.Add(1) // do we need this?
    //     rpcCall(slice)
    // }
    // slice := lines[numLinesPerChunk*numChunks : len(lines)]
    // //wg.Add(1)
    // //go start_thread(slice, numChunks)
    // wg.Wait()

	//dial the server
	serverAddress := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1235")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	// Calling Map
	chunk := "hello"
	
	//rpcCall(array)
	args := chunk //go's version of a constructor, passing the paramter to be saved in the object
	fmt.Printf("Reached");
	var reply map[string]int
	err = client.Call("WordCount.Map", args, &reply)
	if err != nil {
		log.Fatal("map error:", err)
	}
	fmt.Printf("Map: %d", reply)

	// Calling Reduce
	listOfMaps := make([]map[string]int, 1)
	var response map[string]int
	err = client.Call("WordCount.Reduce", listOfMaps, &response) // fix
	if err != nil {
		log.Fatal("reduce error:", err)
	}
	fmt.Printf("Reduce: %d", reply)
	

}