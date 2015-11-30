package main 

import (
    "flag"
    "fmt"
    "os"
    //s "strings"
    //"os"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var fileName *string = flag.String("f", "", "The file to scan")

func main() {
	fmt.Println("LFMGREP");
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    
    //argsWithoutProg := os.Args[1:]
    //fmt.Println(argsWithoutProg)
    //filePtr := flag.String("file", "", "The file to scan")
	if fi, err := os.Stat(*fileName); err == nil {
	  fmt.Println("Processing File:", *fileName)
	}
    
    
    
}


