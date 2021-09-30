package server

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
)

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

// StaticAssetsHandler returns an http.Handler that will serve files from the Assets embed.FS.
// Used to embed static CSS and image files into the templated pages.
func StaticAssetsHandler(f embed.FS, prefix, root string) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		return f.Open(assetPath)
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}
