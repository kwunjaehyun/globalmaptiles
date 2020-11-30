package globalmaptiles

import (
	"fmt"
	"math"
)

var RADIUS = 6378137
var TILESIZE = 256
var PI2 = math.Pi * 2
var PI_DIV_360 = math.Pi / 360
var PI_DIV_180 = math.Pi / 180
var PI_DIV_2 = math.Pi / 2
var _180_DIV_PI = 180 / math.Pi

type GlobalMercator struct {
	tileSize int
	initialResolution float64
	originShift float64
}

var defaultGlobalMercator GlobalMercator

func init() {
	defaultGlobalMercator = GlobalMercator{TILESIZE, PI2 * float64(RADIUS) / float64(TILESIZE), PI2 * float64(RADIUS)/2}
}

func GetDefaultGlobalMercator() GlobalMercator{
	return defaultGlobalMercator
}
func (gm *GlobalMercator) LatLonToMeters(latitude float64, longitude float64) (float64,float64) {
	lat := latitude
	lon := longitude

	x := lon * gm.originShift / 180
	y := math.Log(math.Tan((90 + lat) * PI_DIV_360)) / PI_DIV_180

	y = y * gm.originShift / 180
	return x, y
}

func (gm *GlobalMercator) MetersToLatLon(mx float64, my float64) (float64,float64) {
	lat := mx / gm.originShift * 180
	lon := my / gm.originShift * 180

	lat = _180_DIV_PI * (2 * math.Atan(math.Exp(lat * PI_DIV_180)) - PI_DIV_2)

	return lat, lon
}

func (gm *GlobalMercator) MetersToPixel(mx float64, my float64, zoom int) (float64,float64) {
	res := gm.getResolution(zoom);

	px := (mx + gm.originShift) / res;
	py := (my + gm.originShift) / res;

	return px, py
}

func (gm *GlobalMercator) PixelToMeters(px float64, py float64, zoom int) (float64,float64) {
	res := gm.getResolution(zoom);

	x := px * res - gm.originShift
	y := py * res - gm.originShift

	return x, y
}

func (gm *GlobalMercator) PixelToTile(px float64, py float64) (int,int) {
	tx := math.Round(math.Ceil(px / float64(gm.tileSize)) - 1)
	ty := math.Round(math.Ceil(py / float64(gm.tileSize)) - 1)

	return int(tx), int(ty)
}

func (gm *GlobalMercator) MetersToTile(mx float64, my float64, zoom int) (int,int) {
	px, py := gm.MetersToPixel(mx, my, zoom)
	return gm.PixelToTile(px, py)
}

func (gm *GlobalMercator) GoogleTile(tx int, ty int, zoom int) (int, int) {
	return tx, int(math.Pow(2, float64(zoom)) - 1 - float64(ty))
}

func (gm *GlobalMercator) getResolution(zoom int) float64 {
	resolution := gm.initialResolution / math.Pow(2, float64(zoom))
	return resolution
}

func (gm *GlobalMercator) PrintGmMember(){
	fmt.Println(gm.tileSize)
	fmt.Println(gm.initialResolution)
	fmt.Println(gm.originShift)
}
