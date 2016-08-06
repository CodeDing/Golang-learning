package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	buf := make([]byte, 512)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		fmt.Println(n)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	if _, e := os.Stat("name"); e != nil {
		os.Mkdir("name", 0755)
	} else {
		errors.New("Mkdir error")
	}
	cmd := exec.Command("/bin/ls", "-l")
	err := cmd.Run()
	if err != nil {
		fmt.Println("OK")
	}
}
