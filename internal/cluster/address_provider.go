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

package cluster

import (
	"fmt"
	"net"
	"strconv"

	pubcluster "github.com/semihbkgr/hazelcast-go-client/cluster"
	"github.com/semihbkgr/hazelcast-go-client/internal"
)

type AddressProvider interface {
	Addresses() ([]pubcluster.Address, error)
}

type DefaultAddressProvider struct {
	addresses []pubcluster.Address
}

func ParseAddress(addr string) (pubcluster.Address, error) {
	host, port, err := internal.ParseAddr(addr)
	if err != nil {
		return "", fmt.Errorf("parsing address: %w", err)
	}
	return pubcluster.Address(net.JoinHostPort(host, strconv.Itoa(port))), nil
}

func NewDefaultAddressProvider(networkConfig *pubcluster.NetworkConfig) *DefaultAddressProvider {
	var err error
	addresses := make([]pubcluster.Address, len(networkConfig.Addresses))
	for i, addr := range networkConfig.Addresses {
		if addresses[i], err = ParseAddress(addr); err != nil {
			panic(err)
		}
	}
	return &DefaultAddressProvider{addresses: addresses}
}

func (p DefaultAddressProvider) Addresses() ([]pubcluster.Address, error) {
	return p.addresses, nil
}
