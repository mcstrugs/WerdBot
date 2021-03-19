package util

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sahilm/fuzzy"
)

func HandleBook(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	spaced_words := strings.Split(command, " ")
	if len(spaced_words) != 2 {
		s.ChannelMessageSend(m.ChannelID, "please specify one book")
		return
	}
	book := spaced_words[1]

	dir, _ := os.ReadDir("books")
	var book_names []string
	for i := 0; i < len(dir); i++ {
		book_names = append(book_names, dir[i].Name())
	}
	fmt.Println(book_names)
	matches := fuzzy.Find(book, book_names)
	if len(matches) == 0 {
		s.ChannelMessageSend(m.ChannelID, "no books matching: "+book)
		return
	}
	best_match := matches[0].Str

	var r io.Reader
	r, err := os.Open("books/" + best_match)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "book does not exist")
		return
	}

	s.ChannelFileSend(m.ChannelID, best_match, r)
}
