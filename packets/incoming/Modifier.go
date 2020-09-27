package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *ModifierPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ModifierPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ModifierPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ModifierPacket) SetSend(value bool) {
	packet.Send = value
}

type ModifierPacket struct {
	ID   int64
	Send bool

	Mods       []dso_lib.ServerModRec
	WorldBless dso_lib.WorldBlessRec
}

func (packet *ModifierPacket) Read(b buffer.PacketBuffer) {
	packet.Mods = make([]dso_lib.ServerModRec, 12)
	for i := range packet.Mods {
		packet.Mods[i] = dso_lib.ServerModRec{
			Value: b.ReadDouble(b.Bytes(), b.Index()),
			Timer: int64(b.ReadInt(b.Bytes(), b.Index())),
		}
	}
	packet.WorldBless = dso_lib.WorldBlessRec{
		Active: b.ReadBool(b.Bytes(), b.Index()),
	}
	if packet.WorldBless.Active {
		packet.WorldBless.Owner = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Sprite = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.Hair = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.HairTint = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Paperdoll = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Time = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.Value = b.ReadDouble(b.Bytes(), b.Index())
	}
}

func (packet *ModifierPacket) Write(b buffer.PacketBuffer) {
	for _, mod := range packet.Mods {
		b.WriteDouble(b.Bytes(), mod.Value, b.Index())
		b.WriteInt(b.Bytes(), int32(mod.Timer), b.Index())
	}
	b.WriteBool(b.Bytes(), packet.WorldBless.Active, b.Index())
	if packet.WorldBless.Active {
		b.WriteString(b.Bytes(), packet.WorldBless.Owner, b.Index())
		b.WriteInt(b.Bytes(), packet.WorldBless.Sprite, b.Index())
		b.WriteInt(b.Bytes(), packet.WorldBless.Hair, b.Index())
		b.WriteString(b.Bytes(), packet.WorldBless.HairTint, b.Index())
		b.WriteString(b.Bytes(), packet.WorldBless.Paperdoll, b.Index())
		b.WriteInt(b.Bytes(), packet.WorldBless.Time, b.Index())
		b.WriteDouble(b.Bytes(), packet.WorldBless.Value, b.Index())
	}
}
