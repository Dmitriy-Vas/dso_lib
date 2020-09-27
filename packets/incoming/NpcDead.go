package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib/objects"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *NpcDeadPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *NpcDeadPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *NpcDeadPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *NpcDeadPacket) SetSend(value bool) {
	packet.Send = value
}

type NpcDeadPacket struct {
	ID   int64
	Send bool

	Index      byte
	Variable2  uint16
	Animation  uint16
	Pos        objects.Vector2
	EventCount uint16
	Right      string // Name of player to compare
}

func (packet *NpcDeadPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadUShort(b.Bytes(), b.Index())
	packet.Animation = b.ReadUShort(b.Bytes(), b.Index())
	packet.Pos = objects.Vector2{
		X: int32(b.ReadByte(b.Bytes(), b.Index())),
		Y: int32(b.ReadByte(b.Bytes(), b.Index())),
	}
	packet.EventCount = b.ReadUShort(b.Bytes(), b.Index())
	packet.Right = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *NpcDeadPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Index, b.Index())
	b.WriteUShort(b.Bytes(), packet.Variable2, b.Index())
	b.WriteUShort(b.Bytes(), packet.Animation, b.Index())
	b.WriteByte(b.Bytes(), byte(packet.Pos.X), b.Index())
	b.WriteByte(b.Bytes(), byte(packet.Pos.Y), b.Index())
	b.WriteUShort(b.Bytes(), packet.EventCount, b.Index())
	b.WriteString(b.Bytes(), packet.Right, b.Index())
}
