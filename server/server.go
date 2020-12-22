package server

import (
	"fmt"
	"myz-torrent-api/aria2"
	"myz-torrent-api/common"
	"myz-torrent-api/downloader"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Server app context
type Server struct {
	r     *gin.Engine
	conf  *common.Config
	dmg   *downloader.DownloadManager
	aria2 *aria2.Aria2
}

// Run server
func (s *Server) Run() error {
	if err := s.initConfig(); err != nil {
		return err
	}

	if err := s.initDownloader(); err != nil {
		return err
	}

	s.initRouter()

	return s.r.Run(fmt.Sprintf("%v:%v", s.conf.ServerAddr, s.conf.ServerPortal))
}

func (s *Server) initRouter() {
	r := gin.Default()

	r.GET("job", s.listJob)
	r.GET("active_job", s.listActiveJob)
	r.GET("waiting_job", s.listWaitingJob)
	r.GET("stopped_job", s.listStoppedJob)
	r.POST("add_uri", s.addURI)
	r.POST("download/:id", s.startJob)
	r.PUT("download/:id", s.pauseJob)
	r.DELETE("download/:id", s.deleteJob)

	r.GET("file", s.listFile)
	r.GET("file/:path", s.downloadFile)
	r.DELETE("file/:path", s.deleteFile)

	s.r = r
}

func (s *Server) initConfig() error {
	c, err := common.LoadConfig()
	if err != nil {
		return err
	}

	s.conf = c

	// Create log files
	if len(s.conf.LogPath) > 0 {
		if gin.DefaultWriter, err = os.Create(filepath.Join(s.conf.LogPath, "request.log")); err != nil {
			return fmt.Errorf("error create log file: %v", filepath.Join(s.conf.LogPath, "requests.log"))
		}

		if gin.DefaultErrorWriter, err = os.Create(filepath.Join(s.conf.LogPath, "error.log")); err != nil {
			return fmt.Errorf("error create log file: %v", filepath.Join(s.conf.LogPath, "error.log"))
		}
	}

	return nil
}

func (s *Server) initDownloader() error {
	if s.conf.DownloadConfig.Aria2Address != "" {
		s.aria2 = aria2.NewAria2WithSecret(fmt.Sprintf("%s/jsonrpc", s.conf.DownloadConfig.Aria2Address), s.conf.DownloadConfig.Aria2secret)
	}
	return nil
}
