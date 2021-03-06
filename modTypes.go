package dso_lib

import (
	"github.com/Dmitriy-Vas/dso_lib/objects"
	"time"
)

type EmoticonRec struct {
	Name      []string
	Command   string
	Image     int32
	IsDefault bool
	Animated  bool
	Unlocked  bool
}

type MapListRec struct {
	Name []string
}

type PlayerBuffsRec struct {
	Type     byte
	SpellNum int32
	Timer    int32
	Vital    int32
}

type PlayerBuddiesRec struct {
	Name       string
	Sex        byte
	Classes    byte
	Sprite     uint16
	Hair       uint16
	FriendDate string
	Paperdoll  string
	HairTint   string
	Status     bool
}

type WorldRec struct {
	Name    string
	Sprite  int
	sizeX   int
	sizeY   int
	offsetX int
	offsetY int
	Files   byte
	Column  byte
	Map     []WorldMapRec
}

type WorldMapRec struct {
	Num    int
	Status byte
}

type FontDataRec struct {
	Name       string
	Font       []byte
	Sprite     int
	Separation int
	Height     int
}

type ServerOptionRec struct {
	Mode      []ServerModRec
	EventItem int
	Demo      bool
	Clone     bool
	Promotion PromotionRec
}

type PromotionRec struct {
	Type     byte
	Item     int32
	GameCard []PromotionGameCardRec
}

type PromotionGameCardRec struct {
	Name  string
	Price string
	Bonus int32
}

type ServerModRec struct {
	Value float64
	Timer int64
}

type ServerModeRec struct {
	Experience float64
	Drill      float64
	Drop       float64
	Damage     float64
	Heal       float64
	Junk       float64
	LostExp    float64
	Party      float64
	Profession float64
}

type PlayerCardRec struct {
	Level int32
	Exp   int32
}

type AwardsRec struct {
	Name       string
	ESName     string
	Desc       string
	Type       byte
	Level      []AwardsLevelRec
	LevelCount int32
	Blocked    bool
}

type AwardsLevelRec struct {
	Required  int32
	Icon      int32
	AP        int32
	Calaveras int32
}

type PlayerProfessionRec struct {
	Level   byte
	Upgrade []byte
	Points  byte
	EXP     int32
	Times   int64
}

type ProfessionRec struct {
	Icon    int32
	Upgrade []ProfessionUpgradeRec
}

type ProfessionUpgradeRec struct {
	Name     []string
	Desc     []string
	MaxLevel byte
	Icon     int32
}

type PlayerAwardsRec struct {
	Level   int32
	Count   int32
	GetDate string
}

type RegionRec struct {
	Name     []string
	Music    RegionMusicRec
	Fog      RegionFogRec
	Sunlight bool
	Night    RegionNightRec
}

type RegionNightRec struct {
	Alpha int
}

type RegionFogRec struct {
	Sprite int
	Alpha  byte
	Speed  byte
}

type RegionMusicRec struct {
	Name       string
	LoopState  bool
	StartPoint uint
	EndPoint   uint
}

type EventsRec struct {
	Name    []string
	Players []EventsRankRec
}

type EventsRankRec struct {
	Name    string
	Classes int32
	Score   int64
}

type OptionsInterfaceRec struct {
	Pos []objects.Vector2
}

type OptionsQuestLogRec struct {
	Player []OptionsPlayerQuestLogRec
}

type OptionsPlayerQuestLogRec struct {
	Name      string
	QuestList []int
}

type PlayerElementRec struct {
	Num  int32
	Icon int32
}

type PlayerComboRec struct {
	Num int32
}

type ComboCacheDataRec struct {
	Num     int32
	Anim    int32
	Require int32
}

type MyMusicRec struct {
	Name   string
	Format string
	File   []byte
}

type CreateCharRec struct {
	ClassSelect    int
	SexSelect      byte
	HairSelect     int
	SpriteSelect   int
	HairTintSelect int
	Sprite         []createCharSpriteRec
	Hair           []createCharSpriteRec
	HairTint       []objects.Color
}

