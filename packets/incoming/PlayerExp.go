package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerExpPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerExpPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerExpPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerExpPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerExpPacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 int32
	Variable2 int64
	Variable3 int64
	Variable4 byte
	Variable5 int32
	Variable6 byte
	Variable7 int64
}

func (packet *PlayerExpPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *PlayerExpPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable3, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable5, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable6, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable7, b.Index())
}
