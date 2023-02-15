package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	//cmd := exec.Command("./main")
	//go receiver(20005)
	go startTimer()

	cmd := exec.Command("gnome-terminal", "--", "go", "run", "main.go")

	if err := cmd.Start(); err != nil {
		fmt.Println(err)

	}

	//fmt.Println("Output: ", string(out))

	select {}

}

func startTimer() {

	timer1 := time.NewTimer(1 * time.Second)

	<-timer1.C
	fmt.Println("Timer b fired")

	os.Exit(4)

}

func receiver(port int) {
	stopRecive := 10 * time.Second
	timer1 := time.NewTimer(stopRecive)

	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: port, Zone: ""})
	defer ServerConn.Close()
	buf := make([]byte, 1024)
	for {

		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			panic("panic")
		}
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		timer1.Reset(stopRecive)

	}
}
