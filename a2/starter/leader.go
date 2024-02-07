import {
	"net/rpc"
}
type Leader struct {
	files [] string // files
	// address of worker computer?
}

client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
if err != nil {
	log.Fatal("dialing:", err)
}

//maybe the leader has an RPC that map can call in order to notify  the leader where it has stored info
