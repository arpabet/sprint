/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package fsmod

import (
	"go.arpabet.com/sprint/fs"
	"google.golang.org/protobuf/encoding/protojson"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// default size is 64kb, possible to overwrite
var DefaultBufferSize = 64 * 1024

type fileServiceImpl struct {
	bufferSize int // read/write block buffer size
	marshaler  runtime.JSONPb
}

func FileService() fs.FileService {
	return &fileServiceImpl{
		bufferSize: DefaultBufferSize,
		marshaler: runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}
}

func (t *fileServiceImpl) BufferSize() int {
	return t.bufferSize
}

func (t *fileServiceImpl) SetBufferSize(size int) {
	t.bufferSize = size
}

func (t *fileServiceImpl) MarshalOptions() protojson.MarshalOptions {
	return t.marshaler.MarshalOptions
}

func (t *fileServiceImpl) SetMarshalOptions(opt protojson.MarshalOptions) {
	t.marshaler.MarshalOptions = opt
}

func (t *fileServiceImpl) UnmarshalOptions() protojson.UnmarshalOptions {
	return t.marshaler.UnmarshalOptions
}

func (t * fileServiceImpl) SetUnmarshalOptions(opt protojson.UnmarshalOptions) {
	t.marshaler.UnmarshalOptions = opt
}



