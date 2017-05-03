// Auto-generated by avdl-compiler v1.3.13 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/gregor.avdl

package chat1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type GenericPayload struct {
	Action    string    `codec:"Action" json:"Action"`
	InboxVers InboxVers `codec:"inboxVers" json:"inboxVers"`
}

type NewConversationPayload struct {
	Action       string         `codec:"Action" json:"Action"`
	ConvID       ConversationID `codec:"convID" json:"convID"`
	InboxVers    InboxVers      `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate  `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

type NewMessagePayload struct {
	Action       string         `codec:"Action" json:"Action"`
	ConvID       ConversationID `codec:"convID" json:"convID"`
	Message      MessageBoxed   `codec:"message" json:"message"`
	InboxVers    InboxVers      `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate  `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

type ReadMessagePayload struct {
	Action       string         `codec:"Action" json:"Action"`
	ConvID       ConversationID `codec:"convID" json:"convID"`
	MsgID        MessageID      `codec:"msgID" json:"msgID"`
	InboxVers    InboxVers      `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate  `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

type SetStatusPayload struct {
	Action       string             `codec:"Action" json:"Action"`
	ConvID       ConversationID     `codec:"convID" json:"convID"`
	Status       ConversationStatus `codec:"status" json:"status"`
	InboxVers    InboxVers          `codec:"inboxVers" json:"inboxVers"`
	UnreadUpdate *UnreadUpdate      `codec:"unreadUpdate,omitempty" json:"unreadUpdate,omitempty"`
}

type UnreadUpdate struct {
	ConvID         ConversationID `codec:"convID" json:"convID"`
	UnreadMessages int            `codec:"UnreadMessages" json:"UnreadMessages"`
}

type TLFFinalizeUpdate struct {
	FinalizeInfo ConversationFinalizeInfo `codec:"finalizeInfo" json:"finalizeInfo"`
	ConvIDs      []ConversationID         `codec:"convIDs" json:"convIDs"`
	InboxVers    InboxVers                `codec:"inboxVers" json:"inboxVers"`
}

type TLFResolveUpdate struct {
	ConvID    ConversationID `codec:"convID" json:"convID"`
	InboxVers InboxVers      `codec:"inboxVers" json:"inboxVers"`
}

type GregorInterface interface {
}

func GregorProtocol(i GregorInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.gregor",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type GregorClient struct {
	Cli rpc.GenericClient
}
