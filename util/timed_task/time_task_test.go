package timed_task

import (
	"recommend/golbal"
	"strings"
	"testing"
)

func TestGetBannerSection(t *testing.T) {
	_, err := golbal.LoadConfig("../dev.json")
	if err != nil {
		t.Fatal(err)
	}
	data := addBannerData(1)
	if data != nil {
		t.Log(data)
	}
}

func TestGettCartoonTypes(t *testing.T) {
	t.Log(ItemGetKey())
}

func ItemGetKey(keys ...string) string {
	return strings.Join(keys, "_")
}
