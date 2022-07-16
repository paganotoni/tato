package assets

import (
	"embed"
	"net/http"
)

var (
	//go:embed *.js
	fs     embed.FS
	Server = http.FileServer(http.FS(fs))
)
