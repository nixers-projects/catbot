package main

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
)

func main () {

    if (len(os.Args) != 2) { fmt.Println("args, fix them"); os.Exit(2) }
    url := os.Args[1]

    r, err := http.Get(url)
    if err != nil { os.Exit(1) }
    p, err := ioutil.ReadAll(r.Body)
    if err != nil { os.Exit(1) }
    r.Body.Close()
    re := regexp.MustCompile("<title>.*?</title>")
    buf := re.FindString(string(p))
//    fmt.Println(buf)
    if (len(buf) > 16) {
        fmt.Println(buf[7:len(buf)-8])
    }
}


