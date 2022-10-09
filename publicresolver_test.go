package jazdid

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const networkUrl = "https://data-seed-prebsc-1-s1.binance.org:8545/"
const resolverAddr = "0xDeDcF005618C8c130212ff29EBA2EeaD1125737E"
const registrarAddr = "0xFe166d97c891A47c732c8b564A842e05dac10eD9"
const reverseRegistrarAddr = "0xE498E54f421E03C05631093c66AB3219bD83B102"
const bulkControllerAddr = "0xbd418F36bC5cAB3B3A81ecd9a6a461b454328d92"

func TestResolver(t *testing.T) {
	r, err := NewPublicResolver(context.TODO(), networkUrl, resolverAddr, registrarAddr, reverseRegistrarAddr, bulkControllerAddr)
	if err != nil {
		t.Error(err)
	}
	userAddr := "0x0a0b364093cB37787827E210806f50c30CE4e192"
	domain, name, err := r.Name(common.HexToAddress(userAddr))
	require.Nil(t, err, "Error resolving name")
	fmt.Println("resolving domain:", domain, "name:", name)

	addr, err := r.Address(domain)
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, "0x0a0b364093cB37787827E210806f50c30CE4e192", addr.String(), "Unexpected address")

	b, err := r.MultiAddress(60, domain)
	require.Nil(t, err, "Error resolving multiAddress")
	assert.Equal(t, "0x0a0b364093cb37787827e210806f50c30ce4e192", "0x"+hex.EncodeToString(b), "Unexpected multiAddress")

	avatar, err := r.Text(domain, Avatar)
	require.Nil(t, err, "Error resolving avatar")
	assert.Equal(t, "https://ipfs.io/ipfs/QmVSc7yZKefGpJwjh7QWSRBGNTnu9R5jnWdZFPy9fB2sgt", avatar, "Unexpected avatar")

	url, err := r.Text(domain, Url)
	require.Nil(t, err, "Error resolving url")
	assert.Equal(t, "https://www.jazdid.com/", url, "Unexpected url")

	description, err := r.Text(domain, Description)
	require.Nil(t, err, "Error resolving description")
	assert.Equal(t, "The Best Multichain DID Aggregator", description, "Unexpected description")

	keywords, err := r.Text(domain, Keywords)
	require.Nil(t, err, "Error resolving keywords")
	assert.Equal(t, "kind, clever", keywords, "Unexpected keywords")

	email, err := r.Text(domain, Email)
	require.Nil(t, err, "Error resolving email")
	assert.Equal(t, "dev@jazdid.com", email, "Unexpected email")

	twitter, err := r.Text(domain, Twitter)
	require.Nil(t, err, "Error resolving twitter")
	assert.Equal(t, "https://twitter.com/JAZ_DID", twitter, "Unexpected twitter")

	discord, err := r.Text(domain, Discord)
	require.Nil(t, err, "Error resolving discord")
	assert.Equal(t, "https://discord.gg/jazdid", discord, "Unexpected discord")

	telegram, err := r.Text(domain, Telegram)
	require.Nil(t, err, "Error resolving telegram")
	assert.Equal(t, "https://t.me/JazDIDofficial", telegram, "Unexpected telegram")

	medium, err := r.Text(domain, Medium)
	require.Nil(t, err, "Error resolving medium")
	assert.Equal(t, "https://www.medium.com/@jazdid", medium, "Unexpected medium")

	instagram, err := r.Text(domain, Instagram)
	require.Nil(t, err, "Error resolving instagram")
	assert.Equal(t, "https://www.instagram.com/jazdid", instagram, "Unexpected instagram")

	github, err := r.Text(domain, Github)
	require.Nil(t, err, "Error resolving github")
	assert.Equal(t, "https://github.com/jazdid", github, "Unexpected github")
}

func TestNameRecords(t *testing.T) {
	r, err := NewPublicResolver(context.TODO(), networkUrl, resolverAddr, registrarAddr, reverseRegistrarAddr, bulkControllerAddr)
	if err != nil {
		t.Error(err)
	}
	userAddr := "0x0a0b364093cB37787827E210806f50c30CE4e192"
	domain, _, err := r.Name(common.HexToAddress(userAddr))
	require.Nil(t, err, "Error resolving name")

	addr, vals, err := r.NameRecords(domain, []string{Avatar, Url, Description, Keywords, Email, Twitter, Discord,
		Telegram, Medium, Instagram, Github})
	require.Nil(t, err, "Error resolving name records")

	fmt.Println("resolving addr:", addr, "vals:", vals)
}
