package main
//package server
import (
	"net/rpc"
	"fmt"
	"log"
	"strings"
	//"errors"
	"os"
	"strconv"
	"bufio"
	// "context"

	
)

type Leader struct {
	files [] string 
}


type InputChunk struct{
	ChunkContent string
}


//maybe the leader has an RPC that map can call in order to notify  the leader where it has stored info

func MapCall(client *rpc.Client, content string, port int) map[string]int { // make the channel an input
	
	
	//call map RPC
	// args := content //go's version of a constructor, passing the paramter to be saved in the object
	//fmt.Printf("Reached");
	var reply map[string]int
	err := client.Call("WordCount.Map", content, &reply)
	//<- WordCount.Map.Done()
	if err != nil {
		log.Fatal("map error:", err)
	}
	
	//fmt.Printf("Map: %d", reply)
	

	call := client.Go("WordCount.Map", content, &reply, nil)
	
	//return client.Call("WordCount.Map", content, &reply)
	return reply
}


func ReduceCall(mapList []map[string]int, port int) map[string]int{
	// dial server
	serverAddress := "127.0.0.1:"
	client, err := rpc.DialHTTP("tcp", serverAddress + strconv.Itoa(port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//call reduce service
	var response map[string]int
	err = client.Call("WordCount.Reduce", mapList, &response) // fix
	if err != nil {
		log.Fatal("reduce error:", err)
	}
	//fmt.Printf("Reduce: %d", response)

	//return client.Call("WordCount.Reduce", mapList, &response)
	return response
	
}
	
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

		fi, err := file.Stat()
		 if err != nil {
			log.Fatal(err)
		}
		sizeOfFile := fi.Size();
		fmt.Print("size of file", sizeOfFile);
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
	// listOfMaps := make([]map[string]int, numChunks+1)
	var listOfMaps []map[string]int

	replies := make([]chan map[string]int, ) // finish making a list

	portNumber := 1234
    for index := 0; index < numChunks; index++ {

        lowerBound := index * numLinesPerChunk
        upperBound := lowerBound + numLinesPerChunk
        slice := lines[lowerBound:upperBound]
		//loop thorugh slice, append each element to contentAsString
		contentAsString := strings.Join(slice, " ")
        //wg.Add(1) // do we need this?
		//listOfMaps[index] = MapCall(contentAsString ,portNumber)
        serverAddress := "127.0.0.1:"
		client, err := rpc.DialHTTP("tcp", serverAddress + strconv.Itoa(portNumber))
		// take care of case?
		numCallsToMap := len(contentAsString) / 100
		for i := 0; i < numCallsToMap-1; i++ {
			
			// client, err := rpc.DialHTTP("tcp", serverAddress + strconv.Itoa(port))
			if err != nil {
				log.Fatal("dialing:", err)
			}
			// result := MapCall(client, contentAsString[i*100:(i+1)*100],portNumber)
			// <- result.Done()
			// listOfMaps = append(listOfMaps, result)

			// listOfMaps = append(listOfMaps, MapCall(client, contentAsString[i*100:(i+1)*100],portNumber))
			var reply map[string]int
			//err := client.Call("WordCount.Map", contentAsString[i*100:(i+1)*100], &reply)
			//<- WordCount.Map.Done()
			// if err != nil {
			// 	log.Fatal("map error:", err)
			// }
			
			//fmt.Printf("Map: %d", reply)
			
			call := client.Go("WordCount.Map", contentAsString[i*100:(i+1)*100], &reply, nil)
			reply <- call.Done //channel
			//listOfMaps = append(listOfMaps, reply)
			replies = append(replies, reply)
		}
		// count = 
		portNumber += 1 

		
    }
	// portNumber += 1
	print(portNumber)
    slice := lines[numLinesPerChunk*numChunks : len(lines)]
	contentAsString := strings.Join(slice, " ")
	numCallsToMap := len(contentAsString) / 100
	serverAddress := "127.0.0.1:"
	client, err := rpc.DialHTTP("tcp", serverAddress + strconv.Itoa(portNumber))
	// client, err := rpc.DialHTTP("tcp", serverAddress + strconv.Itoa(port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// <-workerCall.Done
	listOfMaps = append(listOfMaps, MapCall(client, contentAsString[(numCallsToMap)*100:],portNumber))
	
	for i := 0; i <= numCallsToMap; i++ {
		
	}
	//listOfMaps = append(listOfMaps, MapCall(contentAsString ,portNumber))
	//listOfMaps[numChunks] = MapCall(contentAsString ,portNumber)
	//fmt.Print("list of maps", len(listOfMaps))
	
	reduced :=  ReduceCall(listOfMaps, portNumber+1) // run on new port
	//fmt.Print("reduced ", reduced)

	//Remove the output directory if it already exists
    err = os.RemoveAll("output")
    if err != nil {
        fmt.Println(err)
    }
   
    // Create output directory and file
    err2 := os.Mkdir("output", os.ModePerm)
    if err2 != nil {
        fmt.Println("Error creating directory:", err2)
    }

	file, err := os.Create("output" + "/" + "results2.txt")
    if err != nil {
        fmt.Println("failed to create file")
    }

	for key, value := range reduced{
		s := key + " " + strconv.FormatInt(int64(value), 10) + "\n"
		file.WriteString(s)
	}
	

}

//<-workerCall.Done -> name of the rpc request

// makerquest - rpc call
// each worker is workng individually
// calling worker before it finished the process
// which worker is which worker

// not a waitgroup