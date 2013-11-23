// ircflu's central messaging system. Handles incoming commands- and outgoing
// messages-routing.
package msgsystem

import (
	"fmt"
)

// Interface which all messaging sub-systems need to implement
type MsgSubSystem interface {
	// Name of the sub-system
	Name() string
	// Activate the sub-system using these in&out channels
	Run(channelIn, channelOut chan Message)
}

type Message struct {
	To     []string
	Msg    string
	Source string
	Authed bool
}

var (
	CommandsIn  = make(chan Message)
	MessagesOut = make(chan Message)

	subsystems map[string]*MsgSubSystem = make(map[string]*MsgSubSystem)
)

func init() {
	fmt.Println("Initializing messaging subsystem...")
}

// Sub-systems need to call this method to register themselves
func RegisterSubSystem(system MsgSubSystem) {
	fmt.Println("Registering msg-subsystem:", system.Name())

	subsystems[system.Name()] = &system
}

// Returns sub-system with this name
func GetSubSystem(identifier string) *MsgSubSystem {
	system, ok := subsystems[identifier]
	if ok {
		return system
	}

	return nil
}

// Starts all registered messaging sub-systems
func StartSubSystems() {
	for _, system := range subsystems {
		(*system).Run(CommandsIn, MessagesOut)
	}
}
