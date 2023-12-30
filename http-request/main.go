package main

import(

    "fmt"        // library that implements formated I/O like C
    "io/ioutil" // implements some I/O utility functions
    "log"       // implements a simple loggin package
    "net/http"  // provides HTTP client and server implementation
)

func main() {

    // here we are requesting (consuming from a website)
    // like a client
    resp, err := http.Get("https://httpbin.org/get")

    // Error checking from the http.Get() request

    // this line is ensuring that the boddy
    // of the http response "resp.Body" 
    // is closed after all operations are done.

    // defer --> is like a "finally" in a try-catch block
    defer resp.Body.Close()

    // this line is reading the data from the response body
    // into a byte slice (bodyBytes)

    // ioutil.ReadAll function read from the provided
    // io.Reader, res.Body in this case, until an error 
    // occurs or the EOF is reached

    bodyBytes, err := ioutil.ReadAll(resp.Body) 
   
    // checking all the errors from the ioutil.ReadAll()
    // operation
    if err != nil {
        log.Fatal(err)
    }

    // creating a string from the bodyBytes slice
    bodyString := string(bodyBytes)

    // printing statements
    fmt.Println("Printing status code \n",
    resp.StatusCode) 

    // printing the header map
    fmt.Println("Printing the header \n",
    resp.Header)

    // Reading the element 
    // "Content-Type" from the map
    fmt.Println("Printing content type header \n",
    resp.Header["Content-Type"])
    fmt.Println("Printing the key element \n", 
    resp.Header["Content-Type"][0])

    // printing the bodystring
    // that is a JSON type
    fmt.Println("Prinitng the body from the response as String \n",
    bodyString)

}
