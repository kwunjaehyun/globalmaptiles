package globalmaptiles

import "math"

type GlobalGeodetic struct {
	tileSize int
}

var TILESIZE = 256
var defaultGlobalGeodetic GlobalGeodetic

func (gd *GlobalGeodetic) LatLonToPixel(latitude float64, longitude float64, zoom int) (float64, float64) {
	resolution := gd.getResolution(zoom)
	px := (180 + latitude) / resolution
	py := (90 + longitude) / resolution

	return px, py
}

func (gd *GlobalGeodetic) pixelsToTile(px float64, py float64) (int, int) {
	tx := math.Round(math.Ceil(px/float64(gd.tileSize)) - 1)
	ty := math.Round(math.Ceil(py/float64(gd.tileSize)) - 1)

	return int(tx), int(ty)
}

func (gd *GlobalGeodetic) LatLonToTile(latitude float64, longitude float64, zoom int) (int, int) {
	px, py := gd.LatLonToPixel(latitude, longitude, zoom)
	return gd.pixelsToTile(px, py)
}

func (gd *GlobalGeodetic) getResolution(zoom int) float64 {
	return float64(180) / float64(gd.tileSize) / math.Pow(2, float64(zoom))
}

func (gd *GlobalGeodetic) getTileBound(tx int, ty int, zoom int) (float64, float64, float64, float64) {
	resolution := gd.getResolution(zoom)
	return float64(tx)*float64(gd.tileSize)*resolution - 180,
		float64(ty)*float64(gd.tileSize)*resolution - 90,
		float64(tx+1)*float64(gd.tileSize)*resolution - 180,
		float64(ty+1)*float64(gd.tileSize)*resolution - 90
}

func init() {
	defaultGlobalGeodetic = GlobalGeodetic{TILESIZE}
}

func GetDefaultGlobalGeodetic() GlobalGeodetic {
	return defaultGlobalGeodetic
}
