// The part command makes ircflu part a joined channel.
package partCmd

import (
	"fmt"
	"github.com/muesli/ircflu/commands"
	"github.com/muesli/ircflu/msgsystem"
	"github.com/muesli/ircflu/msgsystem/irc"
	_ "log"
	"strings"
)

type PartCommand struct {
	messagesIn  chan msgsystem.Message
	messagesOut chan msgsystem.Message
}

func (h *PartCommand) Name() string {
	return "part"
}

func (h *PartCommand) Parse(msg msgsystem.Message) bool {
	channel := msg.To
	m := strings.Split(msg.Msg, " ")
	cmd := m[0]
	params := strings.Join(m[1:], " ")

	switch cmd {
	case "!part":
		if !msg.Authed {
			r := msgsystem.Message{
				To:  channel,
				Msg: "Security breach. Talk to ircflu admin!",
			}
			h.messagesOut <- r
			return true
		}

		if len(params) > 0 {
			fmt.Println("Parting:", params)

			ircclient := (*msgsystem.SubSystem("irc")).(*irc.IrcSubSystem)
			if ircclient != nil {
				ircclient.Part(params)
			}
		} else {
			r := msgsystem.Message{
				To:  channel,
				Msg: "Usage: !part #chan",
			}
			h.messagesOut <- r
		}
		return true
	}
	return false
}

func (h *PartCommand) Run(channelIn, channelOut chan msgsystem.Message) {
	h.messagesIn = channelIn
	h.messagesOut = channelOut
}

func init() {
	partCmd := PartCommand{}
	commands.RegisterCommand(&partCmd)
}
