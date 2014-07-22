package agent

import (
	"github.com/porty/command-and-control/bundled"
	"net/http"
	"path/filepath"
	"strconv"
)

var mimeTypes = map[string]string{
	".css":  "text/css; charset=utf-8",
	".gif":  "image/gif",
	".htm":  "text/html; charset=utf-8",
	".html": "text/html; charset=utf-8",
	".jpg":  "image/jpeg",
	".js":   "application/javascript",
	".json": "application/json",
	".pdf":  "application/pdf",
	".png":  "image/png",
	".xml":  "text/xml; charset=utf-8",
	".eot":  "application/vnd.ms-fontobject",
	".svg":  "image/svg+xml",
	".ttf":  "application/x-font-ttf",
	".woff": "application/font-woff",
	".otf":  "application/x-font-opentype",
}

type AssetsServer struct {
}

func getMime(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return ""
	}
	return mimeTypes[ext]
}

func (a AssetsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	if path == "" {
		path = "index.html"
	}

	b, err := bundled.Asset(path)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	mime := getMime(path)
	if mime != "" {
		w.Header().Set("Content-type", mime)
	}

	w.Header().Set("Content-length", strconv.Itoa(len(b)))

	w.Write(b)
}
