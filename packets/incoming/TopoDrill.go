package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *TopoDrillPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TopoDrillPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TopoDrillPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TopoDrillPacket) SetSend(value bool) {
	packet.Send = value
}

type TopoDrillPacket struct {
	ID        int64
	Send      bool
	PlayerNum int32
	Type      int32
	DrillMod  dso_lib.DrillModRec
	MoleHP    byte
}

func (packet *TopoDrillPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadInt(b.Bytes(), b.Index())
	switch packet.Type {
	case 2, 3:
		packet.DrillMod = dso_lib.DrillModRec{Dir: byte(b.ReadInt(b.Bytes(), b.Index()))}
	case 1:
		packet.DrillMod = dso_lib.DrillModRec{DirStar: b.ReadInt(b.Bytes(), b.Index())}
		packet.MoleHP = b.ReadByte(b.Bytes(), b.Index())
	}
}

func (packet *TopoDrillPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Type, b.Index())
	switch packet.Type {
	case 2, 3:
		b.WriteInt(b.Bytes(), packet.DrillMod.DirStar, b.Index())
	case 1:
		b.WriteInt(b.Bytes(), int32(packet.DrillMod.Dir), b.Index())
		b.WriteByte(b.Bytes(), packet.MoleHP, b.Index())
	}
}
