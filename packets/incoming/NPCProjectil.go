package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *NPCProjectilPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *NPCProjectilPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *NPCProjectilPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *NPCProjectilPacket) SetSend(value bool) {
	packet.Send = value
}

type NPCProjectilPacket struct {
	ID             int64
	Send           bool
	NpcNum         int32
	Type           byte
	ProjectileNum  []int32
	NPCProjectiles []dso_lib.NPCProjectilRec
	Direction      byte
}

func (packet *NPCProjectilPacket) Read(b buffer.PacketBuffer) {
	packet.NpcNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	switch packet.Type {
	case 0:
		packet.NPCProjectiles = make([]dso_lib.NPCProjectilRec, 1)
		packet.ProjectileNum = make([]int32, 1)
	case 1:
		packet.NPCProjectiles = make([]dso_lib.NPCProjectilRec, 4)
		packet.ProjectileNum = make([]int32, 4)
	}
	for i := range packet.NPCProjectiles {
		packet.ProjectileNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.NPCProjectiles[i] = dso_lib.NPCProjectilRec{
			Num:       int64(b.ReadInt(b.Bytes(), b.Index())),
			Direction: int64(b.ReadByte(b.Bytes(), b.Index())),
		}
	}
	packet.Direction = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *NPCProjectilPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	for i := range packet.NPCProjectiles {
		b.WriteInt(b.Bytes(), packet.ProjectileNum[i], b.Index())
		b.WriteInt(b.Bytes(), int32(packet.NPCProjectiles[i].Num), b.Index())
		b.WriteByte(b.Bytes(), byte(packet.NPCProjectiles[i].Direction), b.Index())
	}
	b.WriteByte(b.Bytes(), packet.Direction, b.Index())
}
