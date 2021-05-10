// Code generated by cmd/blockhash; DO NOT EDIT.

package block

const hashFarmland = 0
const hashLog = 1
const hashEndStone = 2
const hashShroomlight = 3
const hashAndesite = 4
const hashTorch = 5
const hashNetherBrickFence = 6
const hashCobblestone = 7
const hashLeaves = 8
const hashBricks = 9
const hashClay = 10
const hashQuartzBricks = 11
const hashNetherWart = 12
const hashGravel = 13
const hashGranite = 14
const hashLitPumpkin = 15
const hashCoalBlock = 16
const hashGrass = 17
const hashStainedGlass = 18
const hashGoldBlock = 19
const hashEndBrickStairs = 20
const hashGrassPlant = 21
const hashGlassPane = 22
const hashInvisibleBedrock = 23
const hashBeacon = 24
const hashNetheriteBlock = 25
const hashDirtPath = 26
const hashLapisOre = 27
const hashQuartzPillar = 28
const hashConcretePowder = 29
const hashQuartz = 30
const hashPotato = 31
const hashMelon = 32
const hashWater = 33
const hashAncientDebris = 34
const hashKelp = 35
const hashGlowstone = 36
const hashCake = 37
const hashCarpet = 38
const hashEmeraldOre = 39
const hashLantern = 40
const hashPlanks = 41
const hashLapisBlock = 42
const hashGoldOre = 43
const hashDirt = 44
const hashGlass = 45
const hashBlueIce = 46
const hashCoral = 47
const hashLava = 48
const hashDiamondBlock = 49
const hashSoulSoil = 50
const hashWheatSeeds = 51
const hashTerracotta = 52
const hashNetherQuartzOre = 53
const hashNoteBlock = 54
const hashCoralBlock = 55
const hashIronBars = 56
const hashCocoaBean = 57
const hashDiorite = 58
const hashObsidian = 59
const hashSand = 60
const hashBarrier = 61
const hashDiamondOre = 62
const hashPumpkin = 63
const hashSoulSand = 64
const hashGildedBlackstone = 65
const hashStainedTerracotta = 66
const hashWoodFenceGate = 67
const hashEndBricks = 68
const hashNetherrack = 69
const hashAir = 70
const hashWoodStairs = 71
const hashWool = 72
const hashEmeraldBlock = 73
const hashLight = 74
const hashStone = 75
const hashDragonEgg = 76
const hashWoodSlab = 77
const hashIronBlock = 78
const hashChest = 79
const hashCarrot = 80
const hashBeetrootSeeds = 81
const hashBedrock = 82
const hashSponge = 83
const hashStainedGlassPane = 84
const hashWoodTrapdoor = 85
const hashNetherGoldOre = 86
const hashIronOre = 87
const hashWoodFence = 88
const hashFire = 89
const hashGlazedTerracotta = 90
const hashSeaLantern = 91
const hashWoodDoor = 92
const hashBoneBlock = 93
const hashBasalt = 94
const hashCoalOre = 95
const hashPumpkinSeeds = 96
const hashConcrete = 97
const hashMelonSeeds = 98
const hashChiseledQuartz = 99

