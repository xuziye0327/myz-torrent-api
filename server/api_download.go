package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) listActiveJob(c *gin.Context) {
	if stats, err := s.aria2.TellActive(); err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.JSON(http.StatusOK, stats)
	}
}

func (s *Server) listWaitingJob(c *gin.Context) {
	if stats, err := s.aria2.TellWaiting(0, 10); err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.JSON(http.StatusOK, stats)
	}
}

func (s *Server) listStoppedJob(c *gin.Context) {
	if stats, err := s.aria2.TellStopped(0, 10); err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.JSON(http.StatusOK, stats)
	}
}

func (s *Server) addURI(c *gin.Context) {
	var links []string
	if err := c.BindJSON(&links); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "magnet is empty",
		})
		return
	}

	if err := s.aria2.AddURI(links...); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": links,
	})
}

func (s *Server) startJob(c *gin.Context) {
	id := c.Param("id")
	s.dmg.Strart(id)
	c.JSON(http.StatusOK, nil)
}

func (s *Server) pauseJob(c *gin.Context) {
	id := c.Param("id")
	s.dmg.Pause(id)
	c.JSON(http.StatusOK, nil)
}

func (s *Server) deleteJob(c *gin.Context) {
	id := c.Param("id")
	s.dmg.Delete(id)
	c.JSON(http.StatusOK, nil)
}
