// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 06 Jun 2016 15:47:34 MSK.
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

// Version function as declared in mdb/lmdb.h:491
func Version(major *int32, minor *int32, patch *int32) []byte {
	cmajor, _ := (*C.int)(unsafe.Pointer(major)), cgoAllocsUnknown
	cminor, _ := (*C.int)(unsafe.Pointer(minor)), cgoAllocsUnknown
	cpatch, _ := (*C.int)(unsafe.Pointer(patch)), cgoAllocsUnknown
	__ret := C.mdb_version(cmajor, cminor, cpatch)
	__v := (*(*[0x7fffffff]byte)(unsafe.Pointer(__ret)))[:0]
	return __v
}

// StrError function as declared in mdb/lmdb.h:503
func StrError(err int32) []byte {
	cerr, _ := (C.int)(err), cgoAllocsUnknown
	__ret := C.mdb_strerror(cerr)
	__v := (*(*[0x7fffffff]byte)(unsafe.Pointer(__ret)))[:0]
	return __v
}

// EnvCreate function as declared in mdb/lmdb.h:516
func EnvCreate(env **Env) int32 {
	cenv, _ := (**C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	__ret := C.mdb_env_create(cenv)
	__v := (int32)(__ret)
	return __v
}

// EnvOpen function as declared in mdb/lmdb.h:639
func EnvOpen(env *Env, path string, flags uint32, mode Mode) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cpath, _ := unpackPCharString(path)
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cmode, _ := (C.mdb_mode_t)(mode), cgoAllocsUnknown
	__ret := C.mdb_env_open(cenv, cpath, cflags, cmode)
	__v := (int32)(__ret)
	return __v
}

// EnvCopy function as declared in mdb/lmdb.h:655
func EnvCopy(env *Env, path string) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cpath, _ := unpackPCharString(path)
	__ret := C.mdb_env_copy(cenv, cpath)
	__v := (int32)(__ret)
	return __v
}

// EnvCopyfd function as declared in mdb/lmdb.h:670
func EnvCopyfd(env *Env, fd Filehandle) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cfd, _ := (C.mdb_filehandle_t)(fd), cgoAllocsUnknown
	__ret := C.mdb_env_copyfd(cenv, cfd)
	__v := (int32)(__ret)
	return __v
}

// EnvCopy2 function as declared in mdb/lmdb.h:694
func EnvCopy2(env *Env, path string, flags uint32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cpath, _ := unpackPCharString(path)
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	__ret := C.mdb_env_copy2(cenv, cpath, cflags)
	__v := (int32)(__ret)
	return __v
}

// EnvCopyfd2 function as declared in mdb/lmdb.h:713
func EnvCopyfd2(env *Env, fd Filehandle, flags uint32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cfd, _ := (C.mdb_filehandle_t)(fd), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	__ret := C.mdb_env_copyfd2(cenv, cfd, cflags)
	__v := (int32)(__ret)
	return __v
}

// EnvStat function as declared in mdb/lmdb.h:721
func EnvStat(env *Env, stat *Stats) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstat, _ := stat.PassRef()
	__ret := C.mdb_env_stat(cenv, cstat)
	__v := (int32)(__ret)
	return __v
}

// EnvInfo function as declared in mdb/lmdb.h:729
func EnvInfo(env *Env, stat *Envinfo) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstat, _ := stat.PassRef()
	__ret := C.mdb_env_info(cenv, cstat)
	__v := (int32)(__ret)
	return __v
}

// EnvSync function as declared in mdb/lmdb.h:750
func EnvSync(env *Env, force int32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cforce, _ := (C.int)(force), cgoAllocsUnknown
	__ret := C.mdb_env_sync(cenv, cforce)
	__v := (int32)(__ret)
	return __v
}

// EnvClose function as declared in mdb/lmdb.h:760
func EnvClose(env *Env) {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	C.mdb_env_close(cenv)
}

// EnvSetFlags function as declared in mdb/lmdb.h:776
func EnvSetFlags(env *Env, flags uint32, onoff int32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	conoff, _ := (C.int)(onoff), cgoAllocsUnknown
	__ret := C.mdb_env_set_flags(cenv, cflags, conoff)
	__v := (int32)(__ret)
	return __v
}

// EnvGetFlags function as declared in mdb/lmdb.h:788
func EnvGetFlags(env *Env, flags *uint32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cflags, _ := (*C.uint)(unsafe.Pointer(flags)), cgoAllocsUnknown
	__ret := C.mdb_env_get_flags(cenv, cflags)
	__v := (int32)(__ret)
	return __v
}

