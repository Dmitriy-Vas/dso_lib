package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *ProjectilCrossPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ProjectilCrossPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ProjectilCrossPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ProjectilCrossPacket) SetSend(value bool) {
	packet.Send = value
}

type ProjectilCrossPacket struct {
	ID              int64
	Send            bool
	PlayerNum       int64
	ProjectilNum    int64
	ProjectileIndex []int64
	ProjectileCache []dso_lib.ProjectilCacheRec
}

func (packet *ProjectilCrossPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadLong(b.Bytes(), b.Index())
	packet.ProjectilNum = b.ReadLong(b.Bytes(), b.Index())
	packet.ProjectileIndex = make([]int64, 4)
	packet.ProjectileCache = make([]dso_lib.ProjectilCacheRec, 4)
	for i := range packet.ProjectileIndex {
		packet.ProjectileIndex[i] = b.ReadLong(b.Bytes(), b.Index())
		packet.ProjectileCache[i] = dso_lib.ProjectilCacheRec{
			Combo:     byte(b.ReadLong(b.Bytes(), b.Index())),
			Direction: byte(b.ReadLong(b.Bytes(), b.Index())),
		}
	}
}

func (packet *ProjectilCrossPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteLong(b.Bytes(), packet.ProjectilNum, b.Index())
	for i := range packet.ProjectileIndex {
		b.WriteLong(b.Bytes(), packet.ProjectileIndex[i], b.Index())
		b.WriteLong(b.Bytes(), int64(packet.ProjectileCache[i].Combo), b.Index())
		b.WriteLong(b.Bytes(), int64(packet.ProjectileCache[i].Direction), b.Index())
	}
}
