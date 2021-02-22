package server

import (
	"encoding/base64"
	"fmt"
	"log"
	"myz-torrent-api/common"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type path string

func (s *Server) listFile(c *gin.Context) {
	root := s.conf.DownloadConfig.DownloadDir
	fs, err := common.ListFiles(root)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, fs)
}

func (s *Server) downloadFile(c *gin.Context) {
	root := s.conf.DownloadConfig.DownloadDir
	p := path(c.Param("path"))

	target, err := p.validate(root)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	info, err := os.Stat(target)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if info.IsDir() {
		fileName := info.Name() + ".zip"

		c.Writer.WriteHeader(http.StatusOK)
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
		c.Header("Content-Type", "application/zip")

		zip := common.NewZipWriter(c.Writer)
		defer zip.Close()

		if err := zip.AddPath(target); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%v", info.Name()))
		c.File(target)
	}
}

func (s *Server) deleteFile(c *gin.Context) {
	root := s.conf.DownloadConfig.DownloadDir
	p := path(c.Param("path"))
	targat, err := p.validate(root)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := common.DeleteFile(targat); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	}
}

func (p path) validate(root string) (string, error) {
	root, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("get root abs path %v error %v", root, err)
	}

	decoded, err := p.decode()
	if err != nil {
		return "", fmt.Errorf("path decode error %v", err)
	}

	res, err := filepath.Abs(decoded)
	if err != nil {
		return "", fmt.Errorf("get target abs path %v error %v", res, err)
	}

	if !strings.HasPrefix(res, root) {
		return "", fmt.Errorf("invaild path %v", res)
	}

	return res, nil
}

// path => base64 => urlEncode => origin
func (p path) decode() (string, error) {
	b64Decode, err := base64.StdEncoding.DecodeString(string(p))
	if err != nil {
		return "", err
	}

	uDecode, err := url.Parse(string(b64Decode))
	if err != nil {
		return "", err
	}

	return uDecode.Path, nil
}
