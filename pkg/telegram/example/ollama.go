package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/telegram"
	"github.com/funtimecoding/go-library/pkg/telegram/constant"
	"github.com/funtimecoding/go-library/pkg/text/multi_line"
)

func Ollama() {
	o := ollama.NewEnvironment()
	t := telegram.NewEnvironment()
	defer t.Close()
	f := constant.Format
	c := environment.Get(constant.ChannelEnvironment, 1)

	l := multi_line.New()
	l.Add(
		"You are a helpful assistant. Below is a set of lines from a chat channel. Respond with a message in a friendly and informative manner. I will worry about the formatting, time and your name.",
	)

	for _, m := range t.MessagesByChannel(c) {
		if false {
			fmt.Println(m.Format(f))
		}

		if m.Text == "" {
			continue
		}

		l.Add(m.FormatChat())
	}

	if l.Count() > 1 {
		g := o.GenerateFast(l.Render())

		if false {
			fmt.Println(g.Text)
			g.Print()
		}

		t.SendMessage(t.ChannelByName(c).Identifier, g.Text)
	} else {
		fmt.Println("Chat empty")
	}

	if false {
		system.KillSignalBlock()
	}
}
