package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerBuddiesPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerBuddiesPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerBuddiesPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerBuddiesPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerBuddiesPacket struct {
	ID      int64
	Send    bool
	Buddies []dso_lib.PlayerBuddiesRec
}

func (packet *PlayerBuddiesPacket) Read(b buffer.PacketBuffer) {
	packet.Buddies = make([]dso_lib.PlayerBuddiesRec, 35)
	for i := range packet.Buddies {
		name := b.ReadString(b.Bytes(), b.Index(), 0)
		buddie := dso_lib.PlayerBuddiesRec{
			Name: name,
		}
		if name != "" {
			buddie.Classes = b.ReadByte(b.Bytes(), b.Index())
			buddie.Sex = b.ReadByte(b.Bytes(), b.Index())
			buddie.Sprite = b.ReadUShort(b.Bytes(), b.Index())
			buddie.Hair = b.ReadUShort(b.Bytes(), b.Index())
			buddie.Paperdoll = b.ReadString(b.Bytes(), b.Index(), 0)
			buddie.HairTint = b.ReadString(b.Bytes(), b.Index(), 0)
			buddie.FriendDate = b.ReadString(b.Bytes(), b.Index(), 0)
			buddie.Status = b.ReadBool(b.Bytes(), b.Index())
		}
		packet.Buddies[i] = buddie
	}
}

func (packet *PlayerBuddiesPacket) Write(b buffer.PacketBuffer) {
	for _, buddie := range packet.Buddies {
		b.WriteString(b.Bytes(), buddie.Name, b.Index())
		if buddie.Name != "" {
			b.WriteByte(b.Bytes(), buddie.Classes, b.Index())
			b.WriteByte(b.Bytes(), buddie.Sex, b.Index())
			b.WriteUShort(b.Bytes(), buddie.Sprite, b.Index())
			b.WriteUShort(b.Bytes(), buddie.Hair, b.Index())
			b.WriteString(b.Bytes(), buddie.Paperdoll, b.Index())
			b.WriteString(b.Bytes(), buddie.HairTint, b.Index())
			b.WriteString(b.Bytes(), buddie.FriendDate, b.Index())
			b.WriteBool(b.Bytes(), buddie.Status, b.Index())
		}
	}
}
