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

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterLastIndexOf(t *testing.T) {
	did := "did:ont:abcde"
	routers := []string{"did:ont:12345#1", "did:ont:abcde@1"}
	i, err := RouterLastIndexOf(did, routers)
	assert.Nil(t, err)
	assert.Equal(t, i, 1)
	routers = append(routers, "did:ont:ccadfa#1")
	i, err = RouterLastIndexOf(did, routers)
	assert.Nil(t, err)
	assert.Equal(t, i, 1)
	routers = []string{"did:ont:abcde@1"}
	i, err = RouterLastIndexOf(did, routers)
	assert.Nil(t, err)
	assert.Equal(t, i, 0)
	routers = []string{"did:ont:ccadfa#1", "did:ont:12345#1", "did:ont:abcde@1"}
	i, err = RouterLastIndexOf(did, routers)
	assert.Nil(t, err)
	assert.Equal(t, i, 2)
	routers = []string{}
	i, err = RouterLastIndexOf(did, routers)
	assert.NotNil(t, err)
}

func TestIsReceiver(t *testing.T) {
	did := "did:ont:abcde"
	routers := []string{"did:ont:12345#1", "did:ont:abcde@1"}
	f := IsReceiver(did, routers)
	assert.True(t, f)

	routers = []string{"did:ont:abcde@1"}
	f = IsReceiver(did, routers)
	assert.True(t, f)

	routers = []string{"did:ont:abcde#1", "did:ont:121213"}
	f = IsReceiver(did, routers)
	assert.False(t, f)
}

func TestMergeRouter(t *testing.T) {
	myrouter := []string{"did:ont:abcde#1", "did:ont:ccccc#1"}
	theirrouter := []string{"did:ont:ddddd@1", "did:ont:eeeee#1"}
	res := MergeRouter(myrouter, theirrouter)
	assert.Equal(t, len(res), 4)
	l, _ := RouterLastIndexOf("did:ont:ddddd", res)
	assert.Equal(t, l, 3)
}
