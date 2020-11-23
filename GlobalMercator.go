package globalmaptiles

import (
	"fmt"
	"math"
)

var RADIUS = 6378137
var TILESIZE = 256
var PI2 = math.Pi * 2

type globalMercator struct {
	tileSize int
	initialResolution float64
	originShift float64
}

var GlobalMercator globalMercator

func init() {
	GlobalMercator = globalMercator{TILESIZE, PI2 * float64(RADIUS) / float64(TILESIZE), PI2 * float64(RADIUS)/2}
}


func (gm globalMercator) Public (){
	fmt.Println(gm.tileSize)
	fmt.Println(gm.initialResolution)
	fmt.Println(gm.originShift)
}
