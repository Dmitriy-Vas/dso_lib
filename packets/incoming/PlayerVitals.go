package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerVitalsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerVitalsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerVitalsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerVitalsPacket) SetSend(value bool) {
	packet.Send = value
}

// TODO finish packet
type PlayerVitalsPacket struct {
	ID        int64
	Send      bool
	PlayerNum uint16
	Variable2 int64
	Variable3 int64
	Variable4 uint16
	Variable5 byte
	Variable6 int32
}

func (packet *PlayerVitalsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadUShort(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadUShort(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerVitalsPacket) Write(b buffer.PacketBuffer) {
	b.WriteUShort(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable3, b.Index())
	b.WriteUShort(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
}
