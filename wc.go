package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	return bytes, err
}

func wc(file string) (int,int,int,string){
	var byteCount int = 0
	var wordCount int = 0
	var lineCount int = 0

	bytes, _ := readFile(file)

	byteCount = len(bytes)
	var lc byte = 0
	for _, b := range bytes {
		if lc!=('\r') && (b == byte('\n') || b == byte('\r')){
			lineCount++
		}
		if lc <= 0x20 && b > 0x20 {
			wordCount++
		}
		lc = b
	}

	return lineCount,wordCount,byteCount, file
}
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a file")
	}



	lc, wc, bc, fn := wc(args[1])
	fmt.Printf("%d %d %d %s\n", lc, wc, bc, fn)
}