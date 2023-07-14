package web

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var assets embed.FS

// Assets contains either the built web files from web/dist directory
// or it is empty.
var Assets = func() fs.FS {
	dist, err := fs.Sub(assets, "dist")
	if err != nil {
		panic(err)
	}
	return dist
}()
