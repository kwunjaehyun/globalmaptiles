package globalmaptiles


type GlobalGeodetic struct {
	tileSize int
}

var globalGeodetic GlobalGeodetic

func init() {
	globalGeodetic = GlobalGeodetic{TILESIZE}
}

func getGlobalGeodetic() GlobalGeodetic{
	return globalGeodetic
}