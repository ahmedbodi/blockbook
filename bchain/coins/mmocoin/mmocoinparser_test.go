// +build unittest

package mmocoin

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"
	"encoding/hex"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
)

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func Test_GetAddrDescFromAddress_Mainnet(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "P2PKH1",
			args:    args{address: "MAMSLb8aRGKiSzCqQZtVAibdrtHjbKhEdJ"},
			want:    "76a9141ae882e788091732da6910595314447c9e38bd8d88ac",
			wantErr: false,
		},
	}
	parser := NewMmocoinParser(GetChainParams("main"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.GetAddrDescFromAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddrDescFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("GetAddrDescFromAddress() = %v, want %v", h, tt.want)
			}
		})
	}
}

var (
	testTx1 bchain.Tx

	testTxPacked1 = "0a206ba18524d81af732d0226ffdb63d2bcdc0d58a35ac97b5ad731057932d324e1412b401010000001134415d0114caae2bf9a7808aee0798e6245a347405d46c8131dbf55cbbbc689bbee367e902000000484730440220280f3fa80b4e93834fe0a8d9884105310eaa8d36d77b9aff113b6c498138e5bb02204578409f0a14fa1950ea4951314fd495fd503b42a6325efb5c139a6c8253912401ffffffff0200000000000000000005f22f5904000000232102bdb95d89f07e3a29305f3c8de86ec211ed77b7e15cf314c85c532a6b71c2ce07ac000000001891e884ea05200028b88a5432760a001220e967e3be9b68bcbb5cf5db31816cd40574345a24e69807ee8a80a7f92baeca14180222484730440220280f3fa80b4e93834fe0a8d9884105310eaa8d36d77b9aff113b6c498138e5bb02204578409f0a14fa1950ea4951314fd495fd503b42a6325efb5c139a6c825391240128ffffffff0f3a0210003a520a0504583af7fb10011a232102bdb95d89f07e3a29305f3c8de86ec211ed77b7e15cf314c85c532a6b71c2ce07ac2222446d343835624e4a6169474a6d4556746832426e5a345931796763756644736934454001"
)

func init() {
	testTx1 = bchain.Tx{
		Hex:       "010000001134415d0114caae2bf9a7808aee0798e6245a347405d46c8131dbf55cbbbc689bbee367e902000000484730440220280f3fa80b4e93834fe0a8d9884105310eaa8d36d77b9aff113b6c498138e5bb02204578409f0a14fa1950ea4951314fd495fd503b42a6325efb5c139a6c8253912401ffffffff0200000000000000000005f22f5904000000232102bdb95d89f07e3a29305f3c8de86ec211ed77b7e15cf314c85c532a6b71c2ce07ac00000000",
		Blocktime: 1564554257,
		Txid:      "6ba18524d81af732d0226ffdb63d2bcdc0d58a35ac97b5ad731057932d324e14",
		LockTime:  0,
		Time:      1564554257,
		Version:   1,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "4730440220280f3fa80b4e93834fe0a8d9884105310eaa8d36d77b9aff113b6c498138e5bb02204578409f0a14fa1950ea4951314fd495fd503b42a6325efb5c139a6c8253912401",
				},
				Txid:     "e967e3be9b68bcbb5cf5db31816cd40574345a24e69807ee8a80a7f92baeca14",
				Vout:     2,
				Sequence: 4294967295,
			},
		},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(0),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "",
				},
			},
			{
				ValueSat: *big.NewInt(18660128763),
				N:        1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "2102bdb95d89f07e3a29305f3c8de86ec211ed77b7e15cf314c85c532a6b71c2ce07ac",
					Addresses: []string{
						"Dm485bNJaiGJmEVth2BnZ4Y1ygcufDsi4E",
					},
				},
			},
		},
	}
}

func Test_PackTx(t *testing.T) {
	type args struct {
		tx        bchain.Tx
		height    uint32
		blockTime int64
		parser    *MmocoinParser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "mmocoin-1",
			args: args{
				tx:        testTx1,
				height:    1377592,
				blockTime: 1564554257,
				parser:    NewMmocoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.parser.PackTx(&tt.args.tx, tt.args.height, tt.args.blockTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("packTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("packTx() = %v, want %v", h, tt.want)
			}
		})
	}
}

func Test_UnpackTx(t *testing.T) {
	type args struct {
		packedTx string
		parser   *MmocoinParser
	}
	tests := []struct {
		name    string
		args    args
		want    *bchain.Tx
		want1   uint32
		wantErr bool
	}{
		{
			name: "mmocoin-1",
			args: args{
				packedTx: testTxPacked1,
				parser:   NewMmocoinParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx1,
			want1:   1377592,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.packedTx)
			got, got1, err := tt.args.parser.UnpackTx(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackTx(1) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackTx(2) got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("unpackTx(3) got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}