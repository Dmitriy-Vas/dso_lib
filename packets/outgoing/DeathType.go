package outgoing

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *DeathTypePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DeathTypePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DeathTypePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DeathTypePacket) SetSend(value bool) {
	packet.Send = value
}

type DeathTypePacket struct {
	ID    int64
	Send  bool
	Index byte
}

func (packet *DeathTypePacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *DeathTypePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Index, b.Index())
}
