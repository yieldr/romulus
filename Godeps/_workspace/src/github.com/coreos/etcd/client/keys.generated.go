// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package client

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	"reflect"
	"runtime"
	time "time"
)

const (
	codecSelferC_UTF83457         = 1
	codecSelferC_RAW3457          = 0
	codecSelferValueTypeArray3457 = 10
	codecSelferValueTypeMap3457   = 9
)

var (
	codecSelferBitsize3457                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr3457 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer3457 struct{}

func init() {
	if codec1978.GenVersion != 4 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			4, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 time.Time
		_ = v0
	}
}

func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(3)
			} else {
				var yynn2 int = 3
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
			}
			if yyr2 || yy2arr2 {
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF83457, string(x.Action))
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("action"))
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF83457, string(x.Action))
				}
			}
			if yyr2 || yy2arr2 {
				if x.Node == nil {
					r.EncodeNil()
				} else {
					x.Node.CodecEncodeSelf(e)
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("node"))
				if x.Node == nil {
					r.EncodeNil()
				} else {
					x.Node.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				if x.PrevNode == nil {
					r.EncodeNil()
				} else {
					x.PrevNode.CodecEncodeSelf(e)
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("prevNode"))
				if x.PrevNode == nil {
					r.EncodeNil()
				} else {
					x.PrevNode.CodecEncodeSelf(e)
				}
			}
			if yysep2 {
				r.EncodeEnd()
			}
		}
	}
}

func (x *Response) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym8 := z.DecBinary()
	_ = yym8
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		if r.IsContainerType(codecSelferValueTypeMap3457) {
			yyl9 := r.ReadMapStart()
			if yyl9 == 0 {
				r.ReadEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl9, d)
			}
		} else if r.IsContainerType(codecSelferValueTypeArray3457) {
			yyl9 := r.ReadArrayStart()
			if yyl9 == 0 {
				r.ReadEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl9, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr3457)
		}
	}
}

func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys10Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys10Slc
	var yyhl10 bool = l >= 0
	for yyj10 := 0; ; yyj10++ {
		if yyhl10 {
			if yyj10 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		yys10Slc = r.DecodeBytes(yys10Slc, true, true)
		yys10 := string(yys10Slc)
		switch yys10 {
		case "action":
			if r.TryDecodeAsNil() {
				x.Action = ""
			} else {
				x.Action = string(r.DecodeString())
			}
		case "node":
			if r.TryDecodeAsNil() {
				if x.Node != nil {
					x.Node = nil
				}
			} else {
				if x.Node == nil {
					x.Node = new(Node)
				}
				x.Node.CodecDecodeSelf(d)
			}
		case "prevNode":
			if r.TryDecodeAsNil() {
				if x.PrevNode != nil {
					x.PrevNode = nil
				}
			} else {
				if x.PrevNode == nil {
					x.PrevNode = new(Node)
				}
				x.PrevNode.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys10)
		} // end switch yys10
	} // end for yyj10
	if !yyhl10 {
		r.ReadEnd()
	}
}

func (x *Response) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj14 int
	var yyb14 bool
	var yyhl14 bool = l >= 0
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Action = ""
	} else {
		x.Action = string(r.DecodeString())
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		if x.Node != nil {
			x.Node = nil
		}
	} else {
		if x.Node == nil {
			x.Node = new(Node)
		}
		x.Node.CodecDecodeSelf(d)
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		if x.PrevNode != nil {
			x.PrevNode = nil
		}
	} else {
		if x.PrevNode == nil {
			x.PrevNode = new(Node)
		}
		x.PrevNode.CodecDecodeSelf(d)
	}
	for {
		yyj14++
		if yyhl14 {
			yyb14 = yyj14 > l
		} else {
			yyb14 = r.CheckBreak()
		}
		if yyb14 {
			break
		}
		z.DecStructFieldNotFound(yyj14-1, "")
	}
	r.ReadEnd()
}

