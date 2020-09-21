package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/wave_buffer"
)

// GetID returns packet ID.
func (packet *CashShopDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CashShopDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CashShopDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CashShopDataPacket) SetSend(value bool) {
	packet.Send = value
}

type CashShopDataPacket struct {
	ID      int64
	Send    bool
	ShopNum byte
	Item    dso_lib.CashShopRec
}

func (packet *CashShopDataPacket) Read(b buffer.PacketBuffer) {
	packet.ShopNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Item = dso_lib.ReadCashShopData(b)
}

func (packet *CashShopDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ShopNum, b.Index())
	dso_lib.WriteCashShopData(b, packet.Item)
}