// EnvGetPath function as declared in mdb/lmdb.h:802
func EnvGetPath(env *Env, path []string) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cpath, _ := unpackArgSString(path)
	__ret := C.mdb_env_get_path(cenv, cpath)
	packSString(path, cpath)
	__v := (int32)(__ret)
	return __v
}

// EnvGetFd function as declared in mdb/lmdb.h:814
func EnvGetFd(env *Env, fd *Filehandle) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cfd, _ := (*C.mdb_filehandle_t)(unsafe.Pointer(fd)), cgoAllocsUnknown
	__ret := C.mdb_env_get_fd(cenv, cfd)
	__v := (int32)(__ret)
	return __v
}

// EnvSetMapsize function as declared in mdb/lmdb.h:848
func EnvSetMapsize(env *Env, size Size) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	csize, _ := (C.mdb_size_t)(size), cgoAllocsUnknown
	__ret := C.mdb_env_set_mapsize(cenv, csize)
	__v := (int32)(__ret)
	return __v
}

// EnvSetMaxreaders function as declared in mdb/lmdb.h:867
func EnvSetMaxreaders(env *Env, readers uint32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	creaders, _ := (C.uint)(readers), cgoAllocsUnknown
	__ret := C.mdb_env_set_maxreaders(cenv, creaders)
	__v := (int32)(__ret)
	return __v
}

// EnvGetMaxreaders function as declared in mdb/lmdb.h:879
func EnvGetMaxreaders(env *Env, readers *uint32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	creaders, _ := (*C.uint)(unsafe.Pointer(readers)), cgoAllocsUnknown
	__ret := C.mdb_env_get_maxreaders(cenv, creaders)
	__v := (int32)(__ret)
	return __v
}

// EnvSetMaxdbs function as declared in mdb/lmdb.h:899
func EnvSetMaxdbs(env *Env, dbs Dbi) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cdbs, _ := (C.MDB_dbi)(dbs), cgoAllocsUnknown
	__ret := C.mdb_env_set_maxdbs(cenv, cdbs)
	__v := (int32)(__ret)
	return __v
}

// EnvGetMaxkeysize function as declared in mdb/lmdb.h:908
func EnvGetMaxkeysize(env *Env) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	__ret := C.mdb_env_get_maxkeysize(cenv)
	__v := (int32)(__ret)
	return __v
}

// EnvSetUserctx function as declared in mdb/lmdb.h:916
func EnvSetUserctx(env *Env, ctx unsafe.Pointer) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cctx, _ := (unsafe.Pointer)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.mdb_env_set_userctx(cenv, cctx)
	__v := (int32)(__ret)
	return __v
}

// EnvGetUserctx function as declared in mdb/lmdb.h:923
func EnvGetUserctx(env *Env) unsafe.Pointer {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	__ret := C.mdb_env_get_userctx(cenv)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// EnvSetAssert function as declared in mdb/lmdb.h:940
func EnvSetAssert(env *Env, _func AssertFunc) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	c_func, _ := _func.PassRef()
	__ret := C.mdb_env_set_assert(cenv, c_func)
	__v := (int32)(__ret)
	return __v
}

// TxnBegin function as declared in mdb/lmdb.h:980
func TxnBegin(env *Env, parent *Txn, flags uint32, txn **Txn) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cparent, _ := (*C.MDB_txn)(unsafe.Pointer(parent)), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	ctxn, _ := (**C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	__ret := C.mdb_txn_begin(cenv, cparent, cflags, ctxn)
	__v := (int32)(__ret)
	return __v
}

// TxnEnv function as declared in mdb/lmdb.h:986
func TxnEnv(txn *Txn) *Env {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	__ret := C.mdb_txn_env(ctxn)
	__v := *(**Env)(unsafe.Pointer(&__ret))
	return __v
}

// TxnId function as declared in mdb/lmdb.h:997
func TxnId(txn *Txn) Size {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	__ret := C.mdb_txn_id(ctxn)
	__v := (Size)(__ret)
	return __v
}

// TxnCommit function as declared in mdb/lmdb.h:1015
func TxnCommit(txn *Txn) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	__ret := C.mdb_txn_commit(ctxn)
	__v := (int32)(__ret)
	return __v
}

// TxnAbort function as declared in mdb/lmdb.h:1025
func TxnAbort(txn *Txn) {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	C.mdb_txn_abort(ctxn)
}

