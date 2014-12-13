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
	now := time.Now()
	p(now)

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

	file, err := os.Open("./tmp/WhoopStrapProcessed_10Minute.bin")
	check(err)
	defer file.Close()
	sizeOfBin := binary.Size(file)
	p("size of bin")
	p(sizeOfBin)
	p("\n")

	check(err)

	/*
		binBuffer := bytes.NewBuffer(dat)
		version, err := binary.ReadUvarint(binBuffer)
		check(err)
		fmt.Printf("new int is %d\n", version)

		flags, err := binary.ReadUvarint(binBuffer)
		check(err)
		fmt.Printf("new flag is %d\n", flags)

		ts, err := binary.ReadUvarint(binBuffer)
		check(err)
		fmt.Printf("new ts is %d\n", ts)

		gsr, err := binary.ReadUvarint(binBuffer)
		check(err)
		fmt.Printf("new gsr is %d\n", gsr)

		temperature, err := binary.ReadUvarint(binBuffer)
		check(err)
		fmt.Printf("new temperature is %d\n", temperature)
	*/

	var structuredFileData WhoopStrapData
	readErr := binary.Read(file, binary.LittleEndian, &structuredFileData)
	if readErr != nil {
		fmt.Println("binary.Read failed:", readErr)
	}
	fmt.Print(structuredFileData)
	p("\n")

	filetest := time.Now()
	p(filetest)
	diff := filetest.Sub(now)
	p(diff)
}
