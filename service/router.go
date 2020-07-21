package service

import (
	"git.ont.io/ontid/otf/common/packager/ecdsa"
	"git.ont.io/ontid/otf/service/common"
	"git.ont.io/ontid/otf/service/controller"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/vdri"
	"github.com/gin-gonic/gin"
)

func NewApiRouter(packager *ecdsa.Packager, store store.Store,
	msgSvr *common.MsgService, v vdri.VDRI) *gin.Engine {

	systemController := controller.NewSystemController(packager, store, msgSvr)
	credentialController := controller.NewCredentialController(packager, store, msgSvr, v)
	presentationController := controller.NewPresentationController(packager, store, msgSvr, v)
	return NewRouter(credentialController, systemController, presentationController)
}

func NewRouter(routers ...common.Router) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	for _, api := range routers {
		for _, route := range api.Routes() {
			router.Handle(route.Method, route.Pattern, route.HandlerFunc)
		}
	}
	return router
}
