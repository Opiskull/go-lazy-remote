package main

// FileService is used to register commands
import (
	"github.com/opiskull/go-jsonapi"
	"github.com/zenazn/goji/web"
)

// FileService
type FileService struct {
	Mux *web.Mux
}

// NewFileService creates a new instance
func NewFileService(dir string) *FileService {
	var mux = web.New()
	mux.Use(jsonapi.StaticFiles(dir))
	mux.Use(jsonapi.JSONRecovererMiddleware)
	mux.NotFound(jsonapi.JSONRouteNotFoundMiddleware)

	var service = &FileService{
		Mux: mux,
	}

	return service
}