func (q Quartz) Hash() uint64         { return hashQuartz<<48 | uint64(boolByte(q.Smooth))<<0 }
func (LapisOre) Hash() uint64         { return hashLapisOre << 48 }
func (q QuartzPillar) Hash() uint64   { return hashQuartzPillar<<48 | uint64(q.Axis)<<0 }
func (c ConcretePowder) Hash() uint64 { return hashConcretePowder<<48 | uint64(c.Colour.Uint8())<<0 }
func (AncientDebris) Hash() uint64    { return hashAncientDebris << 48 }
func (k Kelp) Hash() uint64           { return hashKelp<<48 | uint64(k.Age)<<0 }
func (Glowstone) Hash() uint64        { return hashGlowstone << 48 }
func (p Potato) Hash() uint64         { return hashPotato<<48 | uint64(p.Growth)<<0 }
func (Melon) Hash() uint64            { return hashMelon << 48 }
func (w Water) Hash() uint64 {
	return hashWater<<48 | uint64(boolByte(w.Still))<<0 | uint64(w.Depth)<<1 | uint64(boolByte(w.Falling))<<9
}
func (l Lantern) Hash() uint64 {
	return hashLantern<<48 | uint64(boolByte(l.Hanging))<<0 | uint64(l.Type.Uint8())<<1
}
func (p Planks) Hash() uint64   { return hashPlanks<<48 | uint64(p.Wood.Uint8())<<0 }
func (LapisBlock) Hash() uint64 { return hashLapisBlock << 48 }
func (c Cake) Hash() uint64     { return hashCake<<48 | uint64(c.Bites)<<0 }
func (c Carpet) Hash() uint64   { return hashCarpet<<48 | uint64(c.Colour.Uint8())<<0 }
func (EmeraldOre) Hash() uint64 { return hashEmeraldOre << 48 }
func (BlueIce) Hash() uint64    { return hashBlueIce << 48 }
func (c Coral) Hash() uint64 {
	return hashCoral<<48 | uint64(c.Type.Uint8())<<0 | uint64(boolByte(c.Dead))<<4
}
func (l Lava) Hash() uint64 {
	return hashLava<<48 | uint64(boolByte(l.Still))<<0 | uint64(l.Depth)<<1 | uint64(boolByte(l.Falling))<<9
}
func (GoldOre) Hash() uint64      { return hashGoldOre << 48 }
func (d Dirt) Hash() uint64       { return hashDirt<<48 | uint64(boolByte(d.Coarse))<<0 }
func (Glass) Hash() uint64        { return hashGlass << 48 }
func (s WheatSeeds) Hash() uint64 { return hashWheatSeeds<<48 | uint64(s.Growth)<<0 }
func (DiamondBlock) Hash() uint64 { return hashDiamondBlock << 48 }
func (SoulSoil) Hash() uint64     { return hashSoulSoil << 48 }
func (c CoralBlock) Hash() uint64 {
	return hashCoralBlock<<48 | uint64(c.Type.Uint8())<<0 | uint64(boolByte(c.Dead))<<4
}
func (IronBars) Hash() uint64        { return hashIronBars << 48 }
func (c CocoaBean) Hash() uint64     { return hashCocoaBean<<48 | uint64(c.Facing)<<0 | uint64(c.Age)<<2 }
func (d Diorite) Hash() uint64       { return hashDiorite<<48 | uint64(boolByte(d.Polished))<<0 }
func (o Obsidian) Hash() uint64      { return hashObsidian<<48 | uint64(boolByte(o.Crying))<<0 }
func (Terracotta) Hash() uint64      { return hashTerracotta << 48 }
func (NetherQuartzOre) Hash() uint64 { return hashNetherQuartzOre << 48 }
func (n NoteBlock) Hash() uint64     { return hashNoteBlock << 48 }
func (p Pumpkin) Hash() uint64 {
	return hashPumpkin<<48 | uint64(boolByte(p.Carved))<<0 | uint64(p.Facing)<<1
}
func (s Sand) Hash() uint64           { return hashSand<<48 | uint64(boolByte(s.Red))<<0 }
func (Barrier) Hash() uint64          { return hashBarrier << 48 }
func (DiamondOre) Hash() uint64       { return hashDiamondOre << 48 }
func (GildedBlackstone) Hash() uint64 { return hashGildedBlackstone << 48 }
func (SoulSand) Hash() uint64         { return hashSoulSand << 48 }
func (Netherrack) Hash() uint64       { return hashNetherrack << 48 }
func (Air) Hash() uint64              { return hashAir << 48 }
func (s WoodStairs) Hash() uint64 {
	return hashWoodStairs<<48 | uint64(s.Wood.Uint8())<<0 | uint64(boolByte(s.UpsideDown))<<4 | uint64(s.Facing)<<5
}
func (w Wool) Hash() uint64       { return hashWool<<48 | uint64(w.Colour.Uint8())<<0 }
func (EmeraldBlock) Hash() uint64 { return hashEmeraldBlock << 48 }
func (t StainedTerracotta) Hash() uint64 {
	return hashStainedTerracotta<<48 | uint64(t.Colour.Uint8())<<0
}
func (f WoodFenceGate) Hash() uint64 {
	return hashWoodFenceGate<<48 | uint64(f.Wood.Uint8())<<0 | uint64(f.Facing)<<4 | uint64(boolByte(f.Open))<<6 | uint64(boolByte(f.Lowered))<<7
}
func (EndBricks) Hash() uint64 { return hashEndBricks << 48 }
func (l Light) Hash() uint64   { return hashLight<<48 | uint64(l.Level)<<0 }
func (IronBlock) Hash() uint64 { return hashIronBlock << 48 }
func (s Stone) Hash() uint64   { return hashStone<<48 | uint64(boolByte(s.Smooth))<<0 }
func (DragonEgg) Hash() uint64 { return hashDragonEgg << 48 }
func (s WoodSlab) Hash() uint64 {
	return hashWoodSlab<<48 | uint64(s.Wood.Uint8())<<0 | uint64(boolByte(s.Top))<<4 | uint64(boolByte(s.Double))<<5
}
func (b Bedrock) Hash() uint64       { return hashBedrock<<48 | uint64(boolByte(b.InfiniteBurning))<<0 }
func (s Sponge) Hash() uint64        { return hashSponge<<48 | uint64(boolByte(s.Wet))<<0 }
func (c Chest) Hash() uint64         { return hashChest<<48 | uint64(c.Facing)<<0 }
func (c Carrot) Hash() uint64        { return hashCarrot<<48 | uint64(c.Growth)<<0 }
func (b BeetrootSeeds) Hash() uint64 { return hashBeetrootSeeds<<48 | uint64(b.Growth)<<0 }
func (IronOre) Hash() uint64         { return hashIronOre << 48 }
func (w WoodFence) Hash() uint64     { return hashWoodFence<<48 | uint64(w.Wood.Uint8())<<0 }
func (f Fire) Hash() uint64          { return hashFire<<48 | uint64(f.Type.Uint8())<<0 | uint64(f.Age)<<4 }
func (t GlazedTerracotta) Hash() uint64 {
	return hashGlazedTerracotta<<48 | uint64(t.Colour.Uint8())<<0 | uint64(t.Facing)<<4
}
func (SeaLantern) Hash() uint64 { return hashSeaLantern << 48 }
func (p StainedGlassPane) Hash() uint64 {
	return hashStainedGlassPane<<48 | uint64(p.Colour.Uint8())<<0
}
func (t WoodTrapdoor) Hash() uint64 {
	return hashWoodTrapdoor<<48 | uint64(t.Wood.Uint8())<<0 | uint64(t.Facing)<<4 | uint64(boolByte(t.Open))<<6 | uint64(boolByte(t.Top))<<7
}
func (NetherGoldOre) Hash() uint64 { return hashNetherGoldOre << 48 }
func (CoalOre) Hash() uint64       { return hashCoalOre << 48 }
func (p PumpkinSeeds) Hash() uint64 {
	return hashPumpkinSeeds<<48 | uint64(p.Growth)<<0 | uint64(p.Direction)<<8
}
func (c Concrete) Hash() uint64 { return hashConcrete<<48 | uint64(c.Colour.Uint8())<<0 }
func (m MelonSeeds) Hash() uint64 {
	return hashMelonSeeds<<48 | uint64(m.Growth)<<0 | uint64(m.Direction)<<8
}
func (ChiseledQuartz) Hash() uint64 { return hashChiseledQuartz << 48 }
func (d WoodDoor) Hash() uint64 {
	return hashWoodDoor<<48 | uint64(d.Wood.Uint8())<<0 | uint64(d.Facing)<<4 | uint64(boolByte(d.Open))<<6 | uint64(boolByte(d.Top))<<7 | uint64(boolByte(d.Right))<<8
}
func (b BoneBlock) Hash() uint64 { return hashBoneBlock<<48 | uint64(b.Axis)<<0 }
func (b Basalt) Hash() uint64 {
	return hashBasalt<<48 | uint64(boolByte(b.Polished))<<0 | uint64(b.Axis)<<1
}
func (Shroomlight) Hash() uint64      { return hashShroomlight << 48 }
func (a Andesite) Hash() uint64       { return hashAndesite<<48 | uint64(boolByte(a.Polished))<<0 }
func (t Torch) Hash() uint64          { return hashTorch<<48 | uint64(t.Facing)<<0 | uint64(t.Type.Uint8())<<3 }
func (NetherBrickFence) Hash() uint64 { return hashNetherBrickFence << 48 }
func (c Cobblestone) Hash() uint64    { return hashCobblestone<<48 | uint64(boolByte(c.Mossy))<<0 }
func (f Farmland) Hash() uint64       { return hashFarmland<<48 | uint64(f.Hydration)<<0 }
func (l Log) Hash() uint64 {
	return hashLog<<48 | uint64(l.Wood.Uint8())<<0 | uint64(boolByte(l.Stripped))<<4 | uint64(l.Axis)<<5
}
func (EndStone) Hash() uint64     { return hashEndStone << 48 }
func (QuartzBricks) Hash() uint64 { return hashQuartzBricks << 48 }
func (n NetherWart) Hash() uint64 { return hashNetherWart<<48 | uint64(n.Age)<<0 }
func (l Leaves) Hash() uint64 {
	return hashLeaves<<48 | uint64(l.Wood.Uint8())<<0 | uint64(boolByte(l.Persistent))<<4 | uint64(boolByte(l.ShouldUpdate))<<5
}
func (Bricks) Hash() uint64           { return hashBricks << 48 }
func (c Clay) Hash() uint64           { return hashClay << 48 }
func (CoalBlock) Hash() uint64        { return hashCoalBlock << 48 }
func (Grass) Hash() uint64            { return hashGrass << 48 }
func (g StainedGlass) Hash() uint64   { return hashStainedGlass<<48 | uint64(g.Colour.Uint8())<<0 }
func (GoldBlock) Hash() uint64        { return hashGoldBlock << 48 }
func (Gravel) Hash() uint64           { return hashGravel << 48 }
func (g Granite) Hash() uint64        { return hashGranite<<48 | uint64(boolByte(g.Polished))<<0 }
func (l LitPumpkin) Hash() uint64     { return hashLitPumpkin<<48 | uint64(l.Facing)<<0 }
func (InvisibleBedrock) Hash() uint64 { return hashInvisibleBedrock << 48 }
func (Beacon) Hash() uint64           { return hashBeacon << 48 }
func (NetheriteBlock) Hash() uint64   { return hashNetheriteBlock << 48 }
func (DirtPath) Hash() uint64         { return hashDirtPath << 48 }
func (s EndBrickStairs) Hash() uint64 {
	return hashEndBrickStairs<<48 | uint64(boolByte(s.UpsideDown))<<0 | uint64(s.Facing)<<1
}
func (g GrassPlant) Hash() uint64 {
	return hashGrassPlant<<48 | uint64(boolByte(g.UpperPart))<<0 | uint64(g.Type.Uint8())<<1
}
func (GlassPane) Hash() uint64 { return hashGlassPane << 48 }