func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym18 := z.EncBinary()
		_ = yym18
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep19 := !z.EncBinary()
			yy2arr19 := z.EncBasicHandle().StructToArray
			var yyq19 [8]bool
			_, _, _ = yysep19, yyq19, yy2arr19
			const yyr19 bool = false
			yyq19[1] = x.Dir != false
			yyq19[6] = x.Expiration != nil
			yyq19[7] = x.TTL != 0
			if yyr19 || yy2arr19 {
				r.EncodeArrayStart(8)
			} else {
				var yynn19 int = 5
				for _, b := range yyq19 {
					if b {
						yynn19++
					}
				}
				r.EncodeMapStart(yynn19)
			}
			if yyr19 || yy2arr19 {
				yym21 := z.EncBinary()
				_ = yym21
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF83457, string(x.Key))
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("key"))
				yym22 := z.EncBinary()
				_ = yym22
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF83457, string(x.Key))
				}
			}
			if yyr19 || yy2arr19 {
				if yyq19[1] {
					yym24 := z.EncBinary()
					_ = yym24
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq19[1] {
					r.EncodeString(codecSelferC_UTF83457, string("dir"))
					yym25 := z.EncBinary()
					_ = yym25
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				}
			}
			if yyr19 || yy2arr19 {
				yym27 := z.EncBinary()
				_ = yym27
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF83457, string(x.Value))
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("value"))
				yym28 := z.EncBinary()
				_ = yym28
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF83457, string(x.Value))
				}
			}
			if yyr19 || yy2arr19 {
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("nodes"))
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			}
			if yyr19 || yy2arr19 {
				yym31 := z.EncBinary()
				_ = yym31
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("createdIndex"))
				yym32 := z.EncBinary()
				_ = yym32
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			}
			if yyr19 || yy2arr19 {
				yym34 := z.EncBinary()
				_ = yym34
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			} else {
				r.EncodeString(codecSelferC_UTF83457, string("modifiedIndex"))
				yym35 := z.EncBinary()
				_ = yym35
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			}
			if yyr19 || yy2arr19 {
				if yyq19[6] {
					if x.Expiration == nil {
						r.EncodeNil()
					} else {
						yym37 := z.EncBinary()
						_ = yym37
						if false {
						} else if yym38 := z.TimeRtidIfBinc(); yym38 != 0 {
							r.EncodeBuiltin(yym38, x.Expiration)
						} else if z.HasExtensions() && z.EncExt(x.Expiration) {
						} else if yym37 {
							z.EncBinaryMarshal(x.Expiration)
						} else if !yym37 && z.IsJSONHandle() {
							z.EncJSONMarshal(x.Expiration)
						} else {
							z.EncFallback(x.Expiration)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq19[6] {
					r.EncodeString(codecSelferC_UTF83457, string("expiration"))
					if x.Expiration == nil {
						r.EncodeNil()
					} else {
						yym39 := z.EncBinary()
						_ = yym39
						if false {
						} else if yym40 := z.TimeRtidIfBinc(); yym40 != 0 {
							r.EncodeBuiltin(yym40, x.Expiration)
						} else if z.HasExtensions() && z.EncExt(x.Expiration) {
						} else if yym39 {
							z.EncBinaryMarshal(x.Expiration)
						} else if !yym39 && z.IsJSONHandle() {
							z.EncJSONMarshal(x.Expiration)
						} else {
							z.EncFallback(x.Expiration)
						}
					}
				}
			}
			if yyr19 || yy2arr19 {
				if yyq19[7] {
					yym42 := z.EncBinary()
					_ = yym42
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq19[7] {
					r.EncodeString(codecSelferC_UTF83457, string("ttl"))
					yym43 := z.EncBinary()
					_ = yym43
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				}
			}
			if yysep19 {
				r.EncodeEnd()
			}
		}
	}
}

func (x *Node) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym44 := z.DecBinary()
	_ = yym44
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		if r.IsContainerType(codecSelferValueTypeMap3457) {
			yyl45 := r.ReadMapStart()
			if yyl45 == 0 {
				r.ReadEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl45, d)
			}
		} else if r.IsContainerType(codecSelferValueTypeArray3457) {
			yyl45 := r.ReadArrayStart()
			if yyl45 == 0 {
				r.ReadEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl45, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr3457)
		}
	}
}

func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys46Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys46Slc
	var yyhl46 bool = l >= 0
	for yyj46 := 0; ; yyj46++ {
		if yyhl46 {
			if yyj46 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		yys46Slc = r.DecodeBytes(yys46Slc, true, true)
		yys46 := string(yys46Slc)
		switch yys46 {
		case "key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = string(r.DecodeString())
			}
		case "dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = bool(r.DecodeBool())
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = string(r.DecodeString())
			}
		case "nodes":
			if r.TryDecodeAsNil() {
				x.Nodes = nil
			} else {
				yyv50 := &x.Nodes
				yyv50.CodecDecodeSelf(d)
			}
		case "createdIndex":
			if r.TryDecodeAsNil() {
				x.CreatedIndex = 0
			} else {
				x.CreatedIndex = uint64(r.DecodeUint(64))
			}
		case "modifiedIndex":
			if r.TryDecodeAsNil() {
				x.ModifiedIndex = 0
			} else {
				x.ModifiedIndex = uint64(r.DecodeUint(64))
			}
		case "expiration":
			if r.TryDecodeAsNil() {
				if x.Expiration != nil {
					x.Expiration = nil
				}
			} else {
				if x.Expiration == nil {
					x.Expiration = new(time.Time)
				}
				yym54 := z.DecBinary()
				_ = yym54
				if false {
				} else if yym55 := z.TimeRtidIfBinc(); yym55 != 0 {
					r.DecodeBuiltin(yym55, x.Expiration)
				} else if z.HasExtensions() && z.DecExt(x.Expiration) {
				} else if yym54 {
					z.DecBinaryUnmarshal(x.Expiration)
				} else if !yym54 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(x.Expiration)
				} else {
					z.DecFallback(x.Expiration, false)
				}
			}
		case "ttl":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				x.TTL = int64(r.DecodeInt(64))
			}
		default:
			z.DecStructFieldNotFound(-1, yys46)
		} // end switch yys46
	} // end for yyj46
	if !yyhl46 {
		r.ReadEnd()
	}
}