// TxnReset function as declared in mdb/lmdb.h:1044
func TxnReset(txn *Txn) {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	C.mdb_txn_reset(ctxn)
}

// TxnRenew function as declared in mdb/lmdb.h:1060
func TxnRenew(txn *Txn) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	__ret := C.mdb_txn_renew(ctxn)
	__v := (int32)(__ret)
	return __v
}

// DbiOpen function as declared in mdb/lmdb.h:1134
func DbiOpen(txn *Txn, name string, flags uint32, dbi *Dbi) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cdbi, _ := (*C.MDB_dbi)(unsafe.Pointer(dbi)), cgoAllocsUnknown
	__ret := C.mdb_dbi_open(ctxn, cname, cflags, cdbi)
	__v := (int32)(__ret)
	return __v
}

// Stat function as declared in mdb/lmdb.h:1148
func Stat(txn *Txn, dbi Dbi, stat *Stats) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	cstat, _ := stat.PassRef()
	__ret := C.mdb_stat(ctxn, cdbi, cstat)
	__v := (int32)(__ret)
	return __v
}

// DbiFlags function as declared in mdb/lmdb.h:1157
func DbiFlags(txn *Txn, dbi Dbi, flags *uint32) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	cflags, _ := (*C.uint)(unsafe.Pointer(flags)), cgoAllocsUnknown
	__ret := C.mdb_dbi_flags(ctxn, cdbi, cflags)
	__v := (int32)(__ret)
	return __v
}

// DbiClose function as declared in mdb/lmdb.h:1175
func DbiClose(env *Env, dbi Dbi) {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	C.mdb_dbi_close(cenv, cdbi)
}

// Drop function as declared in mdb/lmdb.h:1186
func Drop(txn *Txn, dbi Dbi, del int32) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	cdel, _ := (C.int)(del), cgoAllocsUnknown
	__ret := C.mdb_drop(ctxn, cdbi, cdel)
	__v := (int32)(__ret)
	return __v
}

// SetCompare function as declared in mdb/lmdb.h:1207
func SetCompare(txn *Txn, dbi Dbi, cmp CmpFunc) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ccmp, _ := cmp.PassRef()
	__ret := C.mdb_set_compare(ctxn, cdbi, ccmp)
	__v := (int32)(__ret)
	return __v
}

// SetDupsort function as declared in mdb/lmdb.h:1230
func SetDupsort(txn *Txn, dbi Dbi, cmp CmpFunc) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ccmp, _ := cmp.PassRef()
	__ret := C.mdb_set_dupsort(ctxn, cdbi, ccmp)
	__v := (int32)(__ret)
	return __v
}

// SetRelfunc function as declared in mdb/lmdb.h:1250
func SetRelfunc(txn *Txn, dbi Dbi, rel RelFunc) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	crel, _ := rel.PassRef()
	__ret := C.mdb_set_relfunc(ctxn, cdbi, crel)
	__v := (int32)(__ret)
	return __v
}

// SetRelctx function as declared in mdb/lmdb.h:1266
func SetRelctx(txn *Txn, dbi Dbi, ctx unsafe.Pointer) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	cctx, _ := (unsafe.Pointer)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.mdb_set_relctx(ctxn, cdbi, cctx)
	__v := (int32)(__ret)
	return __v
}

// Get function as declared in mdb/lmdb.h:1294
func Get(txn *Txn, dbi Dbi, key *Val, data *Val) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ckey, _ := key.PassRef()
	cdata, _ := data.PassRef()
	__ret := C.mdb_get(ctxn, cdbi, ckey, cdata)
	__v := (int32)(__ret)
	return __v
}

// Put function as declared in mdb/lmdb.h:1343
func Put(txn *Txn, dbi Dbi, key *Val, data *Val, flags uint32) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ckey, _ := key.PassRef()
	cdata, _ := data.PassRef()
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	__ret := C.mdb_put(ctxn, cdbi, ckey, cdata, cflags)
	__v := (int32)(__ret)
	return __v
}

// Del function as declared in mdb/lmdb.h:1368
func Del(txn *Txn, dbi Dbi, key *Val, data *Val) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ckey, _ := key.PassRef()
	cdata, _ := data.PassRef()
	__ret := C.mdb_del(ctxn, cdbi, ckey, cdata)
	__v := (int32)(__ret)
	return __v
}

