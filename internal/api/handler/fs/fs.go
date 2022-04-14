package fs

import (
	"net/http"
	"path/filepath"

	"github.com/GoodUncleFood/gu/internal/api/control/handler"
)

type Params struct {
	Dir    string
	Prefix string
}

func New(params Params) handler.RequestHandler {
	return func(
		p handler.Params, r *http.Request, w http.ResponseWriter,
	) error {
		pfs := &PrivateFS{FS: http.Dir(params.Dir)}

		http.StripPrefix(
			params.Prefix, http.FileServer(pfs),
		).ServeHTTP(w, r)

		return nil
	}
}

// PrivateFS prevents directory listings.
type PrivateFS struct {
	FS http.FileSystem
}

func (p PrivateFS) Open(path string) (http.File, error) {
	f, err := p.FS.Open(path)
	if err != nil {
		return nil, err
	}

	// If the directory contains an index, return it.
	info, err := f.Stat()
	if info.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := p.FS.Open(index); err != nil {
			// Close the directory.
			if err := f.Close(); err != nil {
				return nil, err
			}

			// Handle missing index or failure to open.
			return nil, err
		}
	}

	// Not a directory. Return the file.
	return f, nil
}
