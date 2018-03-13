package geo

import "testing"

func TestIosAlpha2to3(t *testing.T) {
	t.Log(IosAlpha2to3("HK"))
}

func TestIosAlpha3to2(t *testing.T) {
	t.Log(IosAlpha3to2("DZA"))
}