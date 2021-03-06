package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerDataTypePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerDataTypePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerDataTypePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerDataTypePacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerDataTypePacket struct {
	ID         int64
	Send       bool
	Variable0  int64
	Variable1  int32
	Variable2  int32
	Variable3  int32
	Variable4  int32
	Variable5  int32
	Variable6  int32
	Variable7  int32
	Variable8  int32
	Variable9  int32
	Variable10 int32
	Variable11 byte
	Variable12 int64
	Variable13 int32
	Variable14 int32
	Variable15 int64
	Variable16 string
	Variable17 string
	Variable18 int32
	Variable19 int32
	Variable20 byte
	Variable21 bool
}

func (packet *PlayerDataTypePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable11 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable12 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable13 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable14 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable15 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable16 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable17 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable18 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable19 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable20 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable21 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerDataTypePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable10, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable11, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable12, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable13, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable14, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable15, b.Index())
	b.WriteString(b.Bytes(), packet.Variable16, b.Index())
	b.WriteString(b.Bytes(), packet.Variable17, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable18, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable19, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable20, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable21, b.Index())
}
