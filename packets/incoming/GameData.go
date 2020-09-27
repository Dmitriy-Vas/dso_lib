package incoming

import (
	"github.com/Dmitriy-Vas/dso_lib"
	"github.com/Dmitriy-Vas/dso_lib/wrapper"
	"github.com/Dmitriy-Vas/wave_buffer"
	"time"
)

// GetID returns packet ID.
func (packet *GameDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *GameDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *GameDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *GameDataPacket) SetSend(value bool) {
	packet.Send = value
}

type GameDataPacket struct {
	ID   int64
	Send bool

	Demo              bool
	ItemsNum          []int32
	Items             []dso_lib.ItemRec
	RandomItemsNum    []int32
	RandomItems       []dso_lib.RandomItemRec
	PuzzlesNum        []int32
	Puzzles           []dso_lib.PuzzleRec
	AnimationsNum     []int32
	Animations        []dso_lib.AnimationRec
	NpcsNum           []int32
	Npcs              []dso_lib.NpcRec
	ShopsNum          []int32
	Shops             []dso_lib.ShopRec
	SpellsNum         []int32
	Spells            []dso_lib.SpellRec
	ResourcesNum      []int32
	Resources         []dso_lib.ResourceRec
	MoralsNum         []int32
	Morals            []dso_lib.MoralRec
	ProjectilsNum     []int32
	Projectils        []dso_lib.ProjectilRec
	CardsNum          []int32
	Cards             []dso_lib.CardsRec
	CraftsNum         []int32
	Crafts            []dso_lib.CraftingRec
	QuestsNum         []int32
	Quests            []dso_lib.QuestRec
	CashShopsNum      []int32
	CashShops         []dso_lib.CashShopRec
	TopItems          []int32
	HotItems          []dso_lib.HotItemRec
	Variable1         byte
	ConditionsNum     []int32
	Conditions        []dso_lib.ConditionRec
	EmoticonsNum      []int32
	Emoticons         []dso_lib.EmoticonRec
	DailyCheck        []dso_lib.InvItemRec
	MapList           []dso_lib.MapListRec
	Pins              []dso_lib.PinsRec
	Professions       []dso_lib.ProfessionRec
	Newspaper         dso_lib.NewspaperRec
	EventTime         int32
	WorldBoss         []time.Time
	Promotion         dso_lib.PromotionRec
	ServerMods        []dso_lib.ServerModRec
	PlayerBlockedList string
	ServerGlobalVar   int32
	DailyCheckDate    time.Time
	WorldBless        dso_lib.WorldBlessRec
	Speed             dso_lib.SpeedFormulaRec
	Variable2         byte
	Selected          byte
	RestrictedWords   []string

	Data []byte
}

func (packet *GameDataPacket) Read(b buffer.PacketBuffer) {
	packet.Data = b.(*wrapper.Buffer).DefaultBuffer.ReadBytes(b.Bytes(), b.Index(), b.Len()-16)
}

func (packet *GameDataPacket) Write(b buffer.PacketBuffer) {
	b.(*wrapper.Buffer).DefaultBuffer.WriteBytes(b.Bytes(), packet.Data, b.Index())
}

