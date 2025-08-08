package luna

import "embed"

//go:embed build/runtimes/*
var EmbeddedRuntimes embed.FS