type createCharSpriteRec struct {
	Male   int
	Female int
	MFace  int
	FFace  int
}

type PlayPuzzleRec struct {
	Data   PuzzleRec
	Status int
}

type PuzzleRec struct {
	Name        string
	Map         int32
	Point       objects.Vector2
	Size        objects.Vector2
	SwitchImage int32
	SwitchPos   objects.Vector2
	Cube        [][]PuzzleCubeRec
}

type PuzzleCubeRec struct {
	Image        int32
	Block        bool
	Move         bool
	Key          bool
	Switch       int32
	isPlayerMove int32
}

type AdminInfoRec struct {
	Map          []string
	Level        []int64
	PlayersOfMap []int
	Puzzle       []string
	PartyQuest   []string
	Elements     []string
	Update       bool
}

type ControlRec struct {
	ID int
}

type TextureBox struct {
	Sprite int
	Width  int
	Height int
	pos    objects.Vector2
	Color  objects.Color
}

type ChoiceRec struct {
	MyChoice      int32
	ChoiceText    int32
	ChoiceSelText int32
	GenDialog     int32
	Sprite        int32
	DialogType    byte
	ChoiceDialog  string
	Max           int32
	ChoiceArray   []string
	Code          bool
	isCondition   bool
	ConditionLine objects.Vector2
}

type DrillModRec struct {
	DrillDepth int32
	Active     bool
	Pause      bool
	SoundIndex int32
	Dir        byte
	DirStar    byte
}

type ConnectionRec struct {
	IP   string
	Port int
}

type OptionsRec struct {
	Username         string
	Music            bool
	Sound            bool
	VSync            bool
	FullScreen       bool
	Debug            bool
	CheckLang        bool
	WindowsCursor    bool
	MapClean         bool
	SaveID           bool
	Zoom             bool
	Lang             string
	Resolution       byte
	Launcher         byte
	MVolume          byte
	SVolume          byte
	ScreenShotFormat string
	Game             GameOptionsRec
	UI               UIRec
	MyJoystick       JoystickRec
}

type GameOptionsRec struct {
	PlayerName bool
	BadWords   bool
}

type JoystickRec struct {
	Active bool
	Analog bool
	ID     byte
}

type UIRec struct {
	BigEmoticon   bool
	Cooldown      bool
	ExpBarPercent bool
}

type PartyRec struct {
	Leader      int32
	Member      []MemberPartyRec
	MemberCount byte
	Type        byte
	Level       int32
	Experience  int64
	PartyQuest  InPartyQuestRec
	ShareItems  bool
}

type InPartyQuestRec struct {
	Time int
}

type MemberPartyRec struct {
	Index    int32
	Level    int32
	Hair     int32
	Sprite   int32
	Classes  int32
	Map      int32
	Name     string
	HairTint string
	Mask     int32
	Helmet   int32
}

type InvItemRec struct {
	Num   uint16
	Value int64
	Slot  byte
	Stat  []byte
}

type BankRec struct {
	Slot byte
	Item []InvItemRec
}

