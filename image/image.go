package image

import "embed"

//go:embed main/*
var MainFs embed.FS

//go:embed fight/*
var FightFs embed.FS
