package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *AttackPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *AttackPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *AttackPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *AttackPacket) SetSend(value bool) {
	packet.Send = value
}

type AttackPacket struct {
	ID        int64
	Send      bool
	Player    uint16
	Attacking byte
	FileName  string
}

func (packet *AttackPacket) Read(b buffer.PacketBuffer) {
	packet.Player = b.ReadUShort(b.Bytes(), b.Index())
	packet.Attacking = b.ReadByte(b.Bytes(), b.Index())
	packet.FileName = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *AttackPacket) Write(b buffer.PacketBuffer) {
	b.WriteUShort(b.Bytes(), packet.Player, b.Index())
	b.WriteByte(b.Bytes(), packet.Attacking, b.Index())
	b.WriteString(b.Bytes(), packet.FileName, b.Index())
}
