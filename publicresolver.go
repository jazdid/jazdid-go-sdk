package jazdid

import (
	"bytes"
	"compress/zlib"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"jazdid-go-sdk/abis/controller"
	"jazdid-go-sdk/abis/registrar"
	"jazdid-go-sdk/abis/resolver"
	"jazdid-go-sdk/abis/reverseregistrar"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// UnknownAddress is the address to which unknown entries resolve
var UnknownAddress = common.HexToAddress("00")

const (
	Avatar      string = "avatar"
	Url         string = "url"
	Description string = "description"
	Keywords    string = "keywords"
	Email       string = "email"
	Twitter     string = "com.twitter"
	Discord     string = "com.discord"
	Telegram    string = "org.telegram"
	Medium      string = "com.medium"
	Instagram   string = "com.instagram"
	Github      string = "com.github"
)

type PublicResolver struct {
	*ethclient.Client
	Resolver         *resolver.Resolver
	Registrar        *registrar.Registrar
	ReverseRegistrar *reverseregistrar.Reverseregistrar
	BulkController   *controller.Controller
	RegistrarAddr    common.Address
}

func NewPublicResolver(ctx context.Context, networkUrl, resolverAddr, registrarAddr, reverseRegistrarAddr,
	bulkControllerAddr string) (*PublicResolver, error) {
	client, err := ethclient.DialContext(ctx, networkUrl)
	if err != nil {
		return nil, err
	}

	r, err := resolver.NewResolver(common.HexToAddress(resolverAddr), client)
	if err != nil {
		return nil, err
	}
	re, err := registrar.NewRegistrar(common.HexToAddress(registrarAddr), client)
	if err != nil {
		return nil, err
	}
	rr, err := reverseregistrar.NewReverseregistrar(common.HexToAddress(reverseRegistrarAddr), client)
	if err != nil {
		return nil, err
	}
	bc, err := controller.NewController(common.HexToAddress(bulkControllerAddr), client)
	if err != nil {
		return nil, err
	}

	return &PublicResolver{
		client,
		r,
		re,
		rr,
		bc,
		common.HexToAddress(registrarAddr),
	}, nil
}

// Address returns the Ethereum address of the domain
func (r *PublicResolver) Address(domain string) (common.Address, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return UnknownAddress, err
	}
	return r.Resolver.Addr(nil, nameHash)
}

// SetAddress sets the Ethereum address of the domain
func (r *PublicResolver) SetAddress(opts *bind.TransactOpts, domain string, address common.Address) (*types.Transaction, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.SetAddr(opts, nameHash, address)
}

// MultiAddress returns the address of the domain for a given coin type.
// The coin type is as per https://github.com/satoshilabs/slips/blob/master/slip-0044.md
func (r *PublicResolver) MultiAddress(coinType uint64, domain string) ([]byte, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.AddrWithCoinType(nil, nameHash, big.NewInt(int64(coinType)))
}

// SetMultiAddress sets the iaddress of the domain for a given coin type.
// The coin type is as per https://github.com/satoshilabs/slips/blob/master/slip-0044.md
func (r *PublicResolver) SetMultiAddress(opts *bind.TransactOpts, coinType uint64, domain string, address []byte) (*types.Transaction, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.SetAddrWithCoinType(opts, nameHash, big.NewInt(int64(coinType)), address)
}

// PubKey returns the public key of the domain
func (r *PublicResolver) PubKey(domain string) ([32]byte, [32]byte, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}
	res, err := r.Resolver.Pubkey(nil, nameHash)
	return res.X, res.Y, err
}

// SetPubKey sets the public key of the domain
func (r *PublicResolver) SetPubKey(opts *bind.TransactOpts, domain string, x [32]byte, y [32]byte) (*types.Transaction, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.SetPubkey(opts, nameHash, x, y)
}

// Contenthash returns the content hash of the domain
func (r *PublicResolver) Contenthash(domain string) ([]byte, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.Contenthash(nil, nameHash)
}

// SetContenthash sets the content hash of the domain
func (r *PublicResolver) SetContenthash(opts *bind.TransactOpts, domain string, contenthash []byte) (*types.Transaction, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.SetContenthash(opts, nameHash, contenthash)
}

// InterfaceImplementer returns the address of the contract that implements the given interface for the given domain
func (r *PublicResolver) InterfaceImplementer(domain string, interfaceID [4]byte) (common.Address, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return UnknownAddress, err
	}
	return r.Resolver.InterfaceImplementer(nil, nameHash, interfaceID)
}

// SetText sets the text associated with a name
func (r *PublicResolver) SetText(opts *bind.TransactOpts, domain string, key string, value string) (*types.Transaction, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.SetText(opts, nameHash, key, value)
}

// Text obtains the text associated with a name
func (r *PublicResolver) Text(domain string, key string) (string, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return "", err
	}
	return r.Resolver.Text(nil, nameHash, key)
}

// SetABI sets the ABI associated with a name
func (r *PublicResolver) SetABI(opts *bind.TransactOpts, domain string, abi string, contentType *big.Int) (*types.Transaction, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	var data []byte
	switch contentType.Uint64() {
	case 1:
		// Uncompressed JSON
		data = []byte(abi)
	case 2:
		// Zlib-compressed JSON
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		if _, err := w.Write([]byte(abi)); err != nil {
			return nil, err
		}
		w.Close()
		data = b.Bytes()
	default:
		return nil, errors.New("unsupported content type")
	}

	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	return r.Resolver.SetABI(opts, nameHash, contentType, data)
}

// ABI returns the ABI associated with a name
func (r *PublicResolver) ABI(domain string) (string, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	contentTypes := big.NewInt(3)
	nameHash, err := NameHash(name)
	if err != nil {
		return "", err
	}
	contentType, data, err := r.Resolver.ABI(nil, nameHash, contentTypes)
	var a string
	if err == nil {
		if contentType.Cmp(big.NewInt(1)) == 0 {
			// Uncompressed JSON
			a = string(data)
		} else if contentType.Cmp(big.NewInt(2)) == 0 {
			// Zlib-compressed JSON
			b := bytes.NewReader(data)
			var z io.ReadCloser
			z, err = zlib.NewReader(b)
			if err != nil {
				return "", err
			}
			defer z.Close()
			var uncompressed []byte
			uncompressed, err = ioutil.ReadAll(z)
			if err != nil {
				return "", err
			}
			a = string(uncompressed)
		}
	}
	return a, nil
}

func (r *PublicResolver) Name(addr common.Address) (domain, name string, err error) {
	n, err := r.ReverseRegistrar.Node(nil, addr)
	if err != nil {
		return "", "", err
	}
	name, err = r.Resolver.Name(nil, n)
	if err != nil {
		return "", "", err
	}
	sn := strings.Split(name, ".")
	if len(sn) < 1 {
		return "", "", errors.New("invalid domain name")
	}

	a, err := r.Address(sn[0])
	if err != nil {
		return "", "", err
	}
	if addr != a {
		return "", "", nil
	}

	tld, err := r.Registrar.Tld(nil)
	if err != nil {
		return "", "", err
	}
	return sn[0], fmt.Sprintf("%s.%s", sn[0], tld), nil
}

func (r *PublicResolver) NameRecords(domain string, keys []string) (common.Address, []string, error) {
	name := fmt.Sprintf("%s.%s", strings.ToLower(domain), strings.TrimPrefix(r.RegistrarAddr.String(), "0x"))
	nameHash, err := NameHash(name)
	if err != nil {
		return UnknownAddress, nil, err
	}
	return r.BulkController.NameRecords(nil, r.RegistrarAddr, nameHash, keys)
}
