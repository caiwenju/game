package maps

// XNUm 地图每行图片的数量
const XNUm = 13

// YNUm 地图每列图片的数量
const YNUm = 13

// ImageLarge 图片长宽
const ImageLarge = 47

// Panel 控制面板宽度
const Panel = 47

// PanelNum 控制面板数量
const PanelNum = 4

// FloorNum 楼层数量
const FloorNum = 50

// Map 地图初始数据
var Map = [FloorNum][YNUm][XNUm]int{
	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 8, 1, 20, 21, 20, 1, 1, 1, 1, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2},
		{2, 83, 1, 1, 5, 1, 2, 85, 80, 1, 2, 1, 2},
		{2, 1, 27, 1, 2, 1, 2, 86, 83, 1, 2, 1, 2},
		{2, 2, 5, 2, 2, 1, 2, 2, 2, 5, 2, 1, 2},
		{2, 80, 1, 1, 2, 1, 5, 24, 30, 24, 2, 1, 2},
		{2, 1, 28, 1, 2, 1, 2, 2, 2, 2, 2, 1, 2},
		{2, 2, 5, 2, 2, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 2, 2, 5, 2, 2, 2, 5, 2, 2},
		{2, 83, 1, 80, 2, 80, 1, 1, 2, 1, 24, 1, 2},
		{2, 83, 1, 80, 2, 1, 1, 1, 2, 20, 84, 20, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 9, 1, 6, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 2, 2, 1, 38, 1, 38, 1, 2, 2, 2},
		{2, 1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2},
		{2, 1, 2, 80, 80, 2, 1, 1, 1, 2, 1, 60, 2},
		{2, 1, 2, 80, 1, 10, 1, 1, 1, 10, 1, 1, 2},
		{2, 1, 2, 2, 2, 2, 1, 1, 1, 2, 2, 2, 2},
		{2, 1, 2, 62, 1, 2, 1, 1, 1, 2, 1, 61, 2},
		{2, 1, 2, 1, 1, 10, 1, 1, 1, 10, 1, 1, 2},
		{2, 1, 2, 2, 2, 2, 1, 1, 1, 2, 2, 2, 2},
		{2, 1, 2, 84, 84, 2, 1, 1, 1, 2, 1, 1, 2},
		{2, 8, 2, 84, 1, 10, 1, 1, 1, 10, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 80, 86, 2, 80, 84, 80, 2, 1, 2, 1, 83, 2},
		{2, 1, 83, 2, 84, 80, 84, 2, 1, 5, 24, 1, 2},
		{2, 30, 1, 2, 80, 81, 80, 2, 1, 2, 2, 2, 2},
		{2, 5, 2, 2, 2, 1, 2, 2, 1, 2, 1, 60, 2},
		{2, 1, 1, 24, 1, 1, 1, 20, 1, 1, 1, 1, 2},
		{2, 5, 2, 2, 1, 1, 1, 2, 1, 2, 2, 2, 2},
		{2, 27, 1, 2, 2, 63, 2, 2, 1, 2, 1, 83, 2},
		{2, 1, 80, 2, 1, 1, 1, 2, 1, 5, 30, 80, 2},
		{2, 83, 85, 2, 1, 1, 1, 2, 1, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 1, 2, 2, 21, 2, 1, 1, 2},
		{2, 9, 1, 1, 1, 1, 1, 2, 1, 5, 1, 8, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 1, 81, 1, 2, 65, 66, 67, 2, 1, 60, 1, 2},
		{2, 83, 1, 80, 2, 1, 1, 1, 2, 80, 1, 84, 2},
		{2, 1, 1, 1, 2, 1, 1, 1, 2, 1, 28, 1, 2},
		{2, 2, 5, 2, 2, 2, 6, 2, 2, 2, 5, 2, 2},
		{2, 1, 30, 1, 5, 1, 21, 1, 1, 27, 1, 1, 2},
		{2, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 21, 1, 20, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 5, 2, 2, 5, 2, 2, 2, 5, 2, 2, 5, 2},
		{2, 1, 2, 1, 24, 1, 2, 1, 30, 1, 2, 1, 2},
		{2, 1, 2, 20, 1, 80, 2, 85, 1, 83, 2, 1, 2},
		{2, 8, 2, 80, 20, 80, 2, 1, 20, 1, 2, 9, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 8, 2, 1, 21, 5, 1, 2, 1, 1, 5, 1, 2},
		{2, 1, 2, 1, 1, 2, 80, 2, 20, 20, 2, 21, 2},
		{2, 1, 5, 24, 1, 2, 1, 2, 80, 80, 2, 1, 2},
		{2, 2, 2, 2, 5, 2, 24, 2, 80, 80, 2, 1, 2},
		{2, 80, 1, 30, 1, 2, 1, 2, 2, 2, 2, 1, 2},
		{2, 80, 1, 1, 24, 2, 1, 20, 1, 1, 1, 1, 2},
		{2, 2, 28, 2, 2, 2, 1, 2, 2, 2, 2, 21, 2},
		{2, 1, 1, 1, 1, 2, 20, 2, 1, 1, 1, 1, 2},
		{2, 86, 80, 83, 98, 2, 1, 2, 5, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 1, 2, 1, 2, 1, 1, 2},
		{2, 9, 1, 1, 1, 1, 1, 2, 1, 2, 1, 87, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 9, 2, 80, 80, 2, 1, 30, 1, 80, 20, 1, 2},
		{2, 1, 2, 80, 80, 2, 1, 2, 2, 2, 2, 5, 2},
		{2, 1, 2, 2, 21, 2, 1, 2, 83, 1, 27, 1, 2},
		{2, 1, 5, 5, 1, 5, 1, 2, 61, 1, 1, 24, 2},
		{2, 1, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2},
		{2, 1, 1, 21, 30, 1, 80, 1, 27, 28, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 1, 2},
		{2, 30, 1, 1, 60, 2, 1, 5, 5, 1, 5, 1, 2},
		{2, 1, 24, 1, 86, 2, 1, 2, 2, 21, 2, 21, 2},
		{2, 5, 2, 2, 2, 2, 1, 2, 1, 1, 2, 1, 2},
		{2, 1, 20, 1, 1, 27, 1, 2, 83, 83, 2, 8, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 8, 2, 85, 2, 1, 61, 1, 2, 80, 2, 20, 2},
		{2, 1, 2, 83, 2, 1, 1, 1, 2, 80, 2, 21, 2},
		{2, 1, 2, 24, 2, 21, 2, 28, 2, 83, 2, 20, 2},
		{2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2},
		{2, 5, 2, 5, 2, 6, 2, 5, 2, 27, 2, 5, 2},
		{2, 1, 28, 1, 30, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 5, 2, 5, 2, 5, 2, 5, 2, 28, 2, 5, 2},
		{2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2},
		{2, 1, 2, 1, 2, 24, 2, 21, 2, 84, 2, 1, 2},
		{2, 20, 2, 20, 2, 80, 2, 30, 2, 80, 2, 1, 2},
		{2, 1, 21, 1, 2, 80, 2, 84, 2, 80, 2, 9, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 9, 1, 5, 5, 1, 8, 1, 2, 80, 1, 80, 2},
		{2, 1, 1, 2, 2, 1, 1, 20, 2, 1, 82, 1, 2},
		{2, 5, 2, 2, 2, 2, 5, 2, 2, 84, 1, 83, 2},
		{2, 1, 2, 80, 80, 80, 1, 1, 2, 2, 11, 2, 2},
		{2, 83, 2, 2, 2, 2, 2, 30, 2, 37, 1, 37, 2},
		{2, 1, 21, 20, 21, 1, 2, 1, 2, 1, 1, 1, 2},
		{2, 2, 2, 2, 2, 5, 2, 24, 2, 2, 5, 2, 2},
		{2, 1, 1, 1, 24, 1, 27, 1, 30, 1, 1, 1, 2},
		{2, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 5, 2},
		{2, 20, 1, 2, 85, 80, 2, 81, 83, 2, 1, 27, 2},
		{2, 1, 24, 6, 80, 86, 2, 80, 1, 6, 28, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 1, 1, 27, 5, 1, 9, 1, 5, 20, 1, 83, 2},
		{2, 1, 80, 1, 2, 1, 1, 1, 2, 1, 20, 1, 2},
		{2, 28, 2, 2, 2, 2, 6, 2, 2, 2, 2, 1, 2},
		{2, 1, 80, 1, 2, 80, 1, 80, 5, 5, 1, 1, 2},
		{2, 86, 1, 24, 5, 1, 85, 1, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 21, 2, 1, 1, 28, 2},
		{2, 80, 1, 5, 28, 80, 2, 1, 2, 88, 2, 1, 2},
		{2, 28, 1, 2, 1, 1, 2, 1, 2, 2, 2, 5, 2},
		{2, 5, 2, 2, 2, 5, 2, 1, 2, 80, 1, 30, 2},
		{2, 1, 83, 2, 1, 27, 2, 24, 2, 1, 27, 1, 2},
		{2, 8, 1, 6, 1, 1, 5, 1, 5, 30, 1, 83, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		{2, 2, 2, 2, 2, 1, 1, 1, 2, 2, 2, 2, 2},
		{2, 27, 27, 27, 2, 2, 1, 2, 2, 27, 27, 27, 2},
		{2, 1, 28, 1, 11, 1, 29, 1, 11, 1, 28, 1, 2},
		{2, 2, 2, 2, 2, 1, 1, 1, 2, 2, 2, 2, 2},
		{2, 27, 86, 27, 2, 2, 1, 2, 2, 27, 85, 27, 2},
		{2, 1, 28, 1, 2, 2, 1, 2, 2, 1, 28, 1, 2},
		{2, 1, 1, 1, 2, 2, 1, 2, 2, 1, 1, 1, 2},
		{2, 5, 2, 5, 2, 2, 7, 2, 2, 5, 2, 5, 2},
		{2, 1, 2, 1, 2, 1, 1, 1, 2, 1, 2, 1, 2},
		{2, 9, 2, 1, 30, 1, 1, 1, 30, 1, 2, 84, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	},
}

