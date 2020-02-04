package mmocoin

import (
	"blockbook/bchain/coins/btc"
        "github.com/martinboehm/btcd/wire"
        "github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0x61199116
	TestnetMagic wire.BitcoinNet = 0xcbf2c0ef
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{50}
	MainNetParams.ScriptHashAddrID = []byte{100}
	MainNetParams.Bech32HRPSegwit = "mmo"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111}
	TestNetParams.ScriptHashAddrID = []byte{196}
	TestNetParams.Bech32HRPSegwit = "bc"
}

type MmocoinParser struct {
	*btc.BitcoinParser
}

func NewMmocoinParser(params *chaincfg.Params, c *btc.Configuration) *MmocoinParser {
	return &MmocoinParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
	}
}

func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}