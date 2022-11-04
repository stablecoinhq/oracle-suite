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

package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/chronicleprotocol/oracle-suite/pkg/ethereum"
	oracleGeth "github.com/chronicleprotocol/oracle-suite/pkg/price/oracle/geth"
)

func NewMegapokerCmd(opts *options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "megapoker",
		Args:  cobra.ExactArgs(1),
		Short: "commands related to the Megapoker contract",
		Long:  ``,
	}

	cmd.AddCommand(
		NewMegapokerPokeCmd(opts),
	)

	return cmd
}

func NewMegapokerPokeCmd(opts *options) *cobra.Command {
	return &cobra.Command{
		Use:   "poke megapoker_address [json_messages_list]",
		Args:  cobra.ExactArgs(1),
		Short: "directly invokes poke method",
		Long:  ``,
		RunE: func(_ *cobra.Command, args []string) error {
			srv, err := PrepareServices(opts)
			if err != nil {
				return err
			}
			megapoker := oracleGeth.NewMegapoker(srv.Client, ethereum.HexToAddress(args[0]))

			tx, err := megapoker.Poke(context.Background(), true)
			if err != nil {
				return err
			}

			fmt.Printf("Transaction: %s\n", tx.String())

			return nil
		},
	}
}