// MapWidth 地图长宽
const MapWidth = XNUm * ImageLarge
const MapHeight = YNUm*ImageLarge + Panel*PanelNum

// Image 快类型数据
var Image = map[int]map[string]interface{}{
	1:  {"imagePath": "", "type": "road"},
	2:  {"imagePath": "", "type": "wall", "canDestroy": false},
	5:  {"imagePath": "", "type": "door", "color": "yellow"},
	6:  {"imagePath": "", "type": "door", "color": "blue"},
	7:  {"imagePath": "", "type": "door", "color": "red"},
	8:  {"imagePath": "", "type": "stairs", "direct": "up"},
	9:  {"imagePath": "", "type": "stairs", "direct": "down"},
	11: {"imagePath": "", "type": "door", "color": "master"},
	12: {"imagePath": "", "type": "door", "direct": "master"},
	20: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	21: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	24: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	27: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	28: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	29: {"imagePath": "", "type": "master", "attack": 20, "defense": 20, "health": 100},
	30: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	37: {"imagePath": "", "type": "master", "attack": 50, "defense": 10, "health": 100},
	38: {"imagePath": "", "type": "master", "attack": 50, "defense": 10, "health": 100},
	60: {"imagePath": "", "type": "npc"},
	61: {"imagePath": "", "type": "npc"},
	62: {"imagePath": "", "type": "npc"},
	63: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	64: {"imagePath": "", "type": "master", "attack": 10, "defense": 10, "health": 100},
	65: {"imagePath": "", "type": "store"},
	66: {"imagePath": "", "type": "store"},
	67: {"imagePath": "", "type": "store"},
	80: {"imagePath": "", "type": "Key", "color": "yellow", "num": 1},
	81: {"imagePath": "", "type": "Key", "color": "blue", "num": 1},
	82: {"imagePath": "", "type": "Key", "color": "red", "num": 1},
	83: {"imagePath": "", "type": "medicine", "health": 1000},
	84: {"imagePath": "", "type": "medicine", "health": 2000},
	85: {"imagePath": "", "type": "tone", "Attack": 10},
	86: {"imagePath": "", "type": "tone", "Defence": 10},
	87: {"imagePath": "", "type": "tone", "Attack": 1000},
	88: {"imagePath": "", "type": "tone", "Defence": 1000},
	98: {"imagePath": "", "type": "tone"},
}

// PersonPosition 每层人物的起始位置
var PersonPosition = map[int][2]int{
	0:  {6, 11},
	1:  {6, 11},
	2:  {1, 2},
	3:  {2, 11},
	4:  {11, 10},
	5:  {2, 11},
	6:  {1, 2},
	7:  {11, 10},
	8:  {1, 2},
	9:  {6, 2},
	10: {1, 10},
	11: {6, 11},
	12: {6, 11},
	13: {6, 11},
	14: {6, 11},
	15: {6, 11},
	16: {6, 11},
	17: {6, 11},
	18: {6, 11},
	19: {6, 11},
	20: {6, 11},
}

// BackGroundPanel 背景面板
var BackGroundPanel = [4][XNUm]int{
	{99, 104, 104, 99, 99, 99, 99, 103, 104, 104, 99, 99, 99},
	{100, 104, 104, 99, 99, 99, 99, 80, 104, 104, 99, 99, 99},
	{101, 104, 104, 99, 99, 99, 99, 81, 104, 104, 99, 99, 99},
	{102, 104, 104, 99, 99, 99, 99, 82, 104, 104, 99, 99, 99},
}