type PlayerRec struct {
	Name             string
	Classes          byte
	Sprite           int
	Sex              byte
	Access           byte
	Level            int
	EXP              int64
	PK               byte
	Honor            int
	HonorDate        time.Time
	Vital            []int64
	Stat             []int64
	POINTS           int
	Equipment        []InvItemRec
	CashEquipment    []InvItemRec
	Pin              int
	Map              int
	Pos              objects.Vector2
	Dir              byte
	MaxVital         []int64
	WalkingTimer     int64
	Offset           objects.Vector2
	ScreenPos        objects.Vector2
	Moving           bool
	Attacking        byte
	AttackTimer      int64
	MapGetTimer      int64
	Steps            byte
	Hair             int64
	HairTint         objects.Color
	Defensa          byte
	DefensaBool      bool
	shieldHP         int64
	CancelDefensa    byte
	SpellBuffer      PlayerSpellBufferRec
	ProjecTil        []ProjectilRec
	StartFlash       int64
	Flash            byte
	Hide             int64
	CurCard          int64
	AmountOfCrafting int64
	Ghost            byte
	AmountOfTitle    int64
	Title            []int
	curTitle         int64
	Language         int64
	AmountOfQuests   int
	DrillMod         DrillModRec
	Alpha            int64
	Temp             PlayerTempRec
	Reflection       ReflectionRec
	Interaction      int64
	PuzzleOffset     objects.Vector2
	pComboCount      int
	ElementSelect    []int
	Light            int
	AP               int
	moveSpeed        float64
	elapsedTime      int64
	AutoTarget       objects.Vector2
	IsFishing        int
	Polymorph        int
	DamageSkin       byte
	ProjectilSkin    byte
	BuffsEffect      PlayerBuffEffectRec
	serverDelay      int
	Trap             PlayerTrapRec
	Emoticon         PlayerEmoticonRec
	Trapped          TrappedRec
}

type PlayerEmoticonRec struct {
	Num   int
	Timer int64
	isPin bool
}

type PlayerTrapRec struct {
	Spell    int
	X        int
	Y        int
	Duration int64
}

type PlayerBuffEffectRec struct {
	Poisoned bool
	Burned   bool
}

type PlayerTempRec struct {
	Invisible    bool
	PickUpItem   objects.Vector2
	PickUpTimer  int64
	PickUpPos    objects.Vector2
	MasterSprite int
	MasterTimer  int64
	MasterPos    objects.Vector2
	Mount        int
}

type PlayerSpellBufferRec struct {
	SpellAnimNum int
	SpellAnim    int
	SpellTimer   int64
	SpellBoom    bool
}

type ReflectionRec struct {
	pos objects.Vector2
	// rec  Rectangle
	Img  int
	Type int
}

type TileRec struct {
	Type      byte
	Data1     int
	Data2     int
	Data3     int
	Data4     string
	DirBlock  byte
	Topo      int
	DataTopo1 int
	DataTopo2 int
}

type MapRec struct {
	ID              int32
	Name            string
	ESName          string
	Region          int32
	Image           int32
	Revision        int32
	Moral           int32
	Up              int32
	Down            int32
	Left            int32
	Right           int32
	BootMap         int32
	BootX           int32
	BootY           int32
	MaxX            byte
	MaxY            byte
	Tile            []TileRec
	tileOffset      objects.Vector2
	Npc             []int32
	FogOpacity      byte
	Fog             int32
	FogSpeed        int32
	tileA           int32
	tileB           int32
	DayNight        int32
	Parallax        int32
	ParallaxType    byte
	Ambience        string
	Lighting        bool
	NightAlpha      int32
	InstanceMax     int32
	MapCondition    int32
	Puzzle          objects.Vector4
	Drill           objects.Vector4
	Temp            MapTempRec
	BackgroundColor string
}

type MapTempRec struct {
	EventCount int
	Events     []MapEventRec
}

type MapEventRec struct {
	Sprite  int
	Mode    byte
	Pos     objects.Vector2
	myFrame objects.Vector2
	Name    string
	OnJoin  bool
	Mask    int
	Width   int
	Height  int
	Offset  objects.Vector2
}

type SpriteData struct {
	Flash bool
	Hide  bool
}

type ClassRec struct {
	Name     string
	Stat     []byte
	Sprite   []ClassSpriteRec
	Hair     []ClassSpriteRec
	HairTint []ClassHairTintRec
	Vital    []int64
}

type ClassHairTintRec struct {
	Name    string
	Color   string // objects.Color
	Premium bool
}

type ClassSpriteRec struct {
	Name    string
	Male    int32
	Female  int32
	MFace   int32
	FFace   int32
	Premium bool
}

