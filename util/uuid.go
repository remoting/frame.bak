package util

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"time"
)

var rander = rand.Reader

func New() string {
	b58 := NewBitcoinBase58()
	i := time.Now().Unix()
	b := Int64ToBytes(i)
	ss, _ := b58.EncodeToString(b)
	var uuid [16]byte
	io.ReadFull(rander, uuid[:])
	ss1, _ := b58.EncodeToString(uuid[:])
	return ss + "" + ss1
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
