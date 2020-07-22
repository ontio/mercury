package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"git.ont.io/ontid/otf/common/message"
)

const (
	RequestPresentationKey = "RequestPresentation"
	PresentationKey        = "Presentation"
)

func (p *PresentationController) SaveRequestPresentation(did, id string, rr message.RequestPresentation) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestPresentationKey, did, id))
	b, err := p.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("ReqeustPresentation id:%s,all ready exist", id)
	}

	rec := new(message.RequestPresentationRec)
	rec.RerquestPrentation = rr
	rec.RequesterDID = rr.Connection.MyDid
	rec.State = message.RequestPresentationReceived

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return p.store.Put(key, data)
}

func (p *PresentationController) UpdateRequestPresentaion(did, id string, state message.RequestPresentationState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestPresentationKey, did, id))
	data, err := p.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(message.RequestPresentationRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}
	if rec.State <= state {
		return fmt.Errorf("request presentation id:%s state invalid", id)
	}

	rec.State = state
	data, err = json.Marshal(rec)
	if err != nil {
		return err
	}
	return p.store.Put(key, data)
}

func (p *PresentationController) SavePresentation(did, id string, pr message.Presentation) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", PresentationKey, did, id))
	b, err := p.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("ReqeustPresentation id:%s,all ready exist", id)
	}

	rec := new(message.PresentationRec)
	rec.Presentation = pr
	rec.OwnerDID = pr.Connection.TheirDid
	rec.Timestamp = time.Now()

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return p.store.Put(key, data)
}

func (p *PresentationController) QueryPresentationFromStore(did, id string) (message.Presentation, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", PresentationKey, did, id))
	data, err := p.store.Get(key)
	if err != nil {
		return message.Presentation{}, err
	}
	rec := new(message.PresentationRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return message.Presentation{}, err
	}
	return rec.Presentation, nil
}

func (p *PresentationController) DelPresentation(did,id string )error {
	key := []byte(fmt.Sprintf("%s_%s_%s", PresentationKey, did, id))
	return p.store.Delete(key)
}

func (p *PresentationController) DelRequestPresentation(did, id string)error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestPresentationKey, did, id))
	return p.store.Delete(key)
}