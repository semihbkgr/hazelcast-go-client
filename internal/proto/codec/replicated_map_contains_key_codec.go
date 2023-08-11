/*
 * Copyright (c) 2008-2021, Hazelcast, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package codec

import (
	"github.com/semihbkgr/hazelcast-go-client/internal/proto"
	iserialization "github.com/semihbkgr/hazelcast-go-client/internal/serialization"
)

const (
	// hex: 0x0D0400
	ReplicatedMapContainsKeyCodecRequestMessageType = int32(852992)
	// hex: 0x0D0401
	ReplicatedMapContainsKeyCodecResponseMessageType = int32(852993)

	ReplicatedMapContainsKeyCodecRequestInitialFrameSize = proto.PartitionIDOffset + proto.IntSizeInBytes

	ReplicatedMapContainsKeyResponseResponseOffset = proto.ResponseBackupAcksOffset + proto.ByteSizeInBytes
)

// Returns true if this map contains a mapping for the specified key.

func EncodeReplicatedMapContainsKeyRequest(name string, key iserialization.Data) *proto.ClientMessage {
	clientMessage := proto.NewClientMessageForEncode()
	clientMessage.SetRetryable(true)

	initialFrame := proto.NewFrameWith(make([]byte, ReplicatedMapContainsKeyCodecRequestInitialFrameSize), proto.UnfragmentedMessage)
	clientMessage.AddFrame(initialFrame)
	clientMessage.SetMessageType(ReplicatedMapContainsKeyCodecRequestMessageType)
	clientMessage.SetPartitionId(-1)

	EncodeString(clientMessage, name)
	EncodeData(clientMessage, key)

	return clientMessage
}

func DecodeReplicatedMapContainsKeyResponse(clientMessage *proto.ClientMessage) bool {
	frameIterator := clientMessage.FrameIterator()
	initialFrame := frameIterator.Next()

	return FixSizedTypesCodec.DecodeBoolean(initialFrame.Content, ReplicatedMapContainsKeyResponseResponseOffset)
}
