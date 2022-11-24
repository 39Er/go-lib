package id

import "testing"

func TestGenerateRandomID(t *testing.T) {
	t.Log(GenerateRandomID("c-"))
}

func TestRandomStr(t *testing.T) {
	t.Log(RandomStr(Letters, 5))
}
