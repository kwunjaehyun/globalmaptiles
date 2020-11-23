package globalmaptiles


type GlobalGeodetic struct {
	tileSize int
}

var globalGeodetic GlobalGeodetic

func init() {
	globalGeodetic = GlobalGeodetic{TILESIZE}
}

func GetGlobalGeodetic() GlobalGeodetic{
	return globalGeodetic
}