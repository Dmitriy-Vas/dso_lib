package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PartyUpdatePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PartyUpdatePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PartyUpdatePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PartyUpdatePacket) SetSend(value bool) {
	packet.Send = value
}

type PartyUpdatePacket struct {
	ID    int64
	Send  bool
	Num   int32
	Party dso_lib.PartyRec
}

func (packet *PartyUpdatePacket) Read(b buffer.PacketBuffer) {
	if packet.Num = b.ReadInt(b.Bytes(), b.Index()); packet.Num > 0 {
		party := dso_lib.PartyRec{
			Leader:      b.ReadInt(b.Bytes(), b.Index()),
			Type:        b.ReadByte(b.Bytes(), b.Index()),
			MemberCount: b.ReadByte(b.Bytes(), b.Index()),
		}
		party.Member = make([]dso_lib.MemberPartyRec, party.MemberCount)
		for i := range party.Member {
			party.Member[i].Index = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Name = b.ReadString(b.Bytes(), b.Index(), 0)
			party.Member[i].Map = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Classes = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Level = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Sprite = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Hair = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].HairTint = b.ReadString(b.Bytes(), b.Index(), 0)
			party.Member[i].Helmet = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Mask = b.ReadInt(b.Bytes(), b.Index())
		}
		party.ShareItems = b.ReadBool(b.Bytes(), b.Index())
		packet.Party = party
	}
}

func (packet *PartyUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	if packet.Num > 0 {
		b.WriteInt(b.Bytes(), packet.Party.Leader, b.Index())
		b.WriteByte(b.Bytes(), packet.Party.Type, b.Index())
		b.WriteByte(b.Bytes(), packet.Party.MemberCount, b.Index())
		for _, member := range packet.Party.Member {
			b.WriteInt(b.Bytes(), member.Index, b.Index())
			b.WriteString(b.Bytes(), member.Name, b.Index())
			b.WriteInt(b.Bytes(), member.Map, b.Index())
			b.WriteInt(b.Bytes(), member.Classes, b.Index())
			b.WriteInt(b.Bytes(), member.Level, b.Index())
			b.WriteInt(b.Bytes(), member.Sprite, b.Index())
			b.WriteInt(b.Bytes(), member.Hair, b.Index())
			b.WriteString(b.Bytes(), member.HairTint, b.Index())
		}
		b.WriteBool(b.Bytes(), packet.Party.ShareItems, b.Index())
	}
}
