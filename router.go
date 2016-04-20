package main

import (
    "fmt"
    "log"
    "os/exec"
    "encoding/json"
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


func main(){
    
   clear()
   fmt.Println("ʕ◔ϖ◔ʔ  Welcome to the GO NetSim, Router Process!!!  ʕ◔ϖ◔ʔ")
    
        testRouter := Router {
            IP: "192.168.1.1",
            Port: 1232,
            Neighbours: []Neighbour{
                
                Neighbour {
                IP: "192.168.1.1",
                Port: 1232,
                Weight:6},
                
                Neighbour {
                IP: "192.168.1.1",
                Port: 1232,
                Weight:6},
                
                Neighbour {
                IP: "192.168.1.1",
                Port: 1232,
                Weight:6},
                
                },
                
          }//end of struct Declaration
        
           b, _ := json.Marshal(testRouter);
           s := string(b)
           
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