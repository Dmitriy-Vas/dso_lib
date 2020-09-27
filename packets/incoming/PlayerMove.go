package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib/objects"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerMovePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerMovePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerMovePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerMovePacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerMovePacket struct {
	ID        int64
	Send      bool
	Variable0 uint16
	Position  objects.Vector2
	Dir       byte
	Variable1 uint16
}

func (packet *PlayerMovePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadUShort(b.Bytes(), b.Index())
	packet.Position = objects.Vector2{
		X: int32(b.ReadByte(b.Bytes(), b.Index())),
		Y: int32(b.ReadByte(b.Bytes(), b.Index())),
	}
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadUShort(b.Bytes(), b.Index())
	if packet.Variable1 > 0 {
		// TODO finish packet
	}
}

func (packet *PlayerMovePacket) Write(b buffer.PacketBuffer) {
}
