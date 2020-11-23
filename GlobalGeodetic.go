package globalmaptiles


type globalGeodetic struct {
	tileSize int
}

var GlobalGeodetic globalGeodetic

func init() {
	GlobalGeodetic = globalGeodetic{TILESIZE}
}