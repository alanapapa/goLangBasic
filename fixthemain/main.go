package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

var OPEN bool = true
var CLOSE bool = false

func OpenDoor(door *Door) {
	PrintStr("Door Opening...")
	PrintStr("\n")
	door.state = OPEN
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) == true {
		OpenDoor(door)
		return
	}
	if IsDoorOpen(door) == true {
		CloseDoor(door)
		return
	}

}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

func CloseDoor(door *Door) {
	PrintStr("Door Closing...")
	PrintStr("\n")
	door.state = CLOSE
}

func IsDoorOpen(door *Door) bool {
	PrintStr("is the Door opened ?")
	PrintStr("\n")
	if door.state == OPEN {
		return true
	}
	return false
}

func IsDoorClose(door *Door) bool {
	PrintStr("is the Door closed ?")
	PrintStr("\n")
	if door.state == CLOSE {
		return true
	}
	return false
}
