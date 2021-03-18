package util

import (
	"io"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleBook(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	spaced_words := strings.Split(command, " ")
	if len(spaced_words) != 2 {
		s.ChannelMessageSend(m.ChannelID, "please specify one book")
		return
	}
	book := spaced_words[1]

	dir, _ := os.ReadDir("books")
	exists := false
	for i := 0; i < len(dir); i++ {
		if dir[i].Name() == book {
			exists = true
		}
	}
	if !exists {
		s.ChannelMessageSend(m.ChannelID, "book does not exist")
		return
	}

	var r io.Reader
	r, err := os.Open("books/" + spaced_words[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "book does not exist")
		return
	}

	s.ChannelFileSend(m.ChannelID, spaced_words[1], r)
}
