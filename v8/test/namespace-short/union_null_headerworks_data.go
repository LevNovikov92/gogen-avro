// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
)

type UnionNullHeaderworksDataTypeEnum int

const (
	UnionNullHeaderworksDataTypeEnumHeaderworksData UnionNullHeaderworksDataTypeEnum = 1
)

type UnionNullHeaderworksData struct {
	Null            *types.NullVal
	HeaderworksData *HeaderworksData
	UnionType       UnionNullHeaderworksDataTypeEnum
}

func writeUnionNullHeaderworksData(r *UnionNullHeaderworksData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullHeaderworksDataTypeEnumHeaderworksData:
		return writeHeaderworksData(r.HeaderworksData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksData")
}

func NewUnionNullHeaderworksData() *UnionNullHeaderworksData {
	return &UnionNullHeaderworksData{}
}

func (r *UnionNullHeaderworksData) Serialize(w io.Writer) error {
	return writeUnionNullHeaderworksData(r, w)
}

func DeserializeUnionNullHeaderworksData(r io.Reader) (*UnionNullHeaderworksData, error) {
	t := NewUnionNullHeaderworksData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeUnionNullHeaderworksDataFromSchema(r io.Reader, schema string) (*UnionNullHeaderworksData, error) {
	t := NewUnionNullHeaderworksData()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func (r *UnionNullHeaderworksData) Schema() string {
	return "[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"headerworks\",\"type\":\"record\"}]"
}

func (_ *UnionNullHeaderworksData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullHeaderworksData) SetLong(v int64) {
	r.UnionType = (UnionNullHeaderworksDataTypeEnum)(v)
}
func (r *UnionNullHeaderworksData) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.HeaderworksData = NewHeaderworksData()
		return r.HeaderworksData
	}
	panic("Unknown field index")
}
func (_ *UnionNullHeaderworksData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) Finalize()                        {}

func (r *UnionNullHeaderworksData) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullHeaderworksDataTypeEnumHeaderworksData:
		return json.Marshal(map[string]interface{}{"headerworks.Data": r.HeaderworksData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullHeaderworksData")
}

func (r *UnionNullHeaderworksData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["headerworks.Data"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.HeaderworksData)
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksData")
}