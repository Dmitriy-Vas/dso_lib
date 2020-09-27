package incoming

import (
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerQuestPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerQuestPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerQuestPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerQuestPacket) SetSend(value bool) {
	packet.Send = value
}

// TODO finish packet
type PlayerQuestPacket struct {
	ID             int64
	Send           bool
	RegionNum      int32
	Num            int32
	AmountOfQuests int32
	Variable4      bool
	QuestNum       int32
	Status         byte
	ActualTask     byte
	CurrentCount   int32
	PromotionURL   string
	Variable10     int32
	Variable11     byte
	Variable12     byte
	Variable13     int32
}

func (packet *PlayerQuestPacket) Read(b buffer.PacketBuffer) {
	packet.RegionNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.AmountOfQuests = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadBool(b.Bytes(), b.Index())
	packet.QuestNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Status = b.ReadByte(b.Bytes(), b.Index())
	packet.ActualTask = b.ReadByte(b.Bytes(), b.Index())
	packet.CurrentCount = b.ReadInt(b.Bytes(), b.Index())
	packet.PromotionURL = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable10 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable11 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable12 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable13 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerQuestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.RegionNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteInt(b.Bytes(), packet.AmountOfQuests, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.QuestNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Status, b.Index())
	b.WriteByte(b.Bytes(), packet.ActualTask, b.Index())
	b.WriteInt(b.Bytes(), packet.CurrentCount, b.Index())
	b.WriteString(b.Bytes(), packet.PromotionURL, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable10, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable11, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable12, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable13, b.Index())
}
