package services

import (
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/RotemWald/smart-short-link/repositories"
)

func TestSimpleGetUrl(t *testing.T) {
	repository := repositories.NewDummy()
	service := NewSmartUrl(repository)
	url, err := service.GetUrl("a1")
	if err != nil {
		t.Fatal(err)
	}
	if url != "http://www.ynet.co.il" {
		t.Fatal("got wrong url")
	}
}

func TestMassiveGetUrl(t *testing.T) {
	repository := repositories.NewDummy()
	service := NewSmartUrl(repository)
	for i := 0; i < 100; i++ {
		_, _ = service.GetUrl("a1")
		if runtime.NumGoroutine() > 2 {
			t.Fatal("too many calls to refresh urls mechanism")
		}
	}
}

func TestSetUrlsByUuid(t *testing.T) {
	repository := repositories.NewDummy()
	service := NewSmartUrl(repository)
	uuid, err := service.SetUrlsByUUID(nil)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(uuid, "-") != 4 {
		t.Fatal("not uuid")
	}
}

func TestSetUrlsByCounter(t *testing.T) {
	repository := repositories.NewDummy()
	service := NewSmartUrl(repository)

	var wg sync.WaitGroup
	counters := make(chan int, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(c chan<- int) {
			defer wg.Done()
			counter, _ := service.SetUrlsByCounter(nil)
			num, _ := strconv.Atoi(counter[1:])
			c <- num
		}(counters)
	}
	wg.Wait()

	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = <-counters
	}
	sort.Ints(arr)
	if arr[99] != 100 {
		// if the number 100 does not exist in the slice,
		// it means the counter were not incremented atomically,
		// so there are at least two counters with same value
		t.Fatal("wrong counters generated")
	}
}
