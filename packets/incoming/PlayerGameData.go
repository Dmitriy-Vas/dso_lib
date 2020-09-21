package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/dso_lib/objects"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *PlayerGameDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerGameDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerGameDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerGameDataPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerGameDataPacket struct {
	ID   int64
	Send bool
}

func (packet *PlayerGameDataPacket) Read(b buffer.PacketBuffer) {
	dso_lib.Global.PlayerCoins = b.ReadInt(b.Bytes(), b.Index())
	for i := range dso_lib.Global.Hotkeys {
		if num := b.ReadInt(b.Bytes(), b.Index()); num > 0 {
			dso_lib.Global.Hotkeys[i] = num
		}
	}
	for i := range dso_lib.Global.Hotbar {
		dso_lib.Global.Hotbar[i] = dso_lib.HotbarRec{
			Slot: int64(b.ReadInt(b.Bytes(), b.Index())),
			Type: b.ReadByte(b.Bytes(), b.Index()),
		}
	}
	for i := range dso_lib.Global.PlayerInventory {
		dso_lib.PlayerInvData(b, i)
	}
	for i := range dso_lib.Global.PlayerCashInventory {
		vector := objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		}
		dso_lib.SetPlayerCashInvItem(i, vector.X, vector.Y)
	}
	for i := range dso_lib.Global.PlayerSpells {
		dso_lib.Global.PlayerSpells[i] = dso_lib.PlayerSpellRec{
			Spell:  b.ReadInt(b.Bytes(), b.Index()),
			Uses:   b.ReadInt(b.Bytes(), b.Index()),
			Master: b.ReadBool(b.Bytes(), b.Index()),
		}
	}
	for i := range dso_lib.Global.PlayerVars {
		dso_lib.Global.PlayerVars[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	dso_lib.Global.TNL = b.ReadLong(b.Bytes(), b.Index())
	if quests := b.ReadInt(b.Bytes(), b.Index()); quests > 0 {
		for i := 1; int32(i) <= quests; i++ {
			dso_lib.Global.PlayerQuests[i] = dso_lib.PlayerQuestRec{
				QuestNum:     b.ReadInt(b.Bytes(), b.Index()),
				Status:       b.ReadByte(b.Bytes(), b.Index()),
				ActualTask:   b.ReadByte(b.Bytes(), b.Index()),
				CurrentCount: int64(b.ReadInt(b.Bytes(), b.Index())),
				TaskCount:    0,
			}
			for x := 1; x <= 10; x++ { // TODO move int to constants
				if dso_lib.Quest[dso_lib.Global.PlayerQuests[i].QuestNum].Task[x].Order > 0 {
					dso_lib.Global.PlayerQuests[i].TaskCount++
				}
			}
		}
	}
	// TODO Update available quests
	b.ReadInt(b.Bytes(), b.Index())
	b.ReadByte(b.Bytes(), b.Index())
	b.ReadInt(b.Bytes(), b.Index())

	if value := b.ReadInt(b.Bytes(), b.Index()); value > 0 {
		for i := range dso_lib.Global.PlayerCraft {
			dso_lib.Global.PlayerCraft[i] = b.ReadInt(b.Bytes(), b.Index())
		}
	}
	for i := range dso_lib.Global.PlayerCards {
		dso_lib.Global.PlayerCards[i].Level = b.ReadInt(b.Bytes(), b.Index())
		dso_lib.Global.PlayerCards[i].Exp = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := 0; i <= 50; i++ {
		b.ReadInt(b.Bytes(), b.Index()) // TODO player emojis
	}
	for i := range dso_lib.Global.PlayerAwards {
		if dso_lib.Global.PlayerAwards[i].Level == 0 {
			dso_lib.Global.PlayerAwards[i].Level++
		}
		if value := b.ReadInt(b.Bytes(), b.Index()); value > 0 {
			dso_lib.Global.PlayerAwards[value].Level = b.ReadInt(b.Bytes(), b.Index())
			dso_lib.Global.PlayerAwards[value].Count = b.ReadInt(b.Bytes(), b.Index())
			dso_lib.Global.PlayerAwards[value].GetDate = b.ReadString(b.Bytes(), b.Index(), 0)
		} else {
			// Garbage
			// b.ReadInt(b.Bytes(), b.Index())
		}
	}
	for i := range dso_lib.Global.PlayerCalaveras {
		dso_lib.Global.PlayerCalaveras[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range dso_lib.Global.PlayerPin {
		dso_lib.Global.PlayerPin[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	for i := range dso_lib.Global.PlayerProfession {
		b.ReadByte(b.Bytes(), b.Index())
		b.ReadInt(b.Bytes(), b.Index())
		b.ReadByte(b.Bytes(), b.Index())
		b.ReadLong(b.Bytes(), b.Index())
		for x := 0; x <= 9; x++ { // TODO move int to constants
			dso_lib.Global.PlayerProfession[i].Upgrade[x] = b.ReadByte(b.Bytes(), b.Index())
		}
	}
	b.ReadString(b.Bytes(), b.Index(), 0)
	b.ReadByte(b.Bytes(), b.Index())
	b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerGameDataPacket) Write(_ buffer.PacketBuffer) {

}
