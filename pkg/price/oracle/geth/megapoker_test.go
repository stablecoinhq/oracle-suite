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
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/chronicleprotocol/oracle-suite/pkg/ethereum"
	"github.com/chronicleprotocol/oracle-suite/pkg/ethereum/mocks"
)

func TestMegapoker_Poke(t *testing.T) {
	// Prepare test data:
	c := &mocks.Client{}
	a := ethereum.Address{}
	m := NewMegapoker(c, a)

	c.On("SendTransaction", mock.Anything, mock.Anything).Return(&ethereum.Hash{}, nil)

	// Call Poke function:
	_, err := m.Poke(context.Background(), false)
	assert.NoError(t, err)

	// Verify generated transaction:
	tx := c.Calls[0].Arguments.Get(1).(*ethereum.Transaction)
	cd := "18178358"

	assert.Equal(t, a, tx.Address)
	assert.Equal(t, (*big.Int)(nil), tx.MaxFee)
	assert.Equal(t, big.NewInt(gasLimit), tx.GasLimit)
	assert.Equal(t, uint64(0), tx.Nonce)
	assert.Equal(t, cd, hex.EncodeToString(tx.Data))
}
