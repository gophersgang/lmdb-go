package lmdb

import (
	"errors"

	"github.com/zenhotels/lmdb-go/mdb"
)

func Version() (major, minor, patch int32, version string) {
	v := mdb.Version(&major, &minor, &patch)
	version = toString(v, 64)
	return
}

func mdbError(err int32) error {
	switch err {
	case mdb.ErrSuccess:
		return nil
	case mdb.ErrKeyExist:
		return ErrKeyExist
	case mdb.ErrNotFound:
		return ErrNotFound
	case mdb.ErrPageNotFound:
		return ErrPageNotFound
	case mdb.ErrCorrupted:
		return ErrCorrupted
	case mdb.ErrPanic:
		return ErrPanic
	case mdb.ErrVersionMismatch:
		return ErrVersionMismatch
	case mdb.ErrInvalid:
		return ErrInvalid
	case mdb.ErrMapFull:
		return ErrMapFull
	case mdb.ErrDbsFull:
		return ErrDbsFull
	case mdb.ErrReadersFull:
		return ErrReadersFull
	case mdb.ErrTlsFull:
		return ErrTlsFull
	case mdb.ErrTxnFull:
		return ErrTxnFull
	case mdb.ErrCursorFull:
		return ErrCursorFull
	case mdb.ErrPageFull:
		return ErrPageFull
	case mdb.ErrMapResized:
		return ErrMapResized
	case mdb.ErrIncompatible:
		return ErrIncompatible
	case mdb.ErrBadRslot:
		return ErrBadRslot
	case mdb.ErrBadTxn:
		return ErrBadTxn
	case mdb.ErrValueSize:
		return ErrValueSize
	case mdb.ErrBadDbi:
		return ErrBadDbi
	default:
		return errors.New(strError(err))
	}
}

func strError(err int32) string {
	return toString(mdb.StrError(err), 255)
}

var (
	ErrSuccess         = errors.New(strError(mdb.ErrSuccess))
	ErrKeyExist        = errors.New(strError(mdb.ErrKeyExist))
	ErrNotFound        = errors.New(strError(mdb.ErrNotFound))
	ErrPageNotFound    = errors.New(strError(mdb.ErrPageNotFound))
	ErrCorrupted       = errors.New(strError(mdb.ErrCorrupted))
	ErrPanic           = errors.New(strError(mdb.ErrPanic))
	ErrVersionMismatch = errors.New(strError(mdb.ErrVersionMismatch))
	ErrInvalid         = errors.New(strError(mdb.ErrInvalid))
	ErrMapFull         = errors.New(strError(mdb.ErrMapFull))
	ErrDbsFull         = errors.New(strError(mdb.ErrDbsFull))
	ErrReadersFull     = errors.New(strError(mdb.ErrReadersFull))
	ErrTlsFull         = errors.New(strError(mdb.ErrTlsFull))
	ErrTxnFull         = errors.New(strError(mdb.ErrTxnFull))
	ErrCursorFull      = errors.New(strError(mdb.ErrCursorFull))
	ErrPageFull        = errors.New(strError(mdb.ErrPageFull))
	ErrMapResized      = errors.New(strError(mdb.ErrMapResized))
	ErrIncompatible    = errors.New(strError(mdb.ErrIncompatible))
	ErrBadRslot        = errors.New(strError(mdb.ErrBadRslot))
	ErrBadTxn          = errors.New(strError(mdb.ErrBadTxn))
	ErrValueSize       = errors.New(strError(mdb.ErrValueSize))
	ErrBadDbi          = errors.New(strError(mdb.ErrBadDbi))
)