// CursorOpen function as declared in mdb/lmdb.h:1392
func CursorOpen(txn *Txn, dbi Dbi, cursor **Cursor) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ccursor, _ := (**C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	__ret := C.mdb_cursor_open(ctxn, cdbi, ccursor)
	__v := (int32)(__ret)
	return __v
}

// CursorClose function as declared in mdb/lmdb.h:1400
func CursorClose(cursor *Cursor) {
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	C.mdb_cursor_close(ccursor)
}

// CursorRenew function as declared in mdb/lmdb.h:1418
func CursorRenew(txn *Txn, cursor *Cursor) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	__ret := C.mdb_cursor_renew(ctxn, ccursor)
	__v := (int32)(__ret)
	return __v
}

// CursorTxn function as declared in mdb/lmdb.h:1424
func CursorTxn(cursor *Cursor) *Txn {
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	__ret := C.mdb_cursor_txn(ccursor)
	__v := *(**Txn)(unsafe.Pointer(&__ret))
	return __v
}

// CursorDbi function as declared in mdb/lmdb.h:1430
func CursorDbi(cursor *Cursor) Dbi {
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	__ret := C.mdb_cursor_dbi(ccursor)
	__v := (Dbi)(__ret)
	return __v
}

// CursorGet function as declared in mdb/lmdb.h:1451
func CursorGet(cursor *Cursor, key *Val, data *Val, op CursorOp) int32 {
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	ckey, _ := key.PassRef()
	cdata, _ := data.PassRef()
	cop, _ := (C.MDB_cursor_op)(op), cgoAllocsUnknown
	__ret := C.mdb_cursor_get(ccursor, ckey, cdata, cop)
	__v := (int32)(__ret)
	return __v
}

// CursorPut function as declared in mdb/lmdb.h:1513
func CursorPut(cursor *Cursor, key *Val, data *Val, flags uint32) int32 {
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	ckey, _ := key.PassRef()
	cdata, _ := data.PassRef()
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	__ret := C.mdb_cursor_put(ccursor, ckey, cdata, cflags)
	__v := (int32)(__ret)
	return __v
}

// CursorDel function as declared in mdb/lmdb.h:1533
func CursorDel(cursor *Cursor, flags uint32) int32 {
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	__ret := C.mdb_cursor_del(ccursor, cflags)
	__v := (int32)(__ret)
	return __v
}

// CursorCount function as declared in mdb/lmdb.h:1547
func CursorCount(cursor *Cursor, countp *Size) int32 {
	ccursor, _ := (*C.MDB_cursor)(unsafe.Pointer(cursor)), cgoAllocsUnknown
	ccountp, _ := (*C.mdb_size_t)(unsafe.Pointer(countp)), cgoAllocsUnknown
	__ret := C.mdb_cursor_count(ccursor, ccountp)
	__v := (int32)(__ret)
	return __v
}

// Cmp function as declared in mdb/lmdb.h:1559
func Cmp(txn *Txn, dbi Dbi, a *Val, b *Val) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ca, _ := a.PassRef()
	cb, _ := b.PassRef()
	__ret := C.mdb_cmp(ctxn, cdbi, ca, cb)
	__v := (int32)(__ret)
	return __v
}

// Dcmp function as declared in mdb/lmdb.h:1571
func Dcmp(txn *Txn, dbi Dbi, a *Val, b *Val) int32 {
	ctxn, _ := (*C.MDB_txn)(unsafe.Pointer(txn)), cgoAllocsUnknown
	cdbi, _ := (C.MDB_dbi)(dbi), cgoAllocsUnknown
	ca, _ := a.PassRef()
	cb, _ := b.PassRef()
	__ret := C.mdb_dcmp(ctxn, cdbi, ca, cb)
	__v := (int32)(__ret)
	return __v
}

// ReaderList function as declared in mdb/lmdb.h:1588
func ReaderList(env *Env, _func MsgFunc, ctx unsafe.Pointer) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	c_func, _ := _func.PassRef()
	cctx, _ := (unsafe.Pointer)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.mdb_reader_list(cenv, c_func, cctx)
	__v := (int32)(__ret)
	return __v
}

// ReaderCheck function as declared in mdb/lmdb.h:1596
func ReaderCheck(env *Env, dead *int32) int32 {
	cenv, _ := (*C.MDB_env)(unsafe.Pointer(env)), cgoAllocsUnknown
	cdead, _ := (*C.int)(unsafe.Pointer(dead)), cgoAllocsUnknown
	__ret := C.mdb_reader_check(cenv, cdead)
	__v := (int32)(__ret)
	return __v
}
