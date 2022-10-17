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

package gofer

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/chronicleprotocol/oracle-suite/pkg/util/query"

	pkgEthereum "github.com/chronicleprotocol/oracle-suite/pkg/ethereum"
	"github.com/chronicleprotocol/oracle-suite/pkg/price/provider/origins"
)

// averageFromBlocks is a list of blocks distances from the latest blocks from
// which prices will be averaged.
var averageFromBlocks = []int64{0, 10, 20}

func parseParamsSymbolAliases(params yaml.Node) (origins.SymbolAliases, error) {
	var res struct {
		SymbolAliases origins.SymbolAliases `yaml:"symbolAliases"`
	}
	err := params.Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal origin symbol aliases from params: %w", err)
	}
	return res.SymbolAliases, nil
}

func parseParamsAPIKey(params yaml.Node) (string, error) {
	var res struct {
		APIKey string `yaml:"apiKey"`
	}
	err := params.Decode(&res)
	if err != nil {
		return "", fmt.Errorf("failed to marshal origin symbol aliases from params: %w", err)
	}
	return res.APIKey, nil
}

func parseParamsContracts(params yaml.Node) (origins.ContractAddresses, error) {
	var res struct {
		Contracts origins.ContractAddresses `yaml:"contracts"`
	}
	err := params.Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal origin symbol aliases from params: %w", err)
	}
	return res.Contracts, nil
}

//nolint:funlen,gocyclo,whitespace
func NewHandler(
	origin string,
	wp query.WorkerPool,
	cli pkgEthereum.Client,
	baseURL string,
	params yaml.Node,
) (origins.Handler, error) {
	aliases, err := parseParamsSymbolAliases(params)
	if err != nil {
		return nil, err
	}
	switch origin {
	case "balancer":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(origins.Balancer{
			WorkerPool:        wp,
			BaseURL:           baseURL,
			ContractAddresses: contracts,
		}, aliases), nil
	case "binance":
		return origins.NewBaseExchangeHandler(origins.Binance{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "bitfinex":
		return origins.NewBaseExchangeHandler(origins.Bitfinex{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "bitstamp":
		return origins.NewBaseExchangeHandler(origins.Bitstamp{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "bitthumb", "bithumb":
		return origins.NewBaseExchangeHandler(origins.BitThump{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "bittrex":
		return origins.NewBaseExchangeHandler(origins.Bittrex{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "coinbase", "coinbasepro":
		return origins.NewBaseExchangeHandler(origins.CoinbasePro{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "cryptocompare":
		return origins.NewBaseExchangeHandler(origins.CryptoCompare{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "coinmarketcap":
		apiKey, err := parseParamsAPIKey(params)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(
			origins.CoinMarketCap{WorkerPool: wp, BaseURL: baseURL, APIKey: apiKey},
			aliases,
		), nil
	case "currencyapi":
		return origins.NewBaseExchangeHandler(origins.Currencyapi{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "currencylayer":
		return origins.NewBaseExchangeHandler(origins.Currencylayer{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "ddex":
		return origins.NewBaseExchangeHandler(origins.Ddex{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "folgory":
		return origins.NewBaseExchangeHandler(origins.Folgory{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "fixer":
		return origins.NewBaseExchangeHandler(origins.Fixer{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "forex":
		return origins.NewBaseExchangeHandler(origins.Forex{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "ftx":
		return origins.NewBaseExchangeHandler(origins.Ftx{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "fx":
		apiKey, err := parseParamsAPIKey(params)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(
			origins.Fx{WorkerPool: wp, BaseURL: baseURL, APIKey: apiKey},
			aliases,
		), nil
	case "gateio":
		return origins.NewBaseExchangeHandler(origins.Gateio{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "gemini":
		return origins.NewBaseExchangeHandler(origins.Gemini{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "hitbtc":
		return origins.NewBaseExchangeHandler(origins.Hitbtc{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "huobi":
		return origins.NewBaseExchangeHandler(origins.Huobi{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "kraken":
		return origins.NewBaseExchangeHandler(origins.Kraken{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "kucoin":
		return origins.NewBaseExchangeHandler(origins.Kucoin{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "loopring":
		return origins.NewBaseExchangeHandler(origins.Loopring{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "okex":
		return origins.NewBaseExchangeHandler(origins.Okex{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "okx":
		return origins.NewBaseExchangeHandler(origins.Okx{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "openexchangerates":
		apiKey, err := parseParamsAPIKey(params)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(
			origins.OpenExchangeRates{WorkerPool: wp, BaseURL: baseURL, APIKey: apiKey},
			aliases,
		), nil
	case "poloniex":
		return origins.NewBaseExchangeHandler(origins.Poloniex{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "sushiswap":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(origins.Sushiswap{
			WorkerPool:        wp,
			BaseURL:           baseURL,
			ContractAddresses: contracts,
		}, aliases), nil
	case "curve", "curvefinance":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		h, err := origins.NewCurveFinance(cli, contracts, averageFromBlocks)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(*h, aliases), nil
	case "balancerV2":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		h, err := origins.NewBalancerV2(cli, contracts, averageFromBlocks)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(*h, aliases), nil
	case "wsteth":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		h, err := origins.NewWrappedStakedETH(cli, contracts, averageFromBlocks)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(*h, aliases), nil
	case "rocketpool":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		h, err := origins.NewRocketPool(cli, contracts, averageFromBlocks)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(*h, aliases), nil
	case "uniswap", "uniswapV2":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(origins.Uniswap{
			WorkerPool:        wp,
			BaseURL:           baseURL,
			ContractAddresses: contracts,
		}, aliases), nil
	case "uniswapV3":
		contracts, err := parseParamsContracts(params)
		if err != nil {
			return nil, err
		}
		return origins.NewBaseExchangeHandler(origins.UniswapV3{
			WorkerPool:        wp,
			BaseURL:           baseURL,
			ContractAddresses: contracts,
		}, aliases), nil
	case "upbit":
		return origins.NewBaseExchangeHandler(origins.Upbit{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	case "xe":
		return origins.NewBaseExchangeHandler(origins.Xe{WorkerPool: wp, BaseURL: baseURL}, aliases), nil
	}

	return nil, origins.ErrUnknownOrigin
}
