package templates

const (
	GreetingMsg = `Hey @%s, welcome to the Sentinel Socks5 Proxy Service (SPS) for Telegram.`

	// Please select a blockchain network for payments to this bot.`
	NodeAttachedAlready = "You already have a node assigned to your username. Please use /mynode to access it"

	CheckWalletOptionsError = "Error while fetching user wallet address. In case you have not attached your wallet address, please share your wallet address again."
	Success                 = "Congratulations!! please click the button below to connect to the sentinel dVPN node"
	AskToSelectANode        = `Please select a node from the above list: `
	DATACONSUMPTION         = "Data consumed: %.2f MB of 1GB"
	LIMITEXCEEDED           = ` Hey @%s,
You have consumed 1GB data. Proxy will disconnect now. Please disable proxy in telegram settings.

Go to Settings -> Data & Storage -> Proxy Settings -> Long tap on the
connected proxy to Delete/Remove Proxy 
							(or) 
Toggle use proxy to disable proxy.`

	UserInfo = `Bandwidth Duration Left: <b>%0.0f days</b>
Ethereum Wallet Attached: <b>%s</b>`
	AskForEthWallet = "Please share your ethereum wallet address that you want to use for transactions to this bot"
	AskForPayment   = `please send %s $SENTS to the following address in next 30 minutes
and submit the transaction hash here.
Please note that if you submit the transaction hash after 30 minutes, it will be considered as invalid transaction.
However, you can use /refund to claim  your amount`
	AskForTMWallet    = "Please share your tendermint wallet address that you want to use for transactions to this bot"
	AskForBW          = "Please select how much bandwidth you need by clicking on one of the buttons below: "
	BWError           = "Error while storing bandwidth price"
	NodeList          = "%s.) %s, %s\n     Speed: %.2f Mbps\n     CPU-Load: %.2f%s"
	BWPeriods         = "You have opted for %s of unlimited bandwidth"
	Error             = "Could not read user info"
	BWAttachmentError = "Error occurred while adding user details for bandwidth requirements"
	ConnectMessage    = "If you are not connected to the proxy, please click on the button below to connect to Sentinel's SOCKS5 Proxy"
	NoEthNodes        = "No nodes available right now. please check again later or try our Tendermint network"
	NoTMNodes         = "No nodes available right now. please check again later or try our Ethereum network"
	NoAssignedNodes   = "There are no nodes assigned for you.. Get a node from here /sps"
	DisableProxy			= `Proxy is disconnected!! 
Please disable/remove proxy in telegram settings

To disconnect/terminate, go to Settings -> Data & Storage -> Proxy Settings -> Long tap on the connected proxy to Delete/Remove Proxy `
	InvalidOption     = "Invalid option"
	FollowSequence    = "Please follow the flow for the bot to work efficiently for you"
	NoNetworkSelected = "You have not selected a blockchain network to part of. Please select a network"
	NoMinBal          = `You don't have enough balance to use this bot.
BALANCE: %.3f $SUT
Minimum required balance is 10 $SUT. Please get some $SUTs and resubmit your Tendermint Wallet.
If you do not know how to get Sentinel Utilty Tokens, please ask the team @sentinel_co`
	AboutSentinel = `Sentinel Network is a network layer that enables a true p2p and decentralized applications and resources marketplace. Sentinel enables anyone to create Public and Private networks that provide access to both free and incentivized, and also payment method agnostic (pre-paid/escrow/post-paid) Services (or dApps) & distributed resources, enabling clients to become both producers and consumers in the network.

	Sentinel utilizes locking, staking and multi-sig directly from Tendermint core and Cosmos SDK and aims to fully eliminate the disadvantages of previous generation protocols, that couldn’t scale due to limitations of the blockchain they share with other dApps or that have an unsustainable economic model that reduces usability or access to the product they offer.
	
	
	 Website: https://sentinel.co
	 Github: https://github.com/sentinel-official/sentinel
	 Twitter: http://twitter.com/Sentinel_co`
	TXNNotFound = `Could not find the transaction hash in the network. A possible reason could be that you did the
transaction on a different Network as of the bot.`
	NotUniqueWallet = "This wallet has already been attached to another user. Please try with a different wallet address"
)
