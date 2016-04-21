package main

import (
    "fmt"
    "log"
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


func main(){
    
   clear()
   fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO NetSim, Router Process!!!  ʕ◔ϖ◔ʔ")
    
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