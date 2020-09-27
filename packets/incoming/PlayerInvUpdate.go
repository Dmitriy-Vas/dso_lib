package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerInvUpdatePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerInvUpdatePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerInvUpdatePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerInvUpdatePacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerInvUpdatePacket struct {
	ID      int64
	Send    bool
	Item    dso_lib.InvItemRec
	InvSlot byte
	PickUp  bool
}

func (packet *PlayerInvUpdatePacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.Item = dso_lib.InvItemRec{
		Num:   b.ReadUShort(b.Bytes(), b.Index()),
		Value: b.ReadLong(b.Bytes(), b.Index()),
		Slot:  b.ReadByte(b.Bytes(), b.Index()),
	}
	if dso_lib.GetItemInt(packet.Item.Num, dso_lib.ItemEnhancement) > 0 {
		packet.Item.Stat = make([]byte, 6) // TODO int to const
		for i := range packet.Item.Stat {
			packet.Item.Stat[i] = b.ReadByte(b.Bytes(), b.Index())
		}
	}
	packet.PickUp = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerInvUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteUShort(b.Bytes(), packet.Item.Num, b.Index())
	b.WriteLong(b.Bytes(), packet.Item.Value, b.Index())
	b.WriteByte(b.Bytes(), packet.Item.Slot, b.Index())
	if dso_lib.GetItemInt(packet.Item.Num, dso_lib.ItemEnhancement) > 0 {
		for _, stat := range packet.Item.Stat {
			b.WriteByte(b.Bytes(), stat, b.Index())
		}
	}
	b.WriteBool(b.Bytes(), packet.PickUp, b.Index())
}
