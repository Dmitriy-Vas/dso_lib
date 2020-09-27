package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *BuddiesActionPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *BuddiesActionPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *BuddiesActionPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *BuddiesActionPacket) SetSend(value bool) {
	packet.Send = value
}

type BuddiesActionPacket struct {
	ID   int64
	Send bool

	Action   byte
	BuddieID byte
	Buddie   dso_lib.PlayerBuddiesRec
}

func (packet *BuddiesActionPacket) Read(b buffer.PacketBuffer) {
	packet.Action = b.ReadByte(b.Bytes(), b.Index())
	packet.BuddieID = b.ReadByte(b.Bytes(), b.Index())

	if packet.Action == 1 || packet.Action == 3 { // TODO int to const
		packet.Buddie = dso_lib.PlayerBuddiesRec{
			Name:       b.ReadString(b.Bytes(), b.Index(), 0),
			Classes:    b.ReadByte(b.Bytes(), b.Index()),
			Sex:        b.ReadByte(b.Bytes(), b.Index()),
			Sprite:     b.ReadUShort(b.Bytes(), b.Index()),
			Hair:       b.ReadUShort(b.Bytes(), b.Index()),
			Paperdoll:  b.ReadString(b.Bytes(), b.Index(), 0),
			HairTint:   b.ReadString(b.Bytes(), b.Index(), 0),
			FriendDate: b.ReadString(b.Bytes(), b.Index(), 0),
			Status:     b.ReadBool(b.Bytes(), b.Index()),
		}
	}
}

func (packet *BuddiesActionPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Action, b.Index())
	b.WriteByte(b.Bytes(), packet.BuddieID, b.Index())

	if packet.Action == 1 || packet.Action == 3 { // TODO int to const
		b.WriteString(b.Bytes(), packet.Buddie.Name, b.Index())
		b.WriteByte(b.Bytes(), packet.Buddie.Classes, b.Index())
		b.WriteByte(b.Bytes(), packet.Buddie.Sex, b.Index())
		b.WriteUShort(b.Bytes(), packet.Buddie.Sprite, b.Index())
		b.WriteUShort(b.Bytes(), packet.Buddie.Hair, b.Index())
		b.WriteString(b.Bytes(), packet.Buddie.Paperdoll, b.Index())
		b.WriteString(b.Bytes(), packet.Buddie.HairTint, b.Index())
		b.WriteString(b.Bytes(), packet.Buddie.FriendDate, b.Index())
		b.WriteBool(b.Bytes(), packet.Buddie.Status, b.Index())
	}
}
