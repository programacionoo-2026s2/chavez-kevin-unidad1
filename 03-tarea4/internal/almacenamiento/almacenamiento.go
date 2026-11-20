package almacenamiento

import "math"

const (
	B  = 1.0
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func Bytes(unidad string) (float64, bool) {
	switch unidad {
	case "B":
		return B, true
	case "KB":
		return KB, true
	case "MB":
		return MB, true
	case "GB":
		return GB, true
	default:
		return 0, false
	}
}

func Convertir(valor, bytesOrigen, bytesDestino float64) float64 {
	return redondear(valor * bytesOrigen / bytesDestino)
}

func redondear(v float64) float64 {
	return math.Round(v*100) / 100
}
