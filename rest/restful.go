package rest

import (
	"git.ont.io/ontid/otf/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Invite(c *gin.Context) {
	resp := Gin{C: c}
	invite := &Invitation{}
	err := c.Bind(invite)
	if err != nil {
		middleware.Log.Errorf("Invite err:%s", err)
	}
	//todo
	resp.Response(http.StatusOK, 0, nil)
}

func Connect(c *gin.Context) {
	resp := Gin{C: c}
	invite := &ConnectionRequest{}
	err := c.Bind(invite)
	if err != nil {
		middleware.Log.Errorf("ConnectionRequest err:%s", err)
	}
	//todo
	resp.Response(http.StatusOK, 0, nil)
}

func SendCredential(c *gin.Context) {
	resp := Gin{C: c}
	invite := &ProposalCredential{}
	err := c.Bind(invite)
	if err != nil {
		middleware.Log.Errorf("ProposalCredential err:%s", err)
	}
	//todo
	resp.Response(http.StatusOK, 0, nil)
}

func IssueCredential(c *gin.Context) {
	resp := Gin{C: c}
	invite := &OfferCredential{}
	err := c.Bind(invite)
	if err != nil {
		middleware.Log.Errorf("OfferCredential err:%s", err)
	}
	//todo
	resp.Response(http.StatusOK, 0, nil)
}

func RequestProof(c *gin.Context) {
	resp := Gin{C: c}
	invite := &RequestCredential{}
	err := c.Bind(invite)
	if err != nil {
		middleware.Log.Errorf("RequestCredential err:%s", err)
	}
	//todo
	resp.Response(http.StatusOK, 0, nil)
}

func PresentProof(c *gin.Context) {
	resp := Gin{C: c}
	invite := &Presentation{}
	err := c.Bind(invite)
	if err != nil {
		middleware.Log.Errorf("Presentation err:%s", err)
	}
	//todo
	resp.Response(http.StatusOK, 0, nil)
}
