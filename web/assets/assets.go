package assets

import (
	"embed"
	"net/http"
)

var (
	//go:embed *.js *.css
	fs     embed.FS
	Server = http.FileServer(http.FS(fs))
)
