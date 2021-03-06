package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *OpenShopPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *OpenShopPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *OpenShopPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *OpenShopPacket) SetSend(value bool) {
	packet.Send = value
}

type OpenShopPacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 int32
	Variable2 string
	Variable3 int32
	Variable4 string
	Variable5 int32
	Variable6 int64
	Variable7 int32
	Variable8 int32
	Variable9 int32
}

func (packet *OpenShopPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable5 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *OpenShopPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteString(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteString(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable5, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable6, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
}
