package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func HandleRand(command string) string {
	spaced_words := strings.Split(command, " ")
	if len(spaced_words) == 2 {
		max, err := strconv.ParseInt(spaced_words[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			return "rand argument must be integer"
		} else {
			rand.Seed(time.Now().UnixNano())
			return fmt.Sprint(rand.Int() % int(max))
		}
	} else {
		return "rand requires one, and only one argument"
	}
}
