package controller

import (
	"bytes"
	"encoding/json"
	"merdeka/model"
	"merdeka/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPersonControllerGet(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	ps := service.NewMockPersonService()
	ps.Add(model.Person{Name: "Budi", Address: "Bandung"})
	ps.Add(model.Person{Name: "Chandra", Address: "Cilegon"})

	pc := NewPersonController(ps)
	pc.Get(c)

	var persons []model.Person
	if err := json.Unmarshal(rec.Body.Bytes(), &persons); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	if len(persons) != 2 {
		t.Errorf("Expecting len(persons) to be 2, get %d", len(persons))
	}
	if persons[0].Name != "Budi" {
		t.Errorf("Expecting persons[0].Name to be Budi, get %s", persons[0].Name)
	}
	if persons[1].Name != "Chandra" {
		t.Errorf("Expecting persons[0].Name to be Chandra, get %s", persons[0].Name)
	}
}

func TestPersonControllerAdd(t *testing.T) {
	e := echo.New()

	newPersonJson, _ := json.Marshal(map[string]string{
		"name":    "Anton",
		"address": "Aceh",
	})
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newPersonJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	ps := service.NewMockPersonService()
	pc := NewPersonController(ps)
	pc.Add(c)

	persons, err := ps.Get()
	if err != nil {
		t.Error(err)
	}
	if len(persons) != 1 {
		t.Errorf("Expecting len(persons) to be 1, get %d", len(persons))
	}
	if persons[0].Name != "Anton" {
		t.Errorf("Expecting persons[0].Name to be Anton, get %s", persons[0].Name)
	}
	if persons[0].Address != "Aceh" {
		t.Errorf("Expecting persons[0].Address to be Aceh, get %s", persons[0].Name)
	}

	// TODO: also check response
}
