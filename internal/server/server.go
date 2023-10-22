package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewServer() Server {
	s := Server{gin.Default()}
	return s
}
