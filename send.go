package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

var (
	action    = flag.String("action", "toggle", "One of (toggle, open, test, party, vacation)")
	id        = flag.Int64("id", 123456, "Transmitter ID")
	button_id = flag.Int64("button_id", 2, "(2:primary, 1:secondary, 0:none)")
	version   = flag.Int64("version", 1, "(1: remote, 2:keypad, 9:sensor)")
	dry_run   = flag.Bool("dry_run", false, "skip calling sendook")
	count     = flag.Int64("count", 10, "Number of times to transmit the code")
)

func decodeAction(action string) (option, command int64, err error) {
	switch action {
	case "toggle":
		command = 3
	case "open":
		command = 1
	case "party":
		option = 8
	case "vacation":
		option = 9
	case "test":
		option = 15
	default:
		return 0, 0, fmt.Errorf("Unknown action: %q", action)
	}
	return
}

func codeFromFlags() (int64, error) {
	option, command, err := decodeAction(*action)
	if err != nil {
		return 0, err
	}
	// TODO : range checks
	var code int64
	code |= *version << 38
	code |= option << 34
	code |= command << 30
	code |= *button_id << 22
	code |= *id
	return code, nil
}

func toBits(code int64) (out string) {
	v := code
	for i := 0; i < 42; i++ {
		if v%2 == 0 {
			out = "1000" + out
		} else {
			out = "1110" + out
		}
		v /= 2
	}
	return out
}

func main() {
	flag.Parse()
	code, err := codeFromFlags()
	if err != nil {
		log.Fatal(err)
	}
	bits := toBits(code)
	if *dry_run {
		fmt.Printf("Code: {42}%011x\n", code*4)
		fmt.Printf("Binary: %b\n", code)
		fmt.Printf("After pwm: %q\n", bits)
	} else {
		cmd := exec.Command("sudo", "sendook", "-1", "250", "-0", "250", "-r", strconv.FormatInt(*count, 10), "-p", "40000", bits)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
