package zabbix

import (
	"encoding/binary"
	"encoding/json"
	"time"
)

type Packet struct {
	Request  string    `json:"request"`
	Data     []*Metric `json:"data"`
	Clock    int64     `json:"clock"`
	jsonData []byte    // cached json data
	dataLen  []byte    // cached length data
}

func NewPacket(data []*Metric, clock ...int64) *Packet {
	p := &Packet{Request: `sender data`, Data: data}
	if len(clock) > 0 {
		p.Clock = clock[0]
	} else {
		p.Clock = time.Now().Unix()
	}
	return p
}

// cache json data
func (p *Packet) Json() []byte {
	if len(p.jsonData) != 0 {
		return p.jsonData
	}
	jsonData, _ := json.Marshal(p)
	p.jsonData = jsonData
	return p.jsonData
}

// cached length data
func (p *Packet) DataLen() []byte {
	if len(p.dataLen) > 0 {
		return p.dataLen
	}
	dataLen := make([]byte, 8)
	binary.LittleEndian.PutUint32(dataLen, uint32(len(p.Json())))
	p.dataLen = dataLen
	return p.dataLen
}