type ItemRec struct {
	Name          string
	ESName        string
	Pic           int32
	Revision      int32
	Type          byte
	Data1         int32
	Data2         int32
	Tool          int32
	ClassReq      []int64
	AccessReq     int64
	LevelReq      int64
	Mastery       byte
	Price         int64
	AddStat       []int32
	Rarity        byte
	Speed         int64
	Handed        int64
	BindType      byte
	StatReq       []int32
	Animation     int64
	Paperdoll     int64
	Desc          string
	Stackable     bool
	Overol        int32
	paperWidth    int32
	ProjecTil     int32
	MType         int32
	HPBonus       int32
	MPBonus       int32
	WarpMap       int32
	WarpX         int32
	WarpY         int32
	isHair        int32
	Element       []int32
	ElementChance []int32
	BigPic        int32
	VitalMode     []int32
	comboSlot     int64
	Recipe        int32
	Int           []int32
	Bool          []bool
}
type QuestRec struct {
	Name                  string
	ESName                string
	Repeat                bool
	RequiredCompleteQuest bool
	Status                bool
	QuestLog              string
	Story                 []string
	RequiredLevel         int32
	RequiredQuest         int32
	RequiredClass         []int32
	Item                  []QuestRewardItemRec
	Task                  []TaskRec
	QuestHelp             int32
	Owner                 int32
	Type                  int32
	Promotion             byte
	QuestAtEnd            int32
	Bool                  []bool
}
type QuestRewardItemRec struct {
	Item    int32
	Task    int32
	Type    int32
	Classes int32
	Value   int64
}

type TaskRec struct {
	Order          int32
	NPC            int32
	PartyQuest     int32
	Item           int32
	Map            int32
	Resource       int32
	Recipe         int32
	SpellUse       int32
	Tutorial       int32
	Amount         int64
	Speech         string
	TaskLog        string
	QuestEnd       bool
	RewardEXP      int64
	SpellReward    []int32
	ElementReward  int32
	RewardRecipe   int32
	RewardVariable int32
	RewardCalavera int64
	Image          int32
	ImageVec       objects.Vector2
	AutoComplete   bool
	StartQuote     []string
	EndQuote       []string
	Sprite         []int32
	Int            []int32
}

type MapItemRec struct {
	Item   InvItemRec
	Slot   byte
	Stat   []byte
	Frame  byte
	Pos    objects.Vector2
	X2     int64
	Y2     int64
	XTmr   int64
	Timer  int64
	Up     bool
	YTmr   int64
	Jump   byte
	PIndex int32
	PTimer int64
	Angle  int
}

type NpcRec struct {
	Name          string
	AttackSay     string
	Behaviour     byte
	Range         byte
	DropChance    []int32
	DropItem      []int32
	DropItemValue []int32
	Int           []int32
	Vec           []objects.Vector2
	Faction       byte
	CardDrop      int32
	CardNum       int32
	Spell         []int32
	Element       []int64
	Bool          []bool
	RequireVar    int32
	AttackSound   string
	Light         int32
	QuestList     []int32
	ScriptList    []NpcScriptListRec
	Offset        objects.Vector2
}

type NpcScriptListRec struct {
	Num      int32
	Interval int32
	Variable string
}

type MapNpcRec struct {
	Num         int
	Target      int
	targetType  byte
	Vital       []int64
	Map         int
	Pos         objects.Vector2
	Dir         byte
	HPSetTo     int64
	Offset      MapNpcOffsetRec
	XOffset     int
	YOffset     int
	Moving      int
	Attacking   int
	AttackTimer int64
	Steps       int64
	AttackNum   int64
	AnimTimer   int
	ProjecTil   []NPCProjectilRec
	AutoTarget  objects.Vector2
	Tile        []NpcTileRec
	questIcon   int
	Buffs       []int
	BuffsEffect NpcBuffEffectRec
	Trapped     TrappedRec
	Paperdoll   []int
}

