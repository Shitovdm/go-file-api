package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Shitovdm/go-file-api/helpers"
)

// Server instance
type Server struct {
	BasePath string
	manifest helpers.FileItems
}

// NewServer instance
func NewServer(basePath string) *Server {
	return &Server{
		BasePath: basePath,
	}
}

// StartServe - start the file server
func (s *Server) StartServe(port string) error {
	// pre calculate the file manifest
	var err error
	if s.manifest, err = helpers.GetFileItems(s.BasePath); err != nil {
		return err
	}

	router := http.NewServeMux()
	router.HandleFunc("/", s.getFile)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	return httpServer.ListenAndServe()
}

// Unused
func (s *Server) filesRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := r.URL.Query().Get("path")
		pathParts := strings.Split(path, "/")
		files := s.manifest
		if path != "" {
			for _, part := range pathParts {
				if part == "." {
					continue
				}
				if file, ok := files[part]; ok {
					files = file.Items
				} else {
					s.writeMessage(w, http.StatusNotFound, "not found")
					return
				}
			}
		}
		s.writeMessage(w, http.StatusOK, files)
	}
}

func (s *Server) getFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		file, _ := os.Open(s.parsePath(r))
		defer file.Close()
		n, err := io.Copy(w, file)
		if err != nil {
			s.writeMessage(w, http.StatusOK, fmt.Sprintf("File .%s not found!", r.URL.Path))
		}else {
			w.Header().Set("Content-Length", strconv.FormatInt(n, 10))
			w.Header().Set("Content-Type", "application/octet-stream")
		}
	}
}

func (s *Server) parsePath(r *http.Request) string {
	path := r.URL.Path
	path = strings.Replace(path, "..", ".", -1)
	path = filepath.Join(s.BasePath, path)
	return path
}

func (s *Server) writeMessage(w http.ResponseWriter, status int, message interface{}) {
	switch t := message.(type) {
	case string:
		w.WriteHeader(status)
		_, _ = w.Write([]byte(t))
	case *string:
		w.WriteHeader(status)
		_, _ = w.Write([]byte(*t))
	case []byte:
		w.WriteHeader(status)
		_, _ = w.Write(t)
	default:
		if data, err := json.Marshal(message); err == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(status)
			_, _ = w.Write(data)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
		}
	}
}
