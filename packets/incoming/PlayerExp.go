package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerExpPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerExpPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerExpPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerExpPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerExpPacket struct {
	ID               int64
	Send             bool
	Exp              int64
	Num              int64
	PlayerProfession []dso_lib.PlayerProfessionRec
}

func (packet *PlayerExpPacket) Read(b buffer.PacketBuffer) {
	packet.Exp = b.ReadLong(b.Bytes(), b.Index())
	packet.Num = b.ReadLong(b.Bytes(), b.Index())
	packet.PlayerProfession = make([]dso_lib.PlayerProfessionRec, 4)
	for i := range packet.PlayerProfession {
		packet.PlayerProfession[i] = dso_lib.PlayerProfessionRec{
			Level:  b.ReadByte(b.Bytes(), b.Index()),
			EXP:    b.ReadInt(b.Bytes(), b.Index()),
			Points: b.ReadByte(b.Bytes(), b.Index()),
			Times:  b.ReadLong(b.Bytes(), b.Index()),
		}
	}
}

func (packet *PlayerExpPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Exp, b.Index())
	b.WriteLong(b.Bytes(), packet.Num, b.Index())
	for _, profession := range packet.PlayerProfession {
		b.WriteByte(b.Bytes(), profession.Level, b.Index())
		b.WriteInt(b.Bytes(), profession.EXP, b.Index())
		b.WriteByte(b.Bytes(), profession.Points, b.Index())
		b.WriteLong(b.Bytes(), profession.Times, b.Index())
	}
}
