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
   
   data, err := ioutil.ReadFile("routerinfo.json")
   checkError(err)
   var testRouter Router
   
   err = json.Unmarshal(data,&testRouter)
   checkError(err)
    /*
        testRouter := Router {
            IP: "192.168.1.1",
            Port: 1232,
            Neighbours: []Neighbour{
                
                Neighbour {
                IP: "192.168.1.2",
                Port: 1233,
                Weight:1},
                
                Neighbour {
                IP: "192.168.1.3",
                Port: 1234,
                Weight:2},
                
                Neighbour {
                IP: "192.168.1.5",
                Port: 1235,
                Weight:3},
                
                },
                
          }//end of struct Declaration
          */
          go listenForRouter(testRouter)
        
           b, _ := json.Marshal(testRouter);
           s := string(b)
           
           //fmt.Println(s)
           
           tree := graph.New(graph.Undirected)
           nodes := make(map[string]graph.Node, 0)
           nodeWeIs := testRouter.IP + ":" + strconv.Itoa(testRouter.Port)
           nodes[nodeWeIs] = tree.MakeNode()
           for i := 0; i < len(testRouter.Neighbours); i++ {
               
               currentProcessingNode := testRouter.Neighbours[i].IP + ":" + strconv.Itoa(testRouter.Neighbours[i].Port)
               nodes[currentProcessingNode] = tree.MakeNode()
               tree.MakeEdgeWeight(nodes[nodeWeIs], nodes[currentProcessingNode], testRouter.Neighbours[i].Weight)
               
           }
           
           for key, node := range nodes {
                *node.Value = key
            }
           
           mst := tree.DijkstraSearch(nodes[nodeWeIs])
           
           b, _ = json.Marshal(mst);
           s = string(b)
           
           fmt.Println(s)
           boot(testRouter)
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

func clear(){

	cmd := exec.Command("clear")
	stdout, err := cmd.Output()


	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))

}