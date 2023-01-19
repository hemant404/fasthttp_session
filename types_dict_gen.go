package session

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Dict) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "KV":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "KV")
				return
			}
			if z.KV == nil {
				z.KV = make(map[string]interface{}, zb0002)
			} else if len(z.KV) > 0 {
				for key := range z.KV {
					delete(z.KV, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 interface{}
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "KV")
					return
				}
				za0002, err = dc.ReadIntf()
				if err != nil {
					err = msgp.WrapError(err, "KV", za0001)
					return
				}
				z.KV[za0001] = za0002
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Dict) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "KV"
	err = en.Append(0x81, 0xa2, 0x4b, 0x56)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.KV)))
	if err != nil {
		err = msgp.WrapError(err, "KV")
		return
	}
	for za0001, za0002 := range z.KV {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "KV")
			return
		}
		err = en.WriteIntf(za0002)
		if err != nil {
			err = msgp.WrapError(err, "KV", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Dict) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "KV"
	o = append(o, 0x81, 0xa2, 0x4b, 0x56)
	o = msgp.AppendMapHeader(o, uint32(len(z.KV)))
	for za0001, za0002 := range z.KV {
		o = msgp.AppendString(o, za0001)
		o, err = msgp.AppendIntf(o, za0002)
		if err != nil {
			err = msgp.WrapError(err, "KV", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Dict) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "KV":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "KV")
				return
			}
			if z.KV == nil {
				z.KV = make(map[string]interface{}, zb0002)
			} else if len(z.KV) > 0 {
				for key := range z.KV {
					delete(z.KV, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 interface{}
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "KV")
					return
				}
				za0002, bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "KV", za0001)
					return
				}
				z.KV[za0001] = za0002
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Dict) Msgsize() (s int) {
	s = 1 + 3 + msgp.MapHeaderSize
	if z.KV != nil {
		for za0001, za0002 := range z.KV {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.GuessSize(za0002)
		}
	}
	return
}