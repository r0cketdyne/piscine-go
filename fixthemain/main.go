package main

import "github.com/01-edu/z01"

var (
	CLOSE = true
	OPEN  = false
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

type Door struct {
	state bool
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
