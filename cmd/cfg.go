package cmd

import "github.com/spf13/viper"

// config keys
const (
	kBvnGw = "bgw" // cfg key for brevis gateway to submit signed resp

	kBlsPath = "blspath" // file path to bls priv key file
	kBlsPwd  = "blspwd"  // password for bls key file
	// ecdsa keystore json, needed to send onchain tx
	kKsPath = "kspath"
	kKsPwd  = "kspwd"
	// Eigen operator id from registry, 0x prefix optional
	kOpId = "operID"
	// key for onechainconfig, may appear many times in config file
	kMultichain = "multichain"
	kDb         = "db.url"
)

// config related defs and func
func GetMcc(key string) MultiChainConfig {
	var m MultiChainConfig
	viper.UnmarshalKey(key, &m)
	return m
}

type MultiChainConfig []*OneChainConfig

type OneChainConfig struct {
	ChainID       uint64
	Name, Gateway string
	// blk related for monitor
	BlkInterval, BlkDelay        uint64
	MaxBlkDelta, ForwardBlkDelay uint64
	// in case missed event, restart from specific block number
	// StartBlk uint64

	Brevis   string // hex addr of brevisRequest contract on this chain, monitor req event
	BrvEigen string // hex addr of BrevisEigen, needed in sign data to avoid replay attack

	BrvRegCo string // registry coordinator address, needed for join
	AVSDir   string // AVSDirectory contract on this chain, needed for salt sig in join

	// gateway only
	OpStateRetriever, BlsApkReg string
	// start query NewPubkeyRegistration from this blk num
	BlsApkStartBlk uint64
}
