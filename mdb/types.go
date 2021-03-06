// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 05 Sep 2016 13:16:14 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package mdb

/*
#cgo linux CFLAGS: -DMDB_USE_SYSV_SEM=1
#cgo windows LDFLAGS: -lntdll
#include "lmdb.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Mode type as declared in mdb/lmdb.h:179
type Mode uint16

// Size type as declared in mdb/lmdb.h:186
type Size uint

// Env as declared in mdb/lmdb.h:239
type Env C.MDB_env

// Txn as declared in mdb/lmdb.h:246
type Txn C.MDB_txn

// Dbi type as declared in mdb/lmdb.h:249
type Dbi uint32

// Cursor as declared in mdb/lmdb.h:252
type Cursor C.MDB_cursor

// Val as declared in mdb/lmdb.h:268
type Val struct {
	Size           uint
	Data           unsafe.Pointer
	reff5baf417    *C.MDB_val
	allocsf5baf417 interface{}
}

// CmpFunc type as declared in mdb/lmdb.h:271
type CmpFunc func(a *Val, b *Val) int32

// RelFunc type as declared in mdb/lmdb.h:287
type RelFunc func(item *Val, oldptr unsafe.Pointer, newptr unsafe.Pointer, relctx unsafe.Pointer)

// Stats as declared in mdb/lmdb.h:472
type Stats struct {
	Psize          uint32
	Depth          uint32
	BranchPages    Size
	LeafPages      Size
	OverflowPages  Size
	Entries        Size
	ref9ed18a7f    *C.MDB_stat
	allocs9ed18a7f interface{}
}

// Envinfo as declared in mdb/lmdb.h:482
type Envinfo struct {
	Mapaddr        unsafe.Pointer
	Mapsize        Size
	LastPgno       Size
	LastTxnid      Size
	Maxreaders     uint32
	Numreaders     uint32
	refea477e79    *C.MDB_envinfo
	allocsea477e79 interface{}
}

// AssertFunc type as declared in mdb/lmdb.h:931
type AssertFunc func(env *Env, msg string)

// MsgFunc type as declared in mdb/lmdb.h:1579
type MsgFunc func(msg string, ctx unsafe.Pointer) int32
