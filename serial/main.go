package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main(){
		
	c:= &serial.Config{Name:"COM8",Baud:9600}
	ser, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	//buf := make([]byte, 128)
	for{
		//n,err := ser.Read(buf)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//log.Printf("%q", buf[:n])
		reader := bufio.NewReader(ser)
		
		reply, Err := reader.ReadBytes('\x0a')
		fmt.Printf("%x\n",reply)

		if Err != nil {
			log.Fatal(Err)
		}
	}
	


	
}