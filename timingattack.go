package timingattack

import (
	"crypto/subtle"
	"fmt"
	"time"
)

func timeTrack(start time.Time, attempt string) {
	elapsed := time.Since(start)
	fmt.Printf("%s,%v\n", attempt, elapsed.Nanoseconds())
}

var password string

var passwordAsByte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func LoginConstant(password []byte) bool {
	return subtle.ConstantTimeCompare(password, passwordAsByte) == 1
}

func Login(password string) bool {
	return password == "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
}
