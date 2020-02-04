package mmocoin

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xf2dbf1d1
)

// chain parameters
var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{31}
	MainNetParams.ScriptHashAddrID = []byte{78}
	MainNetParams.Bech32HRPSegwit = "dpn"
}

type MmocoinParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

func NewMmocoinParser(params *chaincfg.Params, c *btc.Configuration) *MmocoinParser {
	return &MmocoinParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
}

func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
		chaincfg.RegisterBitcoinParams()
	}
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	return &MainNetParams
}

// PackTx packs transaction to byte array using protobuf
func (p *MmocoinParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *MmocoinParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}