type MapNpcOffsetRec struct {
	X      int
	Y      int
	Type   byte
	Toggle bool
}

type TrappedRec struct {
	Sprite   int
	Duration int64
}

type NpcBuffEffectRec struct {
	Burned   bool
	Freezing bool
	Poisoned bool
}

type NpcTileRec struct {
	Type    byte
	Timer   int64
	Pos     objects.Vector2
	Alpha   byte
	fadeOut bool
}

type TradeItemRec struct {
	Item        int32
	CostItem    int32
	RequiredVar int32
	ItemValue   int64
	CostValue   int64
	Haircut     TradeItemHaircutRec
}

type TradeItemHaircutRec struct {
	Sprite int32
	Color  objects.Color
}

type ShopRec struct {
	Name      string
	BuyRate   int32
	TradeItem []TradeItemRec
	Pic       int32
	priceItem int32
	Order     byte
	Type      byte
}

type SpellRec struct {
	Name         string
	ESName       string
	Desc         string
	Dir          byte
	Type         byte
	AccessReq    byte
	MPCost       int32
	LevelReq     int32
	ClassReq     int32
	CastTime     int64
	CDTime       int64
	Icon         int32
	Map          int32
	X            int32
	Y            int32
	Vital        int64
	Duration     int64
	Interval     int64
	Range        int64
	IsAoE        bool
	AoE          int64
	CastAnim     int64
	SpellAnim    int64
	StunDuration int64
	BuffType     int64
	Frame        int64
	NextRank     int64
	NextUses     int64
	MaxMob       int64
	Stat         int64
	Cases        byte
	ElementBased bool
	ClassBasic   bool
	CastWalking  bool
	HPCost       int32
	Int          []int32
	Bool         []bool
}

type MapResourceRec struct {
	X             int32
	Y             int32
	ResourceState byte
	Frame         byte
	Shadow        byte
	FallDir       byte
	FallAngle     int32
	HitWood       float32
}

type ResourceRec struct {
	Name           string
	ESName         string
	SuccessMessage string
	ResourceType   byte
	ResourceImage  int32
	ExhaustedImage int32
	ToolRequired   byte
	Health         int64
	RespawnTime    int32
	Walkthrough    bool
	Animation      int32
	ItemReward     []int32
	ItemVal        []int64
	ItemLuck       []int32
	NormalAnim     bool
	NormalRandom   bool
	Int            []int32
}

type ActionMsgRec struct {
	Message string
	Created int64
	Type    int64
	Color   int64
	Scroll  int64
	X       int64
	Y       int64
	Timer   int64
}

type BloodRec struct {
	Sprite int64
	Timer  int64
	X      int64
	Y      int64
}

type AnimationRec struct {
	Name       string
	Sprite     []int64
	Frames     []int64
	LoopCount  []int64
	LoopTime   []int64
	XOffset    []int64
	YOffset    []int64
	Sound      string
	SoundValue int32
}

type AnimInstanceRec struct {
	Animation    int64
	X            int64
	Y            int64
	Duration     int64
	lockindex    int
	LockType     byte
	Timer        []int64
	Used         []bool
	LoopIndex    []int64
	FrameIndex   []int64
	MasterSprite int
	MasterPos    objects.Vector2
	MasterTimer  int64
}

type ButtonRec struct {
	Num            int
	state          byte
	X              int64
	Y              int64
	Width          int
	Height         int
	Visible        bool
	PicNum         int64
	Name           string
	Sound          string
	Color          string
	FlipHorizontal bool
	flipV          bool
}

type GUIWindowRec struct {
	Name                  string
	pic                   int
	Type                  int64
	X                     int
	Y                     int
	X2                    int64
	Y2                    int64
	defaultX              int
	defaultY              int
	Width                 int
	Height                int
	isString              bool
	Max                   int64
	button                []int64
	label                 []string
	labelX                []int64
	labelY                []int64
	Drag                  bool
	Press                 bool
	Escape                bool
	Up                    bool
	Zoom                  bool
	IsVisible             bool
	CloseX                int
	CloseY                int
	ScreenPosition        int
	defaultScreenPosition int
	UIName                int
	Text                  string
	Visible               bool
	CanWrite              bool
	Resized               objects.Vector2
}

type ChatBubbleRec struct {
	Msg        string
	Colour     objects.Color
	Target     int64
	TargetType byte
	timer      int64
	active     bool
	pos        objects.Vector2
	tName      string
}

type CharDataRec struct {
	Name      string
	Level     int32
	Classes   int32
	Sprite    int32
	Hair      int32
	Sex       byte
	HairTint  string
	NotHair   bool
	Equip     []int32
	CashEquip []int32
	Invisible []bool
}

type PlayerSpellRec struct {
	Spell  int32
	Uses   int32
	Master bool
}

type PlayerQuestRec struct {
	Status       byte
	ActualTask   byte
	TaskCount    byte
	CurrentCount int64
	QuestNum     int32
}

type HotbarRec struct {
	Slot int64
	Type byte
}

type MiniMapPlayerRec struct {
	X int64
	Y int64
}

type MiniMapNPCRec struct {
	X int64
	Y int64
}

type ServerAnnouncementRec struct {
	Message []string
	Timer   int64
}

type ProjectilRec struct {
	Name       string
	TravelTime int64
	Direction  int64
	X          int64
	Y          int64
	StateX     int64
	StateY     int64
	Pic        uint16
	Range      byte
	Damage     int32
	Speed      uint16
	Timer      int64
	Combo      int64
	Animation  uint16
	Light      bool
	Int        []int32
}

type ProjectilCacheRec struct {
	Num        uint16
	TravelTime int64
	Timer      int64
	Direction  byte
	Combo      byte
	X          int16
	Y          int16
	StateX     int16
	StateY     int16
}

type MoralRec struct {
	Name string
	Int  []int32
	Bool []bool
}

type StepsRec struct {
	Sprite int64
	timer  int64
	X      int64
	Y      int64
}

type FramesRec struct {
	Name          string
	Pic           int64
	PaperCheck    byte
	Frame1        int64
	Frame2        int64
	Frame3        int64
	Frame4        int64
	Timer1        int64
	Timer2        int64
	Timer3        int64
	Timer4        int64
	Paperdoll     int64
	PFrame1       int64
	PFrame2       int64
	PFrame3       int64
	PFrame4       int64
	Left1         int64
	Left2         int64
	Left3         int64
	Left4         int64
	KillPaperdoll byte
}

type CardsRec struct {
	Name   string
	Desc   string
	Num    int32
	Card   int32
	Sprite int32
	Icon   int32
	Offset objects.Vector2
}

type CraftingRec struct {
	Name      string
	Desc      string
	ItemReq   []int32
	ItemVal   []int32
	Reward    int32
	RewardVal int32
}

type NumberRec struct {
	Num     []string
	Width   []int64
	Pos     objects.Vector2
	Scroll  int64
	Created int64
	nColor  objects.Color
	aMsg    string
	Skin    byte
}

type MapSoundRec struct {
	pos         objects.Vector2
	SoundHandle string
	getTick     int64
	timer       int64
	Repeat      int
	SoundIndex  int
	isLoop      bool
	Destroy     bool
	Playing     bool
	delay       int64
}

type MapAnimationRec struct {
	Anim         []EachMapAnimationRec
	WeatherType  byte
	WeatherSpeed int64
	WeatherRate  int64
	EarthType    int64
	EarthSprite1 int64
	EarthSprite2 int64
	EarthSprite3 int64
	World        byte
}

type EachMapAnimationRec struct {
	Sprite     int
	X          int
	Y          int
	Size       byte
	Frames     byte
	Layer      byte
	Looop      int64
	XOffset    int64
	YOffset    int64
	Transport  int64
	CustomSize objects.Vector2
}

