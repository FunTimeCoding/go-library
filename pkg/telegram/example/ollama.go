package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/telegram"
	"github.com/funtimecoding/go-library/pkg/telegram/constant"
	"github.com/funtimecoding/go-library/pkg/text/multi_line"
)

func Ollama() {
	b := bolt.New("tmp/telegram.db")
	defer b.Close()

	o := ollama.NewEnvironment()
	t := telegram.NewEnvironment()
	t.SetDatabase(b)
	t.ReadDatabase()

	if true {
		return
	}

	f := constant.Format
	c := environment.Get(constant.ChannelEnvironment, 1)

	l := multi_line.New()
	l.Add(
		"You are a helpful assistant. Below is a conversation with a user in the format 'Timestamp Username: Message'. Please respond in a friendly and informative manner.",
	)

	for _, m := range t.MessagesByChannel(c) {
		if false {
			fmt.Println(m.Format(f))
		}

		if false {
			fmt.Println(m.FormatChat())
		}

		l.Add(m.FormatChat())
	}

	g := o.GenerateFast(l.Render())
	fmt.Println(g.Text)
	g.Print()

	// TODO: Look up channel to reply in
	//t.SendMessage()
}
