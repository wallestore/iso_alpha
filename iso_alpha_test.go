package iso

import "testing"

func TestIosAlpha2to3(t *testing.T) {
	t.Log(IsoAlpha2to3("hk"))
}

func TestIosAlpha3to2(t *testing.T) {
	t.Log(IsoAlpha3to2("DZA"))
}