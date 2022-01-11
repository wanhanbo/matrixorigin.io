// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unary

import (
	"fmt"

	"github.com/matrixorigin/matrixone/pkg/builtin"
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/extend"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/extend/overload"
	"github.com/matrixorigin/matrixone/pkg/vectorize/length"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func init() {
	extend.FunctionRegistry["length"] = builtin.Length
	extend.UnaryReturnTypes[builtin.Length] = func(_ extend.Extend) types.T {
		return types.T_int64
	}
	extend.UnaryStrings[builtin.Length] = func(e extend.Extend) string {
		return fmt.Sprintf("length(%s)", e)
	}
	overload.OpTypes[builtin.Length] = overload.Unary
	overload.UnaryOps[builtin.Length] = []*overload.UnaryOp{
		{
			Typ:        types.T_varchar,
			ReturnType: types.T_int64,
			Fn: func(lv *vector.Vector, proc *process.Process, _ bool) (*vector.Vector, error) {
				lvs := lv.Col.(*types.Bytes)
				vec, err := process.Get(proc, 8*int64(len(lvs.Lengths)), types.Type{Oid: types.T_int64, Size: 8})
				if err != nil {
					return nil, err
				}
				rs := encoding.DecodeInt64Slice(vec.Data)
				rs = rs[:len(lvs.Lengths)]
				vec.Col = rs
				nulls.Set(vec.Nsp, lv.Nsp)
				vector.SetCol(vec, length.StrLength(lvs, rs))
				return vec, nil
			},
		},
		{
			Typ:        types.T_char,
			ReturnType: types.T_int64,
			Fn: func(lv *vector.Vector, proc *process.Process, _ bool) (*vector.Vector, error) {
				lvs := lv.Col.(*types.Bytes)
				vec, err := process.Get(proc, 8*int64(len(lvs.Lengths)), types.Type{Oid: types.T_int64, Size: 8})
				if err != nil {
					return nil, err
				}
				rs := encoding.DecodeInt64Slice(vec.Data)
				rs = rs[:len(lvs.Lengths)]
				vec.Col = rs
				nulls.Set(vec.Nsp, lv.Nsp)
				vector.SetCol(vec, length.StrLength(lvs, rs))
				return vec, nil
			},
		},
	}
}
