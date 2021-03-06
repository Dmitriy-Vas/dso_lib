package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *TradeUpdatePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TradeUpdatePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TradeUpdatePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TradeUpdatePacket) SetSend(value bool) {
	packet.Send = value
}

type TradeUpdatePacket struct {
	ID         int64
	Send       bool
	Variable0  int64
	Variable1  int32
	Variable2  int32
	Variable3  byte
	Variable4  byte
	Variable5  int32
	Variable6  int32
	Variable7  byte
	Variable8  byte
	Variable9  int64
	Variable10 int64
}

func (packet *TradeUpdatePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *TradeUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable3, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable7, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable8, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable9, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable10, b.Index())
}