type FontRec struct {
	Ascii  []string
	Width  []int64
	X      int64
	Y      int64
	Scroll int64
	Active bool
	Leng   int64
}

type QuestTalkRec struct {
	Name string
	Lang []QuestLangTalkRec
}

type QuestLangTalkRec struct {
	Message []string
	Sprite  []int
}

type LanguageRec struct {
	Text []LanguageTextRec
}

type LanguageTextRec struct {
	Message  string
	Tutorial int
}

type GameFontsRec struct {
	Name string
}

type TitleRec struct {
	Name string
	Desc string
	Item int64
}

type TaskMsgRec struct {
	Msg       string
	FontColor objects.Color
	Created   int64
	dY        int
}

type NPCProjectilRec struct {
	Num        int64
	TravelTime int64
	timer      int64
	Direction  int64
	X          int64
	Y          int64
	stateX     int64
	stateY     int64
	Launched   bool
}

type LogMsgRec struct {
	Msg         string
	nColor      objects.Color
	Created     int64
	Transparent int
	dY          int64
}

type MapTopoRec struct {
	Name    string
	Depth   int64
	Item    []int64
	ItemVal []int64
	Luck    []int64
}

type ParticleRec struct {
	Type     int64
	X        int64
	Y        int64
	Velocity int64
	InUse    int64
	Sprite   int
	timer    int64
	alpha    int64
	rotation int64
	size     int64
	Movement int64
}

type LightRec struct {
	Name       string
	OffSet     objects.Vector2
	Color      objects.Color
	Ratio      int
	BlinkRatio float64
}

type ConditionRec struct {
	Name     string
	Line     []ActionRec
	MaxLines int32
	Sprite   ConditionSpriteRec
	ShowName bool
	OnJoin   bool
	MeetReq  string
	Faceset  int32
	Switch   int32
	Bool     []bool
}

type ConditionSpriteRec struct {
	pic     int32
	Type    int32
	myFrame objects.Vector2
	Mask    int32
	Width   int32
	Height  int32
	Offset  objects.Vector2
}

type ActionRec struct {
	Void       byte
	NextLine   int32
	ActionType byte
	EndCode    bool
	OrCode     bool
	AndCode    bool
	Modifier   byte
	ToEntity   byte
	ElseCode   bool
	isChoice   bool
	Var        []string
	VarType    []byte
	BreakPoint bool
}

type MapEditorDataRec struct {
	Data1 int
	Data2 int
	Data3 int
	Data4 string
}

type ArrowRec struct {
	Num   int32
	Value int32
}

type CashShopRec struct {
	Name string
	Slot []TradeCashItemRec
	Icon int32
}

type TradeCashItemRec struct {
	Item  int32
	Value int64
	Price int64
	Hot   bool
}

type RandomItemRec struct {
	Name string
	Item []RandomItemListRec
}

type RandomItemListRec struct {
	Num   int32
	Value int32
	Luck  byte
}

type HotItemRec struct {
	Category    int32
	SubCategory int32
}

type PinsRec struct {
	Item int32
	Desc string
}

type NewspaperRec struct {
	Title         string
	Url           string
	EventUrl      string
	playersOnline int32
	Message       string
	WorldBoss     []time.Time
	Cover         NewspaperCoverRec
	ServerDay     byte
	ServerMonth   byte
	selRank       byte
	selPlayer     byte
	selPoll       byte
	PollName      string
	Poll          []string
}

type NewspaperCoverRec struct {
	Url     string
	Texture []byte
}

type WorldBlessRec struct {
	Active    bool
	Owner     string
	Paperdoll string
	Time      int32
	Sprite    int32
	Hair      int32
	HairTint  string
	Value     float64
}

type SpeedFormulaRec struct {
	Normal float64
	Buff   float64
	Mount  float64
}
