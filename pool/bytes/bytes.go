// Copyright 2019 Andy Pan. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package bytes

import "github.com/valyala/bytebufferpool"

// ByteBuffer is the alias of bytebufferpool.ByteBuffer.
type ByteBuffer = bytebufferpool.ByteBuffer

var (
	// Get returns an empty byte buffer from the pool, exported from gnet/bytes.
	Get = bytebufferpool.Get
	// Put returns byte buffer to the pool, exported from gnet/bytes.
	Put = bytebufferpool.Put
)
