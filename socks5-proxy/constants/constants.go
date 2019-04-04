package constants

import (
	"time"
)

const (
	EthNetwork          = "Rinkeby@Ethereum"
	TenderMintNetwork   = "SENTTEST@Tendermint"
	EthAddr             = "ETHADDR"
	Timestamp           = "TIMESTAMP"
	TimestampTM         = "TIMESTAMPTM"
	Node                = "NODE"
	NodeTM              = "NODETM"
	Bandwidth           = "BANDWIDTH"
	BandwidthTM         = "BANDWIDTHTM"
	NodeWallet          = "NODEWALLET"
	NodeWalletTM        = "NODEWaLLETTM"
	NodePrice           = "NODEPRICE"
	NodePriceTM         = "NODEPRICETM"
	BlockchainNetwork   = "BLOCKCHAINNETWORK"
	TMHashLength        = 40
	TimeLimit           = 30
	TMPrefix            = "cosmosaccaddr"
	WalletTM            = "WALLETTM"
	TMWalletLength      = 52
	TenDays             = time.Hour * 24 * 10
	Month               = time.Hour * 24 * 30
	ThreeMonths         = time.Hour * 24 * 90
	NodeBasePrice       = "10"
	NodeMonthPrice      = "30"
	NodeThreeMonthPrice = "80"
	ThreeM              = "90 Days"
	OneM                = "30 Days"
	TenD                = "10 Days"
	ZFill               = "000000000000000000000000"
	IPAddr              = "IPADDR"
	IPAddrTM            = "IPADDRTM"
	AssignedNodeURI     = "ASSIGNEDNODEURI"
	AssignedNodeURITM   = "ASSIGNEDNODEURITM"
	IsAuth              = "ISAUTH"
	IsAuthTM            = "ISAUTHTM"
	Password            = "PASSWORD"
	PasswordTM          = "PASSWORDTM"
	TestSentURI1        = `https://api-rinkeby.etherscan.io/api?apikey=Y5BJ5VA3XZ59F63XQCQDDUWU2C29144MMM&module=logs&action=getLogs&fromBlock=0&toBlock=latest&address=0x29317B796510afC25794E511e7B10659Ca18048B&topic0=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef&topic0_1_opr=and&topic1=`
	TestSendURI2        = `&topic1_2_opr=or&topic2=`
	TMTxnURL            = "http://localhost:1317/txs/%s"
	EthRegex            = "^(0x){1}[0-9a-fA-F]{40}$"
	ReplyButton         = "replyButton"
	InlineButton        = "inlineButton"
	ProxyURL            = "https://t.me/socks?server=%s&port=%s&user=%s&pass=%s"
	IPLEAKURL           = "https://ipleak.net/json/"
	NodeBaseUrl         = "http://%s:30002/user"
	GetTXNFromMN        = "http://35.154.179.57:8000/txes?fromAccountAddress=%s"
	PasswordLength      = 12
	SentinelTONURL      = "http://35.154.179.57:8000/nodes?type=OpenVPN&status=up"
	TMBalanceURL        = "http://localhost:1317/accounts/%s"
	NodeType            = "tendermint"
	EthState            = "ETHSTATE"
	TMTimeLimit         = "TMTimeLimit"
	TMState             = "TMSTATE"
	NoState             = -1
	MinBal              = 10000000
)

const (
	EthState0 = iota + 1
	EthState1
	EthState2
	EthState3
	EthState4
)

const (
	TMState0 = iota + 1
	TMState1
	TMState2
	TMState3
	TMState4
	TMState5
)
