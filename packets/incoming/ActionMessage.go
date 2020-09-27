package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib/objects"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *ActionMessagePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ActionMessagePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ActionMessagePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ActionMessagePacket) SetSend(value bool) {
	packet.Send = value
}

type ActionMessagePacket struct {
	ID   int64
	Send bool

	Variable0  byte
	Variable1  uint16
	Variable2  uint16
	Variable3  string
	Variable4  objects.Vector2
	Variable5  string
	Variable6  byte
	Variable7  byte
	Variable8  bool
	Variable9  byte
	Variable10 byte
}

func (packet *ActionMessagePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadUShort(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadUShort(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable4 = objects.Vector2{
		X: int32(b.ReadByte(b.Bytes(), b.Index())),
		Y: int32(b.ReadByte(b.Bytes(), b.Index())),
	}
	packet.Variable5 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ActionMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Variable0, b.Index())
	b.WriteUShort(b.Bytes(), packet.Variable1, b.Index())
	b.WriteUShort(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4.Y, b.Index())
	b.WriteString(b.Bytes(), packet.Variable5, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable6, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable7, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable8, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable9, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable10, b.Index())
}
