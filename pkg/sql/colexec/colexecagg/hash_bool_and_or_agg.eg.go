// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecagg

import (
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
)

// Remove unused warning.
var _ = colexecerror.InternalError

func newBoolAndHashAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolAndHashAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolAndHashAgg struct {
	hashAggregateFuncBase
	col        []bool
	sawNonNull bool
	curAgg     bool
}

var _ AggregateFunc = &boolAndHashAgg{}

func (a *boolAndHashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.col = vec.Bool()
}

func (a *boolAndHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Bool(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		{
			sel = sel[:inputLen]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						a.curAgg = a.curAgg && col[i]
						a.sawNonNull = true
					}

				}
			} else {
				for _, i := range sel {

					var isNull bool
					isNull = false
					if !isNull {
						a.curAgg = a.curAgg && col[i]
						a.sawNonNull = true
					}

				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsage(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *boolAndHashAgg) Flush(outputIdx int) {
	if !a.sawNonNull {
		a.nulls.SetNull(outputIdx)
	} else {
		a.col[outputIdx] = a.curAgg
	}
}

func (a *boolAndHashAgg) Reset() {
	a.curAgg = true
}

type boolAndHashAggAlloc struct {
	aggAllocBase
	aggFuncs []boolAndHashAgg
}

var _ aggregateFuncAlloc = &boolAndHashAggAlloc{}

const sizeOfBoolAndHashAgg = int64(unsafe.Sizeof(boolAndHashAgg{}))
const boolAndHashAggSliceOverhead = int64(unsafe.Sizeof([]boolAndHashAgg{}))

func (a *boolAndHashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(boolAndHashAggSliceOverhead + sizeOfBoolAndHashAgg*a.allocSize)
		a.aggFuncs = make([]boolAndHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	f.Reset()
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

func newBoolOrHashAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolOrHashAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolOrHashAgg struct {
	hashAggregateFuncBase
	col        []bool
	sawNonNull bool
	curAgg     bool
}

var _ AggregateFunc = &boolOrHashAgg{}

func (a *boolOrHashAgg) SetOutput(vec coldata.Vec) {
	a.hashAggregateFuncBase.SetOutput(vec)
	a.col = vec.Bool()
}

func (a *boolOrHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Bool(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		{
			sel = sel[:inputLen]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						a.curAgg = a.curAgg || col[i]
						a.sawNonNull = true
					}

				}
			} else {
				for _, i := range sel {

					var isNull bool
					isNull = false
					if !isNull {
						a.curAgg = a.curAgg || col[i]
						a.sawNonNull = true
					}

				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsage(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *boolOrHashAgg) Flush(outputIdx int) {
	if !a.sawNonNull {
		a.nulls.SetNull(outputIdx)
	} else {
		a.col[outputIdx] = a.curAgg
	}
}

func (a *boolOrHashAgg) Reset() {
	a.curAgg = false
}

type boolOrHashAggAlloc struct {
	aggAllocBase
	aggFuncs []boolOrHashAgg
}

var _ aggregateFuncAlloc = &boolOrHashAggAlloc{}

const sizeOfBoolOrHashAgg = int64(unsafe.Sizeof(boolOrHashAgg{}))
const boolOrHashAggSliceOverhead = int64(unsafe.Sizeof([]boolOrHashAgg{}))

func (a *boolOrHashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(boolOrHashAggSliceOverhead + sizeOfBoolOrHashAgg*a.allocSize)
		a.aggFuncs = make([]boolOrHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	f.Reset()
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
