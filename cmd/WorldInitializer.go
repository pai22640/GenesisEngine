package main

import (
	"fmt"
)

// World 定義了星球的基本結構。
type World struct {
	Name   string
	Biomes []Biome
	// 你可以按需添加更多屬性，例如Climate, Resources等
}

// Biome 描述了星球上的一種生物群系。
type Biome struct {
	Name      string
	Creatures []Creature
	Plants    []Plant
	Climate   ClimateType
	// 同樣，依據需求進一步拓展
}

// Creature 描述了生物群系中的生物。
type Creature struct {
	Species string
	// 其他生物特徵git status
}

// Plant 描述了生物群系中的植物生命。
type Plant struct {
	Species string
	// 其他植物特徵
}

// ClimateType 描述了氣候類型。
type ClimateType string

// 氣候類型常量。
const (
	Tropical  ClimateType = "Tropical"
	Desert    ClimateType = "Desert"
	Tundra    ClimateType = "Tundra"
	Temperate ClimateType = "Temperate"
)

// NewWorld 是 World 的建構函式，用於創建新的 World 實例。
func NewWorld(name string) *World {
	// 初始化一個有基本特徵的星球。
	// 實際情況，你需要一組更複雜的邏輯來決定這些特徵。
	return &World{
		Name: name,
		// 假設我們創建了一個生物多樣的星球.
		Biomes: []Biome{
			{Name: "Equatorial", Climate: Tropical},
			{Name: "Desert", Climate: Desert},
			// 添加更多生物群落...
		},
	}
}

func main() {
	// 產生一個新的星球。
	myWorld := NewWorld("Genesis")
	fmt.Printf("Welcome to %s, a new vibrant world!\n", myWorld.Name)
	// 繼續添加對 myWorld 的初始化以及擴展功能...
}
