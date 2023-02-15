package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	//cmd := exec.Command("./main")

	timerExpiredCh := make(chan bool)
	go receiver(20005, timerExpiredCh)
	go startTimer()

	//fmt.Println("Output: ", string(out))

	select {}

}

func startTimer() {

	timer1 := time.NewTimer(10 * time.Second)

	<-timer1.C
	fmt.Println("Timer b fired")

	os.Exit(4)

}

func receiver(port int, timerExpiredCh chan<- bool) {

	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: port, Zone: ""})
	defer ServerConn.Close()
	buf := make([]byte, 1024)
	for {

		fmt.Println("WTF")
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			panic("panic")
		}
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		timerExpiredCh <- true

	}
}
