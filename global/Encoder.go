package global

import pkg "URL_Shortener/package"

var Encode *pkg.Encodeing

func init() {
	SetEncode()
}

func SetEncode() {
	Encode = pkg.GetEncodeing()

	Encode.Set_encode()
	Encode.Set_decode()
}
