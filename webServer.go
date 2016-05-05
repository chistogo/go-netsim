package main

import(
    "net/http"
    "fmt"
    "os/exec"
    //"os"
    "io/ioutil"
    //"encoding/json"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Inside handler")
    fileName := "graph.json"
    
   
    //Read in JSON file to create the router Struct
    data, err := ioutil.ReadFile(fileName)
    checkError(err)
    //prints the object in the terminal
    fmt.Println(string(data))
    //prints the object on the web page
    fmt.Fprintln(w, string(data))
        
    
}
func main(){
    clear()
    println("ʕ◔ϖ◔ʔ  Welcome to the GO NetSim, Web Process!!!  ʕ◔ϖ◔ʔ")
    http.HandleFunc("/", handler) // redirect all urls to the handler function
    http.ListenAndServe("localhost:8080", nil) 
    
    
    
}//end main

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

//General Error Catching
func checkError(err error)  {
    if err != nil {
        log.Fatal(err)
    }
}