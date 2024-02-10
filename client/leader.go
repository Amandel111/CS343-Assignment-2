package main
//package server
import (
	"net/rpc"
	"fmt"
	//"client"
	"log"
	"strings"
	//"errors"
	"os"
	"strconv"
	"bufio"

	
)

type Leader struct {
	files [] string 
}


type InputChunk struct{
	ChunkContent string
}


//maybe the leader has an RPC that map can call in order to notify  the leader where it has stored info

func MapCall(content string, port int) map[string]int {
	serverAddress := "127.0.0.1:"
	client, err := rpc.DialHTTP("tcp", serverAddress + strconv.Itoa(port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//call map RPC
	// args := content //go's version of a constructor, passing the paramter to be saved in the object
	fmt.Printf("Reached");
	var reply map[string]int
	err = client.Call("WordCount.Map", content, &reply)
	if err != nil {
		log.Fatal("map error:", err)
	}
	fmt.Printf("Map: %d", reply)

	return reply
}

// func ReduceCall(mapList []map[string]int) map[string]int{
// 	// Calling Reduce
// 	//listOfMaps := make([]map[string]int, 1)
// 	var response map[string]int
// 	err = client.Call("WordCount.Reduce", mapList, &response) // fix
// 	if err != nil {
// 		log.Fatal("reduce error:", err)
// 	}
// 	fmt.Printf("Reduce: %d", reply)

// 	return response
	
// }
	
func main() {

	//dial the server
	

	
	directory := os.Args[1]
    files, err := os.ReadDir(directory)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    var fileNames []string
    for _, file := range files {
        // fmt.Println(file.Name())
        fileNames = append(fileNames, directory+"/"+file.Name())
    }

	var lines []string
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		// scan each file line by line
		for fileScanner.Scan() {
			line := fileScanner.Text()
			lines = append(lines, line)
		}
	}
	var numChunks int = 1 //this should always be equal to the number of workers running - 1 if numChunks > 1

	numLinesPerChunk := len(lines) / numChunks
	listOfMaps := make([]map[string]int, numChunks+1)

	portNumber := 1233
    for index := 0; index < numChunks; index++ {

        lowerBound := index * numLinesPerChunk
        upperBound := lowerBound + numLinesPerChunk
        slice := lines[lowerBound:upperBound]
		//loop thorugh slice, append each element to contentAsString
		contentAsString := strings.Join(slice, " ")
        //wg.Add(1) // do we need this?
		portNumber += 1 
        listOfMaps = append(listOfMaps, MapCall(contentAsString ,portNumber))
    }
    slice := lines[numLinesPerChunk*numChunks : len(lines)]
	contentAsString := strings.Join(slice, " ")

	listOfMaps = append(listOfMaps, MapCall(contentAsString ,portNumber))
	fmt.Print("list of maps", listOfMaps)

    //wg.Add(1)
    //go start_thread(slice, numChunks)
    //wg.Wait

}