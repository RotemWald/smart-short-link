package repositories

import (
	"testing"

	"github.com/RotemWald/smart-short-link/entities"
)

func TestSimpleSetAndGet(t *testing.T) {
	repository := NewMemory()
	urls := []*entities.SmartUrl{
		{
			StartHour: 0,
			EndHour:   12,
			Url:       "http://www.ynet.co.il",
		},
		{
			StartHour: 12,
			EndHour:   23,
			Url:       "http://www.ynet-1234.co.il",
		},
	}

	if err := repository.SetUrls("a1", urls); err != nil {
		t.Fatal(err)
	}

	url, err := repository.GetUrl("a1", 10)
	if err != nil {
		t.Fatal(err)
	}
	if url.Url != "http://www.ynet.co.il" {
		t.Fatal("got wrong url")
	}

	url, err = repository.GetUrl("a1", 12)
	if err != nil {
		t.Fatal(err)
	}
	if url.Url != "http://www.ynet-1234.co.il" {
		t.Fatal("got wrong url")
	}
}

func TestSetAndRefreshAndThenGet(t *testing.T) {
	repository := NewMemory()
	urls := []*entities.SmartUrl{
		{
			StartHour: 0,
			EndHour:   12,
			Url:       "http://www.ynet.co.il",
		},
		{
			StartHour: 12,
			EndHour:   23,
			Url:       "http://www.ynet-1234.co.il",
		},
	}

	if err := repository.SetUrls("a1", urls); err != nil {
		t.Fatal(err)
	}

	if err := repository.RefreshUrls("a1"); err != nil {
		t.Fatal(err)
	}

	url, err := repository.GetUrl("a1", 10)
	if err != nil {
		t.Fatal(err)
	}
	if url.Url != "http://www.ynet.co.il" {
		t.Fatal("got wrong url")
	}

	url, err = repository.GetUrl("a1", 12)
	if err == nil {
		t.Fatal("should be nil")
	}
}
