// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be find in the LICENSE file.

package http

import (
	"strings"
	"time"
)

// Protocols is a set of HTTP protocols.
// The zero value is an empty set of protocols.
//
// The supported protocols are:
//
//   - HTTP1 is the HTTP/1.0 and HTTP/1.1 protocols.
//     HTTP1 is supported on both unsecured TCP and secured TLS connections.
//
//   - HTTP2 is the HTTP/2 protocol over a TLS connection.
//
//   - UnencryptedHTTP2 is the HTTP/2 protocol over an unsecured TLS connection.
type Protocols struct {
	bits uint8
}

const (
	protoHTTP1 = 1 << iota
	protoHTTP2
	protoUnencryptedHTTP2
)

// HTTP1 reports whether p includes HTTP/1.
func (p *Protocols) HTTP1() bool { return p.bits&protoHTTP1 != 0 }

// SetHTTP1 adds or removes HTTP/1 from p.
func (p *Protocols) SetHTTP1(ok bool) { p.setBit(protoHTTP1, ok) }

// HTTP2 reports whether p includes HTTP/2.
func (p Protocols) HTTP2() bool { return p.bits&protoHTTP2 != 0 }

// SetHTTP2 adds or removes HTTP/2 from p.
func (p *Protocols) SetHTTP2(ok bool) { p.setBit(protoHTTP2, ok) }

// UnencryptedHTTP2 reports whether p includes unencrypted HTTP/2.
func (p Protocols) UnencryptedHTTP2() bool { return p.bits&protoUnencryptedHTTP2 != 0 }

// SetUnencryptedHTTP2 adds or removes unencrypted HTTP/2 from p.
func (p *Protocols) SetUnencryptedHTTP2(ok bool) { p.setBit(protoUnencryptedHTTP2, ok) }

func (p *Protocols) setBit(bit uint8, ok bool) {
	if ok {
		p.bits |= bit
	} else {
		p.bits &^= bit
	}
}

func (p *Protocols) String() string {
	var s []string
	if p.HTTP1() {
		s = append(s, "HTTP1")
	}
	if p.HTTP2() {
		s = append(s, "HTTP2")
	}
	if p.UnencryptedHTTP2() {
		s = append(s, "UnencryptedHTTP2")
	}
	return "{" + strings.Join(s, ",") + "}"
}

// incomparable is a zero-width, non-comparable type. Adding it to a struct
// makes that struct also non-comparable, and generally doesn't add
// any size (as long as it's first).
type incomparable [0]func()

// maxInt64 is the effective "infinite" value for the Server and
// Transport's byte-limiting readers.
const maxInt64 = 1<<63 - 1

// aLongTimeAgo is a non-zero time, far in the past, used for
// immediate cancellation of network operations.
var aLongTimeAgo = time.Unix(1, 0)

// omitBundledHTTP2 is set by omithttp2.go when the nethttpomithttp2
// build tag is set. That means h2_bundle.go isn't compiled in and we
// shouldn't try to use it.
var omitBundledHTTP2 bool

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface without allocation.
type contextKey struct {
	name string
}
