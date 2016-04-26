package main

import (
    "fmt"
    "net"
    "time"
    "log"
    "io/ioutil"
    "os/exec"
    "encoding/json"
    "github.com/twmb/algoimpl/go/graph"
    "strconv"
)
 
//Router This struct represents a router and contains it's properties


type Neighbour struct{
    IP string
    Port int
    Weight int
    
}

// type Neighbours struct{
//     ne
// }

type Router struct{
    IP string
    Port int
    Neighbours []Neighbour
}

type SpanningTree struct{
    
}

func listenForRouter(router Router) {
    println("listening")
    for true {
        In, err := net.Listen("tcp", ":" + strconv.Itoa(router.Port))
        checkError(err)
        for {
            _, err := In.Accept()
            if err != nil {
                checkError(err)
            }
            println("hi")
                    
                    
        }
    }
}


func main(){
    
   clear()
   println("ʕ◔ϖ◔ʔ  Welcome to the GO NetSim, Router Process!!!  ʕ◔ϖ◔ʔ")
   
   //Read in JSON file to create the router Struct
   data, err := ioutil.ReadFile("routerinfo.json")
   checkError(err)
   
   //Create Router Struct and turn json into Struct
   var testRouter Router
   err = json.Unmarshal(data,&testRouter)
   checkError(err)
    
    //Create a Go Routine that listens for other routers that are trying to comunicate
    go listenForRouter(testRouter)

    //Turn our Router to json , this was for testing
    //b, _ := json.Marshal(testRouter);
    // s := string(b)
    
    //fmt.Println(s)
    
    //Create Graph undirected weighted graph to represent the current network
    tree := graph.New(graph.Undirected)
    nodes := make(map[string]graph.Node, 0)
    nodeWeIs := testRouter.IP + ":" + strconv.Itoa(testRouter.Port)
    nodes[nodeWeIs] = tree.MakeNode()
    
    //Interate through the nodes in the router(the one read from json) and add them to the graph
    for i := 0; i < len(testRouter.Neighbours); i++ {
        
        currentProcessingNode := testRouter.Neighbours[i].IP + ":" + strconv.Itoa(testRouter.Neighbours[i].Port)
        nodes[currentProcessingNode] = tree.MakeNode()
        tree.MakeEdgeWeight(nodes[nodeWeIs], nodes[currentProcessingNode], testRouter.Neighbours[i].Weight)
        
    }
    //Set values of all nodes to key???????? Clarification needed.
    for key, node := range nodes {
        *node.Value = key
    }
    
    //Test to find the minimum spanning tree.
    mst := tree.DijkstraSearch(nodes[nodeWeIs])
    
    //Turn that tree to json for sending
    b, _ = json.Marshal(mst);
    s = string(b)
    //Print out the json for testing
    fmt.Println(s)
    
    
    //This function trys to initiates connection to the other routers and updates the tree if they are connected
    boot(testRouter)
    
    //inifinate loop to keep the program running while the go routines do their thing
    for true {
        
    }
    
    }
 func boot(myRouter Router) {
     // for node in neighbors
     //    are you alive
     for i:=0; i < len(myRouter.Neighbours); i++ {
         fmt.Println(myRouter.Neighbours[i].IP)
         _, err:= net.DialTimeout("tcp", myRouter.Neighbours[i].IP + ":" + strconv.Itoa(myRouter.Neighbours[i].Port), time.Duration(1) * time.Second)
         //checkError(err)
         if(err != nil && err.(net.Error).Timeout()) {
             println("ʕ◔ϖ◔ʔ halp, we timed the fuck out ʕ◔ϖ◔ʔ")
         }
     }
        
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