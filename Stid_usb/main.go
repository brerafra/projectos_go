/***************************************************************

Desarollos vergas presenta:

Programa de lectura mifare desfire BRENTHON RAMIREZ FRANCO

con:

Lenguaje de programación GO, hecho para base de datos mysql (mariadb)
lectora stid architecture (protocolo ph5-7aa)

paquetes:

"github.com/tarm/serial" -> para comunicación serial
"github.com/snksoft/crc" -> para obtención de CRC-CCITT-16

***************************************************************/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/snksoft/crc"
	"github.com/tarm/serial"
)

func run_cmd_serial(cmd []byte){
	//rutina de comunicacion serial, este comando
	//no regresa nada
	c:= &serial.Config{Name:"COM6",Baud:38400}
	ser, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := ser.Write(cmd)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(n)

	time.Sleep(100 * time.Millisecond)

	buf := make([]byte, 128)
	n,err = ser.Read(buf)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])

	err = ser.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func armaComando(cmd []byte, cmd_string string){
	//Rutina de armado de comando serial para lectora stid_architect, 
	//Esta rutina calcula CRC-CCITT16 de cmd, asi como crea la cadena final a enviar 
	var cmd_final = []byte("")

	ccittCrc := crc.CalculateCRC(crc.CCITT, []byte(cmd))
	//fmt.Printf("CRC is 0x%04X\n", ccittCrc) // prints "CRC is 0x29B1"
	//fmt.Printf("CRC is 0x%x\n", ccittCrc[0]) // prints "CRC is 0x29B1"

	byte_MIB := byte(ccittCrc)
	//fmt.Printf("0x%x\n",byte_MIB)
	byte_LSB := byte(ccittCrc>>8)
	//fmt.Printf("%x\n",byte_LSB)

	comando_output :=fmt.Sprintf("%s%s","\x02",cmd_string)//,byte_LSB,byte_MIB)
	tamaño := len([]byte(comando_output))
	//fmt.Printf("tamaño: %d \n",tamaño)
	
	cmd_final = []byte(comando_output)
	cmd_final[tamaño-2] = byte_LSB
	cmd_final[tamaño-1] = byte_MIB

	//fmt.Println(cmd_final)
	run_cmd_serial(cmd_final)
}

func run_cmd_serial_retorno(cmd []byte) []byte {
	//rutina de comunicacion serial, este comando
	//no regresa nada
	c:= &serial.Config{Name:"COM6",Baud:38400}
	ser, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := ser.Write(cmd)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(n)

	time.Sleep(200 * time.Millisecond)

	buf := make([]byte, 128)
	n,err = ser.Read(buf)

	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%q", buf[:n])

	err = ser.Close()
	if err != nil {
		log.Fatal(err)
	}
	return buf[:n]
}

func armaComandoRetorno(cmd []byte, cmd_string string) []byte   {
	//Rutina de armado de comando serial para lectora stid_architect, 
	//Esta rutina calcula CRC-CCITT16 de cmd, asi como crea la cadena final a enviar 
	var cmd_final = []byte("")
	

	ccittCrc := crc.CalculateCRC(crc.CCITT, []byte(cmd))
	//fmt.Printf("CRC is 0x%04X\n", ccittCrc) // prints "CRC is 0x29B1"
	//fmt.Printf("CRC is 0x%x\n", ccittCrc[0]) // prints "CRC is 0x29B1"

	byte_MIB := byte(ccittCrc)
	//fmt.Printf("0x%x\n",byte_MIB)
	byte_LSB := byte(ccittCrc>>8)
	//fmt.Printf("%x\n",byte_LSB)

	comando_output :=fmt.Sprintf("%s%s","\x02",cmd_string)//,byte_LSB,byte_MIB)
	tamaño := len([]byte(comando_output))
	//fmt.Printf("tamaño: %d \n",tamaño)
	
	cmd_final = []byte(comando_output)
	cmd_final[tamaño-2] = byte_LSB
	cmd_final[tamaño-1] = byte_MIB

	//fmt.Println(cmd_final)
	return run_cmd_serial_retorno(cmd_final)
}

func comando_output(cmd string){
	//Rutina de control de leds y buzzer de lectora STID_ARCH
	//Hay que ingresar cada color y evento con un string  para que entre en el
	//switch, esta rutina invoca a las rutinas de control de serial
	switch cmd {
	case "ROSA":
		cmd_rosa_static_a:="\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\xE1\x00\x98\xFF\x00\x00\x00"
		var cmd_rosa_static = []byte("\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\xE1\x00\x98\xFF\x00")

		armaComando(cmd_rosa_static, cmd_rosa_static_a)

	case "WHITE":
		cmd_white_static_a:="\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\x3E\x60\x89\xFF\x00\x00\x00"
		var cmd_white_static = []byte("\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\x3E\x60\x89\xFF\x00")

		armaComando(cmd_white_static, cmd_white_static_a)

	case "OK":
		cmd_ok_a :="\x00\x0B\x00\x00\x00\x00\x00\x07\xAA\x55\x00\x03\x01\x02\x01\x30\x68\x00\x00"
		var cmd_ok = []byte("\x00\x0B\x00\x00\x00\x00\x00\x07\xAA\x55\x00\x03\x01\x02\x01\x30\x68")

		armaComando(cmd_ok, cmd_ok_a)

	case "NOK":

		cmd_nok_a :="\x00\x0B\x00\x00\x00\x00\x00\x07\xAA\x55\x00\x03\x02\x01\x01\x3C\x6B\x00\x00"
		var cmd_nok = []byte("\x00\x0B\x00\x00\x00\x00\x00\x07\xAA\x55\x00\x03\x02\x01\x01\x3C\x6B")

		armaComando(cmd_nok, cmd_nok_a)
		armaComando(cmd_nok, cmd_nok_a)

	case "YELLOW":
		cmd_yellow_static_a:="\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\xff\x7e\x00\xFF\x00\x00\x00"
		var cmd_yellow_static = []byte("\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\xff\x7e\x00\xFF\x00")

		armaComando(cmd_yellow_static, cmd_yellow_static_a)

	case "BLUE":
		cmd_blue_static_a:="\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\x00\x00\x80\xFF\x00\x00\x00"
		var cmd_blue_static = []byte("\x00\x0E\x00\x00\x00\x00\x00\x50\xAA\x55\x00\x06\x80\x00\x00\x80\xFF\x00")

		armaComando(cmd_blue_static, cmd_blue_static_a)
	}
}

