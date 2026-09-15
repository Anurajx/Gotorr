package main

import (
	"flag"
	"fmt"
)

const banner = `             ___________                  
   ____   ___\__    ___/_________________ 
  / ___\ /  _ \|    | /  _ \_  __ \_  __ \
 / /_/  >  <_> )    |(  <_> )  | \/|  | \/
 \___  / \____/|____| \____/|__|   |__|   
/_____/                                   `

func main() {
	// Define flags
	name := flag.String("name", "World", "Name to greet")
	verbose := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	if *verbose {
		fmt.Println("Running in verbose mode...")
	}
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Println(banner)

}