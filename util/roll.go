package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func HandleRoll(roll_str string) string {
	spaced_words := strings.Split(roll_str, " ")
	if len(spaced_words) == 2 {
		spaced_nums := strings.Split(spaced_words[1], "d")
		num_dice, err := strconv.ParseInt(spaced_nums[0], 10, 64)
		if err != nil {
			return "number of dice must be integer"
		}
		dice_size, err2 := strconv.ParseInt(spaced_nums[1], 10, 64)
		if err2 != nil {
			return "dice size must be integer"
		}
		message := ":game_die:\n**" + spaced_words[1] + "**: "
		total := 0
		for i := 0; i < int(num_dice); i++ {
			rand.Seed(time.Now().UnixNano())
			roll := rand.Int()%int(dice_size) + 1
			message += fmt.Sprint(roll) + " "
			total += roll
		}
		message += "\n	**Total**: " + fmt.Sprint(total)
		return message
	} else {
		return "!roll requires one argument such as 2d6"
	}
}
