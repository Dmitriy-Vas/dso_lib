package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *UpdateCraftingPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateCraftingPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateCraftingPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateCraftingPacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateCraftingPacket struct {
	ID    int64
	Send  bool
	Num   int64
	Craft dso_lib.CraftingRec
}

func (packet *UpdateCraftingPacket) Read(b buffer.PacketBuffer) {
	if packet.Num = b.ReadLong(b.Bytes(), b.Index()); packet.Num >= 0 && packet.Num <= 255 {
		craft := dso_lib.CraftingRec{
			Name:    b.ReadString(b.Bytes(), b.Index(), 0),
			Desc:    b.ReadString(b.Bytes(), b.Index(), 0),
			ItemReq: make([]int32, 3),
			ItemVal: make([]int32, 3),
		}
		for i := range craft.ItemReq {
			craft.ItemReq[i] = b.ReadInt(b.Bytes(), b.Index())
			craft.ItemVal[i] = b.ReadInt(b.Bytes(), b.Index())
		}
		craft.Reward = b.ReadInt(b.Bytes(), b.Index())
		craft.RewardVal = b.ReadInt(b.Bytes(), b.Index())
		packet.Craft = craft
	}
}

func (packet *UpdateCraftingPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Num, b.Index())
	if packet.Num >= 0 && packet.Num <= 255 {
		b.WriteString(b.Bytes(), packet.Craft.Name, b.Index())
		b.WriteString(b.Bytes(), packet.Craft.Desc, b.Index())
		for i := range packet.Craft.ItemReq {
			b.WriteInt(b.Bytes(), packet.Craft.ItemReq[i], b.Index())
			b.WriteInt(b.Bytes(), packet.Craft.ItemVal[i], b.Index())
		}
		b.WriteInt(b.Bytes(), packet.Craft.Reward, b.Index())
		b.WriteInt(b.Bytes(), packet.Craft.RewardVal, b.Index())
	}
}
