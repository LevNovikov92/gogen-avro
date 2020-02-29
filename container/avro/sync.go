// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     block.avsc
 *     header.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
)

func writeSync(r Sync, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type Sync SyncWrapper
type SyncWrapper [16]byte

func (_ *SyncWrapper) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *SyncWrapper) SetInt(v int32) { panic("Unsupported operation") }
func (_ *SyncWrapper) SetLong(v int64) { panic("Unsupported operation") }
func (_ *SyncWrapper) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *SyncWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (r *SyncWrapper) SetBytes(v []byte) { 
	copy((*r)[:], v)
}
func (_ *SyncWrapper) SetString(v string) { panic("Unsupported operation") }
func (_ *SyncWrapper) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *SyncWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *SyncWrapper) Clear(i int) { panic("Unsupported operation") }
func (_ *SyncWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *SyncWrapper) ClearMap(key string) { panic("Unsupported operation") }
func (_ *SyncWrapper) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *SyncWrapper) Finalize() { }
func (_ *SyncWrapper) SetDefault(i int) { panic("Unsupported operation") }

