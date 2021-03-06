package main

import (
    "fmt"
    "net"
    "time"
    "log"
    "math/rand"
    "io/ioutil"
    "io"
    "os/exec"
    "bufio"
    "encoding/json"
    "encoding/binary"
    "bytes"
    "strconv"
	"os"
)


//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< GRAPH >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>



func (graph *Graph) Dijkstra(source string, target string)(dist map[string]int,path []string){
    
      //Unvisted
      Q := make(map[string]*Node)
      dist = make(map[string]int)
      prev := make(map[string]*Node)
      
      for key , val := range graph.Nodes {             // Initialization
          dist[key] = 9999999999                  // Unknown distance from source to v
          prev[key] = nil                 // Previous node in optimal path from source
          Q[key] = val                     // All nodes initially in Q (unvisited nodes)
      }
      
      dist[source] = 0                        // Distance from source to source
      
      for(len(Q) != 0){
          u ,keyy := min(Q,dist)    // Source node will be selected first    
          delete(Q,keyy) //remove u from Q 
          //for each neighbor v of u: 
          for key, value := range u.Edges{          // where v is still in Q.
          
              alt := dist[keyy] + value
              if alt < dist[key]{               // A shorter path to v has been found
                  dist[key] = alt 
                  prev[key] = u 
              }
          }
          
      }
      
      
      /*
Now we can read the shortest path from source to target by reverse iteration:
1  S ← empty sequence
2  u ← target
3  while prev[u] is defined:                  // Construct the shortest path with a stack S
4      insert u at the beginning of S         // Push the vertex onto the stack
5      u ← prev[u]                            // Traverse from target to source
6  insert u at the beginning of S             // Push the source onto the stack


*/
      
        var S []string
        u := target
        // Construct the shortest path with a stack S
        for(prev[u] != nil){
            // Push the vertex onto the stack
            S = append([]string{u},S...)
            u = graph.getKey(prev[u])                            // Traverse from target to source
        }
        
      return dist, S
    
    
    
}

//u := *Node in Q with min dist[u] 

func (graph *Graph) getKey(node *Node) (key string) {
    for key ,val := range graph.Nodes{
        if(val == node){
            return key
        }
    }
    
    return "ERROR"
}

func min(Q map[string]*Node,dist map[string]int) (*Node,string) {
    
    var min int
    var minNode *Node
    firstLoop := true
    var keyy string
    
    for key ,value := range Q{
        if(firstLoop){
            min = dist[key]
            minNode = value
            firstLoop = false
            keyy = key
        }else{
            if(dist[key]< min){
                min = dist[key]
                minNode = value    
                keyy = key   
            }
        }
    }
    
    
    return minNode ,keyy
}


type Node struct{
    Edges map[string]int
}

type Graph struct{
    Name string
    TotalNodes int
    Nodes map[string]*Node
}


func (graph *Graph) addUndirectedWeightedVertice(idOfNode1 string ,idOfNode2 string ,weight int){
    
    graph.Nodes[idOfNode1].Edges[idOfNode2] = weight
    graph.Nodes[idOfNode2].Edges[idOfNode1] = weight
    
}

func createGraph(graphName string) *Graph {
    graphThis := Graph{ 
        Name : graphName,
        TotalNodes : 0,
    }
    
    graphThis.Nodes = make(map[string]*Node)
    
    return &graphThis
}

func (graph *Graph) addNode(idOfNode string) {
    
    //TODO: Add Real error
    if _ , isKey := graph.Nodes[idOfNode]; isKey{
        fmt.Println("Key already exists. Did not add node")
    }else{
        
    
        graph.TotalNodes = graph.TotalNodes + 1
        newNode := Node{} 

        graph.Nodes[idOfNode] = &newNode

        newNode.Edges = make(map[string]int)
            
        
    }
    
}

func (graph *Graph) toJson()([]byte){
    theJSON, _ := json.Marshal(graph)
    return theJSON
}

