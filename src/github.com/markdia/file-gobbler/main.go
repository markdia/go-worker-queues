package main

import (
	//"bufio"
	"fmt"
	//"io"
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	//"os"
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
	dat, err := ioutil.ReadFile("./tmp/WhoopStrapProcessed_1Day.bin")
	check(err)
	binBuffer := bytes.NewBuffer(dat)
	version, err := binary.ReadUvarint(binBuffer)
	check(err)
	fmt.Printf("new int is %d\n", version)
	//fmt.Print(string(dat))

	//file, err := os.OpenFile("./tmp/WhoopStrapProcessed_10Minute.bin")
	//check(err)
	//defer file.Close()
	//dataChunk1 := make([]byte, 39)
	//countOfReadBytes, err := binary.Read(dataChunk1)
	//check(err)

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

	filetest := time.Now()
	p(filetest)
	diff := filetest.Sub(now)
	p(diff)
}
