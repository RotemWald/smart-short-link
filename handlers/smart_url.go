package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RotemWald/smart-short-link/entities"
	"github.com/RotemWald/smart-short-link/services"
)

const (
	uuidMethod    = "uuid"
	counterMethod = "counter"
)

type SmartUrl struct {
	service *services.SmartUrl
}

func NewSmartUrl(service *services.SmartUrl) *SmartUrl {
	return &SmartUrl{
		service: service,
	}
}

func (h *SmartUrl) CreateByUUID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.create(uuidMethod, w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
}

func (h *SmartUrl) CreateByCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.create(counterMethod, w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
}

func (h *SmartUrl) create(method string, w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var urls []*entities.SmartUrl
	err = json.Unmarshal(bytes, &urls)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var key string
	switch method {
	case uuidMethod:
		key, err = h.service.SetUrlsByUUID(urls)
		break
	case counterMethod:
		key, err = h.service.SetUrlsByCounter(urls)
		break
	default:
		err = fmt.Errorf("only uuid or counter methods are supported")
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	url := entities.Url{
		Url: fmt.Sprintf("http://%s/%s/%s", r.Host, method, key),
	}
	bytes, err = json.Marshal(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