func (x *Node) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj57 int
	var yyb57 bool
	var yyhl57 bool = l >= 0
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = string(r.DecodeString())
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		x.Dir = bool(r.DecodeBool())
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = string(r.DecodeString())
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.Nodes = nil
	} else {
		yyv61 := &x.Nodes
		yyv61.CodecDecodeSelf(d)
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.CreatedIndex = 0
	} else {
		x.CreatedIndex = uint64(r.DecodeUint(64))
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.ModifiedIndex = 0
	} else {
		x.ModifiedIndex = uint64(r.DecodeUint(64))
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		if x.Expiration != nil {
			x.Expiration = nil
		}
	} else {
		if x.Expiration == nil {
			x.Expiration = new(time.Time)
		}
		yym65 := z.DecBinary()
		_ = yym65
		if false {
		} else if yym66 := z.TimeRtidIfBinc(); yym66 != 0 {
			r.DecodeBuiltin(yym66, x.Expiration)
		} else if z.HasExtensions() && z.DecExt(x.Expiration) {
		} else if yym65 {
			z.DecBinaryUnmarshal(x.Expiration)
		} else if !yym65 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(x.Expiration)
		} else {
			z.DecFallback(x.Expiration, false)
		}
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		r.ReadEnd()
		return
	}
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		x.TTL = int64(r.DecodeInt(64))
	}
	for {
		yyj57++
		if yyhl57 {
			yyb57 = yyj57 > l
		} else {
			yyb57 = r.CheckBreak()
		}
		if yyb57 {
			break
		}
		z.DecStructFieldNotFound(yyj57-1, "")
	}
	r.ReadEnd()
}

func (x Nodes) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym68 := z.EncBinary()
		_ = yym68
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			h.encNodes((Nodes)(x), e)
		}
	}
}

func (x *Nodes) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym69 := z.DecBinary()
	_ = yym69
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		h.decNodes((*Nodes)(x), d)
	}
}

func (x codecSelfer3457) encNodes(v Nodes, e *codec1978.Encoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv70 := range v {
		if yyv70 == nil {
			r.EncodeNil()
		} else {
			yyv70.CodecEncodeSelf(e)
		}
	}
	r.EncodeEnd()
}

func (x codecSelfer3457) decNodes(v *Nodes, d *codec1978.Decoder) {
	var h codecSelfer3457
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv71 := *v
	yyh71, yyl71 := z.DecSliceHelperStart()

	var yyrr71, yyrl71 int
	var yyc71, yyrt71 bool
	_, _, _ = yyc71, yyrt71, yyrl71
	yyrr71 = yyl71

	if yyv71 == nil {
		if yyrl71, yyrt71 = z.DecInferLen(yyl71, z.DecBasicHandle().MaxInitLen, 8); yyrt71 {
			yyrr71 = yyrl71
		}
		yyv71 = make(Nodes, yyrl71)
		yyc71 = true
	}

	if yyl71 == 0 {
		if len(yyv71) != 0 {
			yyv71 = yyv71[:0]
			yyc71 = true
		}
	} else if yyl71 > 0 {

		if yyl71 > cap(yyv71) {
			yyrl71, yyrt71 = z.DecInferLen(yyl71, z.DecBasicHandle().MaxInitLen, 8)
			yyv71 = make([]*Node, yyrl71)
			yyc71 = true

			yyrr71 = len(yyv71)
		} else if yyl71 != len(yyv71) {
			yyv71 = yyv71[:yyl71]
			yyc71 = true
		}
		yyj71 := 0
		for ; yyj71 < yyrr71; yyj71++ {
			if r.TryDecodeAsNil() {
				if yyv71[yyj71] != nil {
					*yyv71[yyj71] = Node{}
				}
			} else {
				if yyv71[yyj71] == nil {
					yyv71[yyj71] = new(Node)
				}
				yyw72 := yyv71[yyj71]
				yyw72.CodecDecodeSelf(d)
			}

		}
		if yyrt71 {
			for ; yyj71 < yyl71; yyj71++ {
				yyv71 = append(yyv71, nil)
				if r.TryDecodeAsNil() {
					if yyv71[yyj71] != nil {
						*yyv71[yyj71] = Node{}
					}
				} else {
					if yyv71[yyj71] == nil {
						yyv71[yyj71] = new(Node)
					}
					yyw73 := yyv71[yyj71]
					yyw73.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		for yyj71 := 0; !r.CheckBreak(); yyj71++ {
			if yyj71 >= len(yyv71) {
				yyv71 = append(yyv71, nil) // var yyz71 *Node
				yyc71 = true
			}

			if yyj71 < len(yyv71) {
				if r.TryDecodeAsNil() {
					if yyv71[yyj71] != nil {
						*yyv71[yyj71] = Node{}
					}
				} else {
					if yyv71[yyj71] == nil {
						yyv71[yyj71] = new(Node)
					}
					yyw74 := yyv71[yyj71]
					yyw74.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		yyh71.End()
	}
	if yyc71 {
		*v = yyv71
	}

}
