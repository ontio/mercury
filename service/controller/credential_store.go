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

package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ontio/mercury/common/message"
)

const (
	CredentialKey        = "Credential"
	RequestCredentialKey = "RequestCredential"
	OfferCredentialKey   = "OfferCredential"
)

func (c *CredentialController) SaveOfferCredential(did, id string, propsal *message.OfferCredential) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", OfferCredentialKey, did, id))
	b, err := c.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	data, err := json.Marshal(propsal)
	if err != nil {
		return err
	}
	return c.store.Put(key, data)
}

func (c *CredentialController) SaveCredential(did, id string, credential message.IssueCredential) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", CredentialKey, did, id))
	b, err := c.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	rec := message.CredentialRec{
		OwnerDID:   credential.Connection.TheirDid,
		Credential: credential,
		Timestamp:  time.Now(),
	}
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return c.store.Put(key, data)
}

func (c *CredentialController) SaveRequestCredential(did, id string, requestCredential message.RequestCredential) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestCredentialKey, did, id))
	b, err := c.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	rec := message.RequestCredentialRec{
		RequesterDID:      requestCredential.Connection.MyDid,
		RequestCredential: requestCredential,
		State:             message.RequestCredentialReceived,
	}
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return c.store.Put(key, data)
}

func (c *CredentialController) QueryCredentialFromStore(did, id string) (message.IssueCredential, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", CredentialKey, did, id))

	data, err := c.store.Get(key)
	if err != nil {
		return message.IssueCredential{}, err
	}

	rec := new(message.CredentialRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return message.IssueCredential{}, err
	}
	return rec.Credential, nil
}

func (c *CredentialController) UpdateRequestCredential(did, id string, state message.RequestCredentialState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestCredentialKey, did, id))
	data, err := c.store.Get(key)
	if err != nil {
		return err
	}

	rec := new(message.RequestCredentialRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}
	if rec.State >= state {
		return fmt.Errorf("UpdateRequestCredential id :%s state invalid\n", id)
	}
	rec.State = state
	data, err = json.Marshal(rec)
	if err != nil {
		return err
	}
	return c.store.Put(key, data)
}

func (c *CredentialController) DelRequestCredential(did, id string) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestCredentialKey, did, id))
	return c.store.Delete(key)
}

func (c *CredentialController) DelCredential(did, id string) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", CredentialKey, did, id))
	return c.store.Delete(key)
}
