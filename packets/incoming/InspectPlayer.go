package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *InspectPlayerPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *InspectPlayerPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *InspectPlayerPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *InspectPlayerPacket) SetSend(value bool) {
	packet.Send = value
}

type InspectPlayerPacket struct {
	ID   int64
	Send bool

	Type                uint16
	PlayerName          string
	InspectPlayerInv    []dso_lib.InvItemRec
	InspectPlayerSkulls []int64
}

func (packet *InspectPlayerPacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadUShort(b.Bytes(), b.Index())
	packet.PlayerName = b.ReadString(b.Bytes(), b.Index(), 0)
	switch {
	case packet.Type == 1:
		packet.ReadInspectPlayers(240, b)
		packet.InspectPlayerSkulls = []int64{
			int64(b.ReadInt(b.Bytes(), b.Index())),
		}
	case packet.Type != 0:
		packet.ReadInspectPlayers(49, b)
		packet.InspectPlayerSkulls = make([]int64, 2)
		for i := range packet.InspectPlayerSkulls {
			packet.InspectPlayerSkulls[i] = b.ReadLong(b.Bytes(), b.Index())
		}
	}
}

func (packet *InspectPlayerPacket) Write(b buffer.PacketBuffer) {
	b.WriteUShort(b.Bytes(), packet.Type, b.Index())
	b.WriteString(b.Bytes(), packet.PlayerName, b.Index())

	for _, inv := range packet.InspectPlayerInv {
		b.WriteUShort(b.Bytes(), inv.Num, b.Index())
		b.WriteLong(b.Bytes(), inv.Value, b.Index())
	}
	switch {
	case packet.Type == 1:
		b.WriteInt(b.Bytes(), int32(packet.InspectPlayerSkulls[0]), b.Index())
	case packet.Type != 0:
		for _, skull := range packet.InspectPlayerSkulls {
			b.WriteLong(b.Bytes(), skull, b.Index())
		}
	}
}

func (packet *InspectPlayerPacket) ReadInspectPlayers(amount int, b buffer.PacketBuffer) {
	packet.InspectPlayerInv = make([]dso_lib.InvItemRec, amount)
	for i := range packet.InspectPlayerInv {
		packet.InspectPlayerInv[i] = dso_lib.InvItemRec{
			Num:   b.ReadUShort(b.Bytes(), b.Index()),
			Value: b.ReadLong(b.Bytes(), b.Index()),
		}
	}
}
