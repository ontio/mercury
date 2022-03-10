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
package utils

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCutDid(t *testing.T) {
	did := CutDId("did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr")
	assert.Equal(t, did, "did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr")
}

func TestGetPubKeyIndex(t *testing.T) {
	index := GetIndex("did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr@1#2")
	assert.Equal(t, index, "1")
}

func TestCutRouter(t *testing.T) {
	router := CutRouter("did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr@1#2")
	assert.Equal(t, router, "did:ont:TL9d9JddeyUZznz9eiTNwLEWQAipULr4mr#2")
}
