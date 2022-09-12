package main

import (
	"testing"
)

func TestGis(t *testing.T) {
	// 公主坟
	var p *Gis = &Gis{
		116.31015, 39.907341,
	}

	// 军博
	x := p.distance(&Gis{
		116.321437, 39.907407,
	})

	t.Logf("The distance is %v", x)
	t.Fail()
}
