package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *DefensaPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DefensaPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DefensaPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DefensaPacket) SetSend(value bool) {
	packet.Send = value
}

type DefensaPacket struct {
	ID        int64
	Send      bool
	PlayerNum uint16
	ShieldHP  uint16
	Status    bool
}

func (packet *DefensaPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadUShort(b.Bytes(), b.Index())
	packet.ShieldHP = b.ReadUShort(b.Bytes(), b.Index())
	packet.Status = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *DefensaPacket) Write(b buffer.PacketBuffer) {
	b.WriteUShort(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteUShort(b.Bytes(), packet.ShieldHP, b.Index())
	b.WriteBool(b.Bytes(), packet.Status, b.Index())
}
