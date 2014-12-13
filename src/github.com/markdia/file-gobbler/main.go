package main

import (
	//"bufio"
	"fmt"
	//"io"
	//"bytes"
	"encoding/binary"
	//"io/ioutil"
	"log"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func main() {
	p := fmt.Println

	//testing a simple readfile and dump
	/*
		Read in Entire File

		dat, err := ioutil.ReadFile("./tmp/WhoopStrapProcessed_10Minute.bin")
		check(err)
		sizeOfBin := binary.Size(dat)
		p("size of bin")
		p(sizeOfBin)
		p("\n")
	*/

	file, err := os.Open("./tmp/WhoopStrapProcessed_1Day.bin")
	check(err)
	defer file.Close()
	sizeOfBin := binary.Size(file)
	p("size of bin")
	p(sizeOfBin)
	p("\n")
	check(err)
	now := time.Now()
	p(now)

	sliceOfStructs := make([]WhoopStrapData, 172801)
	i := 0
	for {
		var structuredFileData WhoopStrapData
		readErr := binary.Read(file, binary.LittleEndian, &structuredFileData)
		if readErr != nil {
			fmt.Println("binary.Read failed:", readErr)
			break
		}
		sliceOfStructs = append(sliceOfStructs, structuredFileData)
		//fmt.Print(structuredFileData)
		p("\n")
		i++
	}
	p("number of lines processed:")
	p(len(sliceOfStructs))
	p("\n")
	filetest := time.Now()
	p(filetest)
	diff := filetest.Sub(now)
	p(diff)
}
