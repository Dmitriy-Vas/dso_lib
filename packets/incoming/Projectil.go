package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *ProjectilPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ProjectilPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ProjectilPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ProjectilPacket) SetSend(value bool) {
	packet.Send = value
}

type ProjectilPacket struct {
	ID           int64
	Send         bool
	ProjectilNum uint16
	PlayerNum    byte
	Variable3    uint16
	Combo        byte
	Dir          byte
}

func (packet *ProjectilPacket) Read(b buffer.PacketBuffer) {
	packet.ProjectilNum = b.ReadUShort(b.Bytes(), b.Index())
	packet.PlayerNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadUShort(b.Bytes(), b.Index())
	packet.Combo = b.ReadByte(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ProjectilPacket) Write(b buffer.PacketBuffer) {
	b.WriteUShort(b.Bytes(), packet.ProjectilNum, b.Index())
	b.WriteByte(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteUShort(b.Bytes(), packet.Variable3, b.Index())
	b.WriteByte(b.Bytes(), packet.Combo, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
