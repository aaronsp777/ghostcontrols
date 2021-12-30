package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var (
	action    = flag.String("action", "toggle", "One of (toggle, test, party, vacation)")
	id        = flag.Int64("id", 123456, "Transmitter ID")
	button_id = flag.Int64("button_id", 1, "(1:primary, 0:secondary)")
	version   = flag.Int64("version", 1, "(1: remote, 2:keypad)")
	dry_run   = flag.Bool("dry_run", false, "skip calling sendook")
)

func decodeAction(action string) (option, command int64, err error) {
	switch action {
	case "toggle":
		command = 3
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
	code |= *button_id << 23
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
		cmd := exec.Command("sudo", "sendook", "-1", "250", "-0", "250", "-r", "10", "-p", "40000", bits)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
