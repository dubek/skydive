/*
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package flow

import (
	"encoding/binary"
	"hash"

	"github.com/google/gopacket"
)

// Hash computes the hash of a ICMP layer
func (fl *ICMPLayer) Hash(hasher hash.Hash) {
	if fl == nil {
		return
	}

	value32 := make([]byte, 4)
	binary.BigEndian.PutUint32(value32, uint32(fl.Type)<<24|uint32(fl.Code<<16|uint32(fl.ID)))
	hasher.Write(value32)
}

// Hash computes the hash of a transport layer
func (tl *TransportLayer) Hash(hasher hash.Hash, swap bool) {
	if tl == nil {
		return
	}

	value32 := make([]byte, 4)
	if swap {
		binary.BigEndian.PutUint32(value32, uint32(tl.B<<16|tl.A))
	} else {
		binary.BigEndian.PutUint32(value32, uint32(tl.A<<16|tl.B))
	}
	hasher.Write(value32)

	valueID := make([]byte, 8)
	binary.BigEndian.PutUint64(valueID, uint64(tl.ID))
	hasher.Write(valueID)
}

// Hash calculates a unique symetric flow layer hash
func (fl *FlowLayer) Hash(hasher hash.Hash, swap bool) {
	if fl == nil {
		return
	}

	if swap {
		hasher.Write([]byte(fl.B))
		hasher.Write([]byte(fl.A))
	} else {
		hasher.Write([]byte(fl.A))
		hasher.Write([]byte(fl.B))
	}

	value64 := make([]byte, 8)
	binary.BigEndian.PutUint64(value64, uint64(fl.ID))
	hasher.Write(value64)
}

// hashFlow flow with custom hasher function
func hashFlow(f gopacket.Flow, hasher hash.Hash, swap bool) {
	src, dst := f.Endpoints()
	if swap {
		hasher.Write(dst.Raw())
		hasher.Write(src.Raw())
	} else {
		hasher.Write(src.Raw())
		hasher.Write(dst.Raw())
	}

	value64 := make([]byte, 8)
	binary.BigEndian.PutUint64(value64, uint64(f.EndpointType()))
	hasher.Write(value64)
}
