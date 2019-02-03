package config

import(
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
)

// funcion para encrypt
func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

var (
	t = time.Now().UTC().Unix()
	// timestamp
	ts = strconv.FormatInt(t,10)
	// llave publica del API
	publicKey = "a71fa6c845e117349196fe9c4a58b2ae"
	// llave privada del API
	privateKey = "7c55664c217553d96ed6e84098a931dff0204776"
	// hash md5
	hash = GetMD5Hash(ts+privateKey+publicKey)
	// url 
	URL = "http://gateway.marvel.com/v1/public/characters?ts="+ts+"&apikey="+publicKey+"&hash="+hash
)