package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/dso_lib/objects"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *SpawnMobItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SpawnMobItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SpawnMobItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SpawnMobItemPacket) SetSend(value bool) {
	packet.Send = value
}

type SpawnMobItemPacket struct {
	ID        int64
	Send      bool
	ItemIndex int32
	MapItem   dso_lib.MapItemRec
}

func (packet *SpawnMobItemPacket) Read(b buffer.PacketBuffer) {
	if packet.ItemIndex = b.ReadInt(b.Bytes(), b.Index()); packet.ItemIndex != 0 && packet.ItemIndex <= 255 {
		item := dso_lib.MapItemRec{
			Item: dso_lib.InvItemRec{
				Num: b.ReadUShort(b.Bytes(), b.Index()),
			},
		}
		if item.Item.Num > 0 {
			item.Item.Value = int64(b.ReadInt(b.Bytes(), b.Index()))
			item.PIndex = b.ReadInt(b.Bytes(), b.Index())
			item.Pos = objects.Vector2{
				X: b.ReadInt(b.Bytes(), b.Index()),
				Y: b.ReadInt(b.Bytes(), b.Index()),
			}
			item.X2 = int64(b.ReadInt(b.Bytes(), b.Index()))
			item.Y2 = int64(b.ReadInt(b.Bytes(), b.Index()))
		}
		packet.MapItem = item
	}
}

func (packet *SpawnMobItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ItemIndex, b.Index())
	if packet.ItemIndex != 0 && packet.ItemIndex <= 255 {
		b.WriteUShort(b.Bytes(), packet.MapItem.Item.Num, b.Index())
		if packet.MapItem.Item.Num > 0 {
			b.WriteInt(b.Bytes(), int32(packet.MapItem.Item.Value), b.Index())
			b.WriteInt(b.Bytes(), packet.MapItem.PIndex, b.Index())
			b.WriteInt(b.Bytes(), packet.MapItem.Pos.X, b.Index())
			b.WriteInt(b.Bytes(), packet.MapItem.Pos.Y, b.Index())
			b.WriteInt(b.Bytes(), int32(packet.MapItem.X2), b.Index())
			b.WriteInt(b.Bytes(), int32(packet.MapItem.Y2), b.Index())
		}
	}
}
