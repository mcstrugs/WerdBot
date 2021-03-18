package util

import (
	"io"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleBook(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	spaced_words := strings.Split(command, " ")
	var book string
	if len(spaced_words) != 2 {
		s.ChannelMessageSend(m.ChannelID, "please specify one book")
		return
	}
	book = spaced_words[1]
	var r io.Reader
	r, err := os.Open("books/" + spaced_words[1])
	dir, _ := os.ReadDir("books")
	exists := false
	for i := 0; i < len(dir); i++ {
		if strings.Compare(dir[i].Name(), book) == 0 {
			exists = true
		}
	}
	if !exists {
		s.ChannelMessageSend(m.ChannelID, "book does not exist")
		return
	}
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "book does not exist")
		return
	}
	s.ChannelFileSend(m.ChannelID, spaced_words[1], r)
}
