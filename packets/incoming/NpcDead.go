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
	ID         int64
	Send       bool
	Index      int32
	Variable2  int32
	Animation  int32
	Pos        objects.Vector2
	EventCount int32
	Right      string // Name of player to compare
}

func (packet *NpcDeadPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Animation = b.ReadInt(b.Bytes(), b.Index())
	packet.Pos = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	packet.EventCount = b.ReadInt(b.Bytes(), b.Index())
	packet.Right = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *NpcDeadPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Animation, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.Y, b.Index())
	b.WriteInt(b.Bytes(), packet.EventCount, b.Index())
	b.WriteString(b.Bytes(), packet.Right, b.Index())
}