func (graph *Graph) removeNodesWithNoEdges(){

    for _, val := range graph.Nodes{
        
        if(len(val.Edges) == 0){
            //delete(graph.Nodes,key)
            //graph.TotalNodes = graph.TotalNodes - 1
        }
        
    }


}
func (graph *Graph) removeNode(idOfNode string){
    
 
   for key, val := range graph.Nodes{
       
       for keyy := range val.Edges{
           if keyy == idOfNode {
               delete(graph.Nodes[key].Edges,keyy)
           }
       }
       
       
       if(key == idOfNode){
           delete(graph.Nodes,key)
           graph.TotalNodes = graph.TotalNodes -1
       }
   }
    
}





//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<END OF GRAPH >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func createRouter(ip string) *Router{
    this := Router{}
    this.IP = ip
    this.Neighbours = make(map[string]int)
    return &this
}
 
 
type Router struct{
    IP string
    Neighbours map[string]int
}

//Change IP of router
func (router *Router) setIP(ip string){
    router.IP = ip
}

//Add Neighbours
func (router *Router) addNeighbour(ip string, weight int){
    router.Neighbours[ip] = weight
}
//Add Neighbours
func (router *Router) removeNeighbour(ip string){
    delete(router.Neighbours,ip)
}


func main(){
   
   clear()
   println("ʕ◔ϖ◔ʔ  Welcome to the GO NetSim, Client Process!!!  ʕ◔ϖ◔ʔ")
   
    fileName := "client1.json"
    if(len(os.Args[1:]) == 1){
        fileName = os.Args[1]
    }
   
   //Read in JSON file to create the router Struct
    data, err := ioutil.ReadFile(fileName)
    checkError(err)
   
   //Create Router Struct and turn json into Struct
    var router *Router
    err = json.Unmarshal(data,&router)
    checkError(err)
    
    // Something old
    timeStamp := int64(0);
    
    graph := createGraph(strconv.FormatInt(timeStamp, 10))
    
    fmt.Println("Timestamp:"+ strconv.FormatInt(timeStamp, 10))
    
    graph.addNode(router.IP)
    //theJSON, _ := json.Marshal(router)
    //theJSON, _ = json.Marshal(graph)
    
    scanForNeighbours(router,graph)
    
    //fmt.Println(string(theJSON))
    
    go listenForScan(router,graph)
    
    go sendFiles(router, graph)
    
    for{
        time.Sleep(time.Second * 1)
        scanForNeighbours(router,graph)
        
        //jsonOfRouter , _ := json.Marshal(router)
        
        //fmt.Println("State of CLIENT: "+string(jsonOfRouter))
        //fmt.Println("State of Graph:"+string(graph.toJson()))
    }

    
 }

func listenForScan(router *Router,graph *Graph)  {
    ln, err := net.Listen("tcp",router.IP)
    checkError(err)
    for{
        conn, err := ln.Accept()
        checkError(err)
       
        go handleListenForScan(conn,router,graph)
    }
}

func handleListenForScan(conn net.Conn,router *Router,graph *Graph)  {
   
    receivingGraph := byte(0x61)
    receivingScan  := byte(0x62)
    receivingData  := byte(0x01)

    //Generate Random Weight (SEED RANDOM)
    rand.Seed(time.Now().Unix())
    weight := rand.Intn(9) + 1
    
    //Read in the IP From scanner
    connectorIPBytes, err := bufio.NewReader(conn).ReadBytes(byte(0x4))
    
    EOFerror := checkNetworkRead(err)
    
    if(!EOFerror){
        
        if(receivingScan == connectorIPBytes[0] ){
            fmt.Fprintf(conn,strconv.Itoa(weight)+string(byte(0x4)))
            
            pleaseRemoveNewLine := []byte(connectorIPBytes)
            connectorIPBytes = pleaseRemoveNewLine[:len(pleaseRemoveNewLine)-1]

            //fmt.Println(string(connectorIPBytes))
            
        }else if(receivingGraph == connectorIPBytes[0]){
            //println("I'm a client, so I don't need this graph")

        } else if(receivingData == connectorIPBytes[0]) {
            // we are the client, we want to write this data to a file
            // I am thinking length of dest field dest field 
            // length of name field name field
            // length of data field data field
            
            currLoc := uint32(0)
            connectorIPBytes := connectorIPBytes[1:len(connectorIPBytes) - 1]
            sizeOfField := binary.LittleEndian.Uint32(connectorIPBytes[0:4]) // inclusive:exclusive
            println("Size of dest is " + strconv.FormatUint(uint64(sizeOfField), 10))
            dest := connectorIPBytes[4:4+sizeOfField]
            
            // make sure the data is for us
            if(string(dest) != router.IP) {
                fmt.Println("I got data addressed to " + string(dest))
                return
            } else {
                fmt.Println("I got a letter!")
            }
            println("data recieved " + strconv.Itoa(len(connectorIPBytes)))
            
            currLoc = 4 + sizeOfField
            sizeOfField = binary.LittleEndian.Uint32(connectorIPBytes[currLoc:currLoc + 4])
            currLoc += 4
            
            data := connectorIPBytes[currLoc:len(connectorIPBytes)]
            
            // now write to file
            println("Recieved: " + string(data))
            checkError(err)
            
            
        }else{
            fmt.Println("UNEXPECTED CASE HAS HAPPENED ERROR : 325")
        }
    }
    
    conn.Close()
}

