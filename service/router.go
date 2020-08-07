/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package service

import (
	"github.com/ontio/mercury/common/packager/ecdsa"
	"github.com/ontio/mercury/service/common"
	"github.com/ontio/mercury/service/controller"
	"github.com/ontio/mercury/store"
	"github.com/ontio/mercury/vdri"
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