func (packet *GameDataPacket) RealRead(b buffer.PacketBuffer) {
	// TODO decompress
	packet.Demo = b.ReadBool(b.Bytes(), b.Index())
	itemsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ItemsNum = make([]int32, itemsAmount)
	packet.Items = make([]dso_lib.ItemRec, itemsAmount)
	for i := range packet.Items {
		packet.ItemsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Items[i] = dso_lib.ReadItemData(b)
	}
	randomItemsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.RandomItemsNum = make([]int32, randomItemsAmount)
	packet.RandomItems = make([]dso_lib.RandomItemRec, randomItemsAmount)
	for i := range packet.RandomItems {
		packet.RandomItemsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.RandomItems[i] = dso_lib.ReadRandomItemData(b)
	}
	puzzlesAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.PuzzlesNum = make([]int32, puzzlesAmount)
	packet.Puzzles = make([]dso_lib.PuzzleRec, puzzlesAmount)
	for i := range packet.Puzzles {
		packet.PuzzlesNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Puzzles[i] = dso_lib.ReadPuzzleData(b)
	}
	animationAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.AnimationsNum = make([]int32, animationAmount)
	packet.Animations = make([]dso_lib.AnimationRec, animationAmount)
	for i := range packet.Animations {
		packet.AnimationsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Animations[i] = dso_lib.ReadAnimationData(b)
	}
	npcsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.NpcsNum = make([]int32, npcsAmount)
	packet.Npcs = make([]dso_lib.NpcRec, npcsAmount)
	for i := range packet.Npcs {
		packet.NpcsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Npcs[i] = dso_lib.ReadNpcData(b)
	}
	shopsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ShopsNum = make([]int32, shopsAmount)
	packet.Shops = make([]dso_lib.ShopRec, shopsAmount)
	for i := range packet.Shops {
		packet.ShopsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Shops[i] = dso_lib.ReadShopData(b)
	}
	spellsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.SpellsNum = make([]int32, spellsAmount)
	packet.Spells = make([]dso_lib.SpellRec, spellsAmount)
	for i := range packet.Spells {
		packet.SpellsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Spells[i] = dso_lib.ReadSpellData(b)
	}
	resourcesAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ResourcesNum = make([]int32, resourcesAmount)
	packet.Resources = make([]dso_lib.ResourceRec, resourcesAmount)
	for i := range packet.Resources {
		packet.ResourcesNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Resources[i] = dso_lib.ReadResourceData(b)
	}
	moralsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.MoralsNum = make([]int32, moralsAmount)
	packet.Morals = make([]dso_lib.MoralRec, moralsAmount)
	for i := range packet.Morals {
		packet.MoralsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Morals[i] = dso_lib.ReadMoralData(b)
	}
	projectilsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ProjectilsNum = make([]int32, projectilsAmount)
	packet.Projectils = make([]dso_lib.ProjectilRec, projectilsAmount)
	for i := range packet.Projectils {
		packet.ProjectilsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		projectil := dso_lib.ProjectilRec{
			Name:      b.ReadString(b.Bytes(), b.Index(), 0),
			Damage:    b.ReadInt(b.Bytes(), b.Index()),
			Pic:       b.ReadUShort(b.Bytes(), b.Index()),
			Range:     b.ReadByte(b.Bytes(), b.Index()),
			Speed:     b.ReadUShort(b.Bytes(), b.Index()),
			Animation: b.ReadUShort(b.Bytes(), b.Index()),
			Light:     b.ReadBool(b.Bytes(), b.Index()),
			Int:       make([]int32, 3), // TODO int to const
		}
		for x := range projectil.Int {
			projectil.Int[x] = b.ReadInt(b.Bytes(), b.Index())
		}
	}
	cardsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.CardsNum = make([]int32, cardsAmount)
	packet.Cards = make([]dso_lib.CardsRec, cardsAmount)
	for i := range packet.Cards {
		packet.CardsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Cards[i] = dso_lib.ReadCardData(b)
	}
	craftsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.CraftsNum = make([]int32, craftsAmount)
	packet.Crafts = make([]dso_lib.CraftingRec, craftsAmount)
	for i := range packet.Cards {
		packet.CraftsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Crafts[i] = dso_lib.ReadCraftData(b)
	}
	questsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.QuestsNum = make([]int32, questsAmount)
	packet.Quests = make([]dso_lib.QuestRec, questsAmount)
	for i := range packet.Quests {
		packet.QuestsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Quests[i] = dso_lib.ReadQuestData(b)
	}
	cashShopsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.CashShopsNum = make([]int32, cashShopsAmount)
	packet.CashShops = make([]dso_lib.CashShopRec, cashShopsAmount)
	for i := range packet.CashShops {
		packet.CashShopsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.CashShops[i] = dso_lib.ReadCashShopData(b)
	}
	packet.TopItems = make([]int32, 8) // TODO int to const
	for i := range packet.TopItems {
		packet.TopItems[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	packet.HotItems = make([]dso_lib.HotItemRec, 24) // TODO int to const
	for i := range packet.HotItems {
		packet.HotItems[i] = dso_lib.HotItemRec{
			Category:    b.ReadInt(b.Bytes(), b.Index()),
			SubCategory: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	if packet.Variable1 == 4 {
		conditionsAmount := b.ReadInt(b.Bytes(), b.Index())
		packet.ConditionsNum = make([]int32, conditionsAmount)
		packet.Conditions = make([]dso_lib.ConditionRec, conditionsAmount)
		for i := range packet.Conditions {
			packet.ConditionsNum[i] = b.ReadInt(b.Bytes(), b.Index())
			packet.Conditions[i] = dso_lib.ReadConditionData(b)
		}
	}
	emoticonsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.EmoticonsNum = make([]int32, emoticonsAmount)
	packet.Emoticons = make([]dso_lib.EmoticonRec, emoticonsAmount)
	for i := range packet.Emoticons {
		packet.EmoticonsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Emoticons[i] = dso_lib.ReadEmoticonData(b)
	}
	packet.DailyCheck = make([]dso_lib.InvItemRec, 28) // TODO int to const
	for i := range packet.DailyCheck {
		packet.DailyCheck[i] = dso_lib.InvItemRec{
			Num:   b.ReadUShort(b.Bytes(), b.Index()),
			Value: b.ReadLong(b.Bytes(), b.Index()),
		}
	}
	packet.MapList = make([]dso_lib.MapListRec, 200) // TODO int to const
	for i := range packet.MapList {
		packet.MapList[i].Name = make([]string, 2) // TODO int to const
		for x := range packet.MapList[i].Name {
			packet.MapList[i].Name[x] = b.ReadString(b.Bytes(), b.Index(), 0)
		}
	}
	packet.Pins = make([]dso_lib.PinsRec, 51)
	for i := range packet.Pins {
		packet.Pins[i] = dso_lib.PinsRec{
			Item: b.ReadInt(b.Bytes(), b.Index()),
			Desc: b.ReadString(b.Bytes(), b.Index(), 0),
		}
	}
	packet.Professions = dso_lib.ReadProfessionData(b)
	packet.Newspaper = dso_lib.ReadNewspaperData(b)
	packet.EventTime = b.ReadInt(b.Bytes(), b.Index())
	packet.WorldBoss = make([]time.Time, 2) // TODO int to const
	for i := range packet.WorldBoss {
		packet.WorldBoss[i] = wrapper.ReadDate(b)
	}
	packet.Promotion = dso_lib.PromotionRec{
		Type: b.ReadByte(b.Bytes(), b.Index()),
		Item: b.ReadInt(b.Bytes(), b.Index()),
		GameCard: make([]dso_lib.PromotionGameCardRec,
			b.ReadByte(b.Bytes(), b.Index())),
	}
	for i := range packet.Promotion.GameCard {
		card := dso_lib.PromotionGameCardRec{
			Name:  b.ReadString(b.Bytes(), b.Index(), 0),
			Price: b.ReadString(b.Bytes(), b.Index(), 0),
			Bonus: b.ReadInt(b.Bytes(), b.Index()),
		}
		packet.Promotion.GameCard[i] = card
	}
	packet.ServerMods = make([]dso_lib.ServerModRec, dso_lib.ServerModQuestExp+1)
	for i := dso_lib.ServerModExperience; i <= dso_lib.ServerModQuestExp; i++ {
		packet.ServerMods[i] = dso_lib.ServerModRec{
			Value: b.ReadDouble(b.Bytes(), b.Index()),
			Timer: int64(b.ReadInt(b.Bytes(), b.Index())),
		}
	}
	packet.PlayerBlockedList = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.WorldBless = dso_lib.WorldBlessRec{Active: b.ReadBool(b.Bytes(), b.Index())}
	if packet.WorldBless.Active {
		packet.WorldBless.Owner = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Sprite = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.Hair = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.HairTint = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Paperdoll = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Time = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.Value = b.ReadDouble(b.Bytes(), b.Index())
	}
	packet.Speed = dso_lib.SpeedFormulaRec{
		Normal: b.ReadDouble(b.Bytes(), b.Index()),
		Buff:   b.ReadDouble(b.Bytes(), b.Index()),
		Mount:  b.ReadDouble(b.Bytes(), b.Index()),
	}
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
	packet.Selected = b.ReadByte(b.Bytes(), b.Index())
	packet.RestrictedWords = make([]string, 101) // TODO int to const
	for i := range packet.RestrictedWords {
		packet.RestrictedWords[i] = b.ReadString(b.Bytes(), b.Index(), 0)
		if packet.RestrictedWords[i] == "" {
			break
		}
	}
}

func (packet *GameDataPacket) RealWrite(b buffer.PacketBuffer) {
	// TODO compress
}