func sendFiles(router *Router, graph *Graph) {
    for {
        fmt.Print("Please enter your message: ")
        reader := bufio.NewReader(os.Stdin)
        message, _:= reader.ReadString('\n')
        message = message[0:len(message) - 1]
        message = string(message)
        
        fmt.Print("Please enter your destination with port number: ")
        dest, _:= reader.ReadString('\n')
        dest = dest[0:len(dest) - 1]
        dest = string(dest)
        
        stringToSend := string(byte(0x1))
        buf := new(bytes.Buffer)
        err := binary.Write(buf, binary.LittleEndian, uint32(len(dest)))
        checkError(err)

        stringToSend += string(buf.Bytes()) + dest 
        
        buf = new(bytes.Buffer)
        err = binary.Write(buf, binary.LittleEndian, uint32(len(message)))
        checkError(err)
        stringToSend += string(buf.Bytes()) + message + string(byte(0x4))

        // since we are a client, we only have one neighbor
        for ip := range router.Neighbours {
            conn, err := net.Dial("tcp", ip)
            checkError(err)
            fmt.Fprintf(conn, stringToSend)
        }
    }
    
}


func scanForNeighbours(router *Router,graph *Graph){
    
    for ip, val := range router.Neighbours{
        conn, err := net.Dial("tcp", ip)
        
        
        if(checkNetworkError(err) && val == 0){
            fmt.Fprintf(conn,string(0x62) +router.IP + string(byte(0x4)))
            message , _ := bufio.NewReader(conn).ReadString(byte(0x4)) 
            conn.Close()
            
            //Update Weight in router
            
            messageBytes := []byte(message)
            message = string(messageBytes[:len(messageBytes)-1])
            fmt.Println("Weight: "+message)
            router.Neighbours[ip], _ = strconv.Atoi(message)   
            
            //Add To Graph
            
            println("adding node " + ip)
            
            
            graph.addNode(ip)
            
            graph.addUndirectedWeightedVertice(router.IP,ip,router.Neighbours[ip])  
            
        //If there is a network error happens and the val in router is not 0
        }else if(!checkNetworkError(err) && val != 0){ //DOEST connects and the weight in NOT 0
          router.Neighbours[ip] = 0
          graph.removeNode(ip)
          //graph.Name = strconv.FormatInt(timeStamp,10)
          println("removing node " + ip)
           
        }
        
    }
    
    graph.removeNodesWithNoEdges()  
}

//Makes strings more gophery
func println(dis string) {
    dis = "ʕ◔ϖ◔ʔ " + dis + " ʕ◔ϖ◔ʔ"
    fmt.Println(dis)
}

//General Error Catching
func checkError(err error)  {
    if err != nil {
        log.Fatal(err)
    }
}

func checkNetworkError(err error) (shouldRun bool) {
  skip := false
  if err == nil {
    skip = true
  } else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
    println("Timeout")
  } else if opError, ok := err.(*net.OpError); ok {
    if opError.Op == "dial" {
      //println("Unknown host")
    } else if opError.Op == "read" {
      println("Connection refused")
    }
  }
  return skip
}



//All this function does is executes the clear command.
func clear(){

	cmd := exec.Command("clear")
	stdout, err := cmd.Output()


	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))

}



func checkNetworkRead(err error)bool  {
    if(err == io.EOF){
        fmt.Println("Suppressing EOF ERROR")
        return true
    }else{
        checkError(err)
    }
    return false
}