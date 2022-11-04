//  Copyright (C) 2020 Maker Ecosystem Growth Holdings, INC.
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Affero General Public License as
//  published by the Free Software Foundation, either version 3 of the
//  License, or (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Affero General Public License for more details.
//
//  You should have received a copy of the GNU Affero General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

package geth

import (
	"context"
	"math/big"

	"github.com/chronicleprotocol/oracle-suite/pkg/ethereum"
)

// Megapoker implements the Megapoker interface using go-ethereum packages.
type Megapoker struct {
	ethereum ethereum.Client
	address  ethereum.Address
}

// NewMegapoker creates the new Median instance.
func NewMegapoker(ethereum ethereum.Client, address ethereum.Address) *Megapoker {
	return &Megapoker{
		ethereum: ethereum,
		address:  address,
	}
}

func (m *Megapoker) Poke(ctx context.Context, simulateBeforeRun bool) (*ethereum.Hash, error) {
	if simulateBeforeRun {
		if _, err := m.read(ctx, "poke"); err != nil {
			return nil, err
		}
	}

	return m.write(ctx, "poke")
}

func (m *Megapoker) read(ctx context.Context, method string, args ...interface{}) ([]interface{}, error) {
	cd, err := megapokerABI.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	var data []byte
	err = retry(maxReadRetries, delayBetweenReadRetries, func() error {
		data, err = m.ethereum.Call(ctx, ethereum.Call{Address: m.address, Data: cd})
		return err
	})
	if err != nil {
		return nil, err
	}

	return megapokerABI.Unpack(method, data)
}

func (m *Megapoker) write(ctx context.Context, method string, args ...interface{}) (*ethereum.Hash, error) {
	cd, err := megapokerABI.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	return m.ethereum.SendTransaction(ctx, &ethereum.Transaction{
		Address:  m.address,
		GasLimit: new(big.Int).SetUint64(gasLimit),
		Data:     cd,
	})
}
