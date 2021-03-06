package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *CooldownPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CooldownPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CooldownPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CooldownPacket) SetSend(value bool) {
	packet.Send = value
}

type CooldownPacket struct {
	ID   int64
	Send bool

	Num       byte
	Icon      uint16
	SpellTime uint16
	Variable4 bool
}

func (packet *CooldownPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadByte(b.Bytes(), b.Index())
	packet.Icon = b.ReadUShort(b.Bytes(), b.Index())
	packet.SpellTime = b.ReadUShort(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *CooldownPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Num, b.Index())
	b.WriteUShort(b.Bytes(), packet.Icon, b.Index())
	b.WriteUShort(b.Bytes(), packet.SpellTime, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable4, b.Index())
}
