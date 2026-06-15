/*
 * Copyright (c) 2025 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package raftmod

var RaftServices = []interface{}{
	RaftLogStoreFactory(),
	RaftStableStoreFactory(),
	RaftSnapshotFactory(),
	SerfConfigFactory(),
	ServerLookup(),
	SerfRPCServer(),
	RaftServer(),
	RaftClientPool(),
}
