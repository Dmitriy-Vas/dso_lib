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
	PlayerNum uint16
	Type      byte
	DrillMod  dso_lib.DrillModRec
	MoleHP    byte
}

func (packet *TopoDrillPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadUShort(b.Bytes(), b.Index())
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	switch {
	case packet.Type == 1:
		packet.DrillMod = dso_lib.DrillModRec{Dir: b.ReadByte(b.Bytes(), b.Index())}
	case packet.Type-2 <= 1:
		packet.DrillMod = dso_lib.DrillModRec{DirStar: b.ReadByte(b.Bytes(), b.Index())}
		packet.MoleHP = b.ReadByte(b.Bytes(), b.Index())
	}
}

func (packet *TopoDrillPacket) Write(b buffer.PacketBuffer) {
	b.WriteUShort(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	switch packet.Type {
	case 2, 3:
		b.WriteByte(b.Bytes(), packet.DrillMod.DirStar, b.Index())
	case 1:
		b.WriteByte(b.Bytes(), packet.DrillMod.Dir, b.Index())
		b.WriteByte(b.Bytes(), packet.MoleHP, b.Index())
	}
}