func lectura_tarjeta(){
	var cadSerialRx []byte
	var lenCadRx int
	cmd_auth_desfire_a := "\x00\x0C\x00\x00\x00\x02\x00\x0B\xAA\x55\x00\x04\x01\x00\x00\x00\x00\x00"
	var cmd_auth_desfire = []byte("\x00\x0C\x00\x00\x00\x02\x00\x0B\xAA\x55\x00\x04\x01\x00\x00\x00")

	cmd_select_app_a := "\x00\x0B\x00\x00\x00\x02\x00\x5A\xAA\x55\x00\x03\x02\x00\xF5\xE3\x4B\x00\x00"
	var cmd_select_app = []byte("\x00\x0B\x00\x00\x00\x02\x00\x5A\xAA\x55\x00\x03\x02\x00\xF5\xE3\x4B")

	cmd_read_data_personal_a := "\x00\x10\x00\x00\x00\x02\x00\xBD\xAA\x55\x00\x08\x00\x01\x00\x00\x00\x20\x00\x00\x18\x29\x00\x00"
	var cmd_read_data_personal = []byte("\x00\x10\x00\x00\x00\x02\x00\xBD\xAA\x55\x00\x08\x00\x01\x00\x00\x00\x20\x00\x00\x18\x29")

	cadSerialRx = armaComandoRetorno(cmd_auth_desfire, cmd_auth_desfire_a)
	//fmt.Printf("Respuesta: %q \n", cadSerialRx)

	lenCadRx = len(cadSerialRx)
	
	if(lenCadRx>=10){
		if cadSerialRx[9]=='\x02'{
			//fmt.Println("Tarjeta Authenticada")

			//inicia proceso de lectura_tarjeta

			//Paso 3 seleccionar app a usar, en el caso de acceso es la F50002
			cadSerialRx = armaComandoRetorno(cmd_select_app, cmd_select_app_a)
			//fmt.Printf("Respuesta: %q \n", cadSerialRx)
			lenCadRx = len(cadSerialRx)
			if(lenCadRx>=10){
				if (cadSerialRx[9]=='\x02' &&  cadSerialRx[10]=='\x00'){
					//fmt.Println("Aplicacion encontrada")

					cadSerialRx = armaComandoRetorno(cmd_auth_desfire, cmd_auth_desfire_a)
					//fmt.Printf("Respuesta: %q \n", cadSerialRx)

					//\x02\x00\x26\x00\x00\x00\xBD\x00[\x20][\x16\x6F\x0C\x6E\xEC\x3D\x99\xC6\x7A\xDD\x67\x48\xC1\xAA\xA5\x9F\x9E\x1C\x2C\x37\xAC\x9B\xBF\x54\x88\x53\x6E\x1D\xA0\xA3\xF3\x30]\x02\x00\x9A\xCF
					
					cadSerialRx = armaComandoRetorno(cmd_read_data_personal, cmd_read_data_personal_a)
					//fmt.Printf("Respuesta: %q \n", cadSerialRx)
					lenCadRx = len(cadSerialRx)
					if(lenCadRx>=10){
						len_datos := cadSerialRx[8]
						//fmt.Println(len_datos)

						datos_leidos := cadSerialRx[9:9+len_datos]
						fmt.Printf("Respuesta: %q \n",datos_leidos)

						folio_tarjeta := datos_leidos[:8]
						fmt.Printf("Folio: %q \n",folio_tarjeta)
					}

				}else{
					fmt.Println("Aplicacion no encontrada")
				}
			}
		}
	}
}

func mifare() {

	var cadSerialRx []byte
	var uidDesfire = []byte ("\x00\x00\x00\x00\x00\x00\x00")

	cmd_scan_a:="\x00\x08\x00\x00\x00\x00\x00\x02\xAA\x55\x00\x00\x00\x00"	
	var cmd_scan = []byte("\x00\x08\x00\x00\x00\x00\x00\x02\xAA\x55\x00\x00")
	
	for {
		cadSerialRx = armaComandoRetorno(cmd_scan, cmd_scan_a)
		//fmt.Printf("Respuesta: %q \n", cadSerialRx)
		presenciaTarjeta:=cadSerialRx[9]
		if presenciaTarjeta == '\x01'{
			tipoTarjeta:=cadSerialRx[10]
			if tipoTarjeta == '\x02'{
				uidTarjeta:=cadSerialRx[14:21]
				//fmt.Println("Tarjeta Desfire detectada")
				//fmt.Printf("UID: %x \n",uidTarjeta)
				
				uidlen := len(uidTarjeta)
				for posicion,letra:= range uidTarjeta{
					uidDesfire[uidlen-posicion-1] = letra
					//fmt.Println(posicion,letra)
				}
				fmt.Printf("UID: %x \n",uidDesfire)

				lectura_tarjeta()
				
			} else {
				fmt.Println("Otro tipo de tarjeta detectada")
			}
		}
		
		//fmt.Printf("%x \n",cadSerialRx)
		time.Sleep(1*time.Second)
	}

}

func main() {
	comando_output("YELLOW")

	go mifare()

	for{
	}
}