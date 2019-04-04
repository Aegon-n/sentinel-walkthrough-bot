package models

import "time"

type Nodes struct {
	NodesList   []TONNode `json:"nodes"`
}
type Location struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
}
type NetSpeed struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}
type TONNode struct {
	Location       Location `json:"location"`
	NetSpeed       NetSpeed `json:"netSpeed"`
	APIPort        int      `json:"APIPort"`
	PricePerGB     int      `json:"pricePerGB"`
	Description    string   `json:"description"`
	RatingPoints   int      `json:"ratingPoints"`
	RatingCount    int      `json:"ratingCount"`
	AccountAddress string   `json:"accountAddress"`
	IP             string   `json:"IP"`
	EncMethod      string   `json:"encMethod"`
	NodeType       string   `json:"nodeType"`
	Version        string   `json:"version"`
	TxHash         string   `json:"txHash"`
}

type KV struct {
	Key   string
	Value string
}

type User struct {
	Timestamp        time.Time
	Auth             bool
	Node             string
	Password         string
	BW               string
	URI              string
	EthAddr          string
	TelegramUsername string
}

type TxReceipt struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	Timestamp        string   `json:"timestamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUser          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type TXDetails struct {
	Result struct {
		BlockHash        string `json:"blockHash"`
		BlockNumber      string `json:"blockNumber"`
		From             string `json:"from"`
		Gas              string `json:"gas"`
		GasPrice         string `json:"gasPrice"`
		Hash             string `json:"hash"`
		Input            string `json:"input"`
		Nonce            string `json:"nonce"`
		To               string `json:"to"`
		TransactionIndex string `json:"transactionIndex"`
		Value            string `json:"value"`
		V                string `json:"v"`
		R                string `json:"r"`
		S                string `json:"s"`
	} `json:"result"`
}

type TMTxn struct {
	Tx struct {
		Value struct {
			Msg []TMMsg `json:"msg"`
			Fee struct {
				Amount []TMAmount `json:"amount"`
			} `json:"fee"`
		} `json:"value"`
	} `json:"tx"`
}

type TMMsg struct {
	Type  string `json:"type"`
	Value struct {
		Coins   []TMAmount `json:"Coins"`
		From    string     `json:"From"`
		To      string     `json:"To"`
		Address string     `json:"address"`
	} `json:"value"`
}

type TMAmount struct {
	Denom  string
	Amount string
}

type TxReceiptList struct {
	Results []TxReceipt `json:"result"`
	Status  string      `json:"status"`
}

type GeoLocation struct {
	//As string `json:"as"`
	City        string `json:"city_name"`
	Country     string `json:"country_name"`
	CountryCode string `json:"country_code"`
	//Isp string `json:"isp"`
	//Lat float64 `json:"lat"`
	//Lon float64 `json:"lon"`
	//Query string `json:"query"`
	//Region string `json:"region"`
	RegionName string `json:"region_name"`
	//Status string `json:"status"`
	//TimeZone string `json:"timezone"`
	//Zip string `json:"zip"`
}

type AddUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RemoveUser struct {
	Username string `json:"username"`
}

type ExpiredUsers struct {
	Key   string
	Value string
}

type UserResp struct {
	Message  string `json:"messages"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type InlineButtonOptions struct {
	Label string
	URL   string
}

type ButtonHelper struct {
	Type               string
	Labels             []string
	InlineKeyboardOpts []InlineButtonOptions
}

type Transactions struct {
	List []MNTXs `json:"txes"`
}

type MNTXs struct {
	From      string `json:"fromAccountAddress"`
	To        string `json:"toAccountAddress"`
	Hash      string `json:"txHash"`
	Timestamp string `json:"addedOn"`
}
