package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib/wrapper"
	"github.com/Dmitriy-Vas/wave_buffer"
	"time"
)

// GetID returns packet ID.
func (packet *SendTimePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SendTimePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SendTimePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SendTimePacket) SetSend(value bool) {
	packet.Send = value
}

type SendTimePacket struct {
	ID             int64
	Send           bool
	DayTime        bool
	ServerTimezone time.Time
}

func (packet *SendTimePacket) Read(b buffer.PacketBuffer) {
	packet.DayTime = b.ReadBool(b.Bytes(), b.Index())
	packet.ServerTimezone = wrapper.ReadDate(b)
}

func (packet *SendTimePacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.DayTime, b.Index())
	wrapper.WriteDate(b, packet.ServerTimezone)
}
