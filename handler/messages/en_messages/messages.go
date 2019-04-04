package en_messages

var (
	LangSelectMsg = map[string]string{"English":"Please choose your language from the list below: ","Russian":"Пожалуйста, выберите ваш язык из списка ниже"}
	LangChosenMsg = "LangChosenMsg"
	SelectwalkthroughMsg = "Please select /help to know the available commands for this bot"
	WelcomeGreetMsg = "Hey %s , Welcome to sentinel-bot."
	ExitMsg = "\n\n\n\tThank you for using sentinel app walkthrough session\t\n\n\n"
	AppSelectMsg = "Choose App from the list below: "
	DesktopOSSelectMsg = "*Sentinel->Desktop*\n\nChoose OS from the list below: "
	MobileOSSelectMsg = "*Sentinel->Mobile*\n\nChoose OS from the list below: "

	LinuxNetworkSelectMsg = "*Sentinel->Desktop->Linux*\n\nChoose a network from the list below: "
	WindowsNetworkSelectMsg = "*Sentinel->Desktop->Windows*\n\nChoose a network from the list below: "
	MacNetworkSelectMsg = "*Sentinel->Desktop->Mac*\n\nChoose a network from the list below: "
	AndroidNetworkSelectMsg = "*Sentinel->Desktop->Android*\n\nChoose a network from the list below: "
	IOSNetworkSelectMsg = "*Sentinel->Desktop->IOS*\n\nChoose a network from the list below: "
	SelectUpdateBlog = "select a button below for getting latest updates"
	EthWinListOfModulesMsg = `*Here is the list of Modules for Ethereum TestNet*

	1. Downloading & installation of sentinel-Desktop Application

	2. Wallet Creation on Ethereum

	3. Sentinel-MainNet

	4. TestNet Activation

	5. Receiving free test tokens

	6. Connecting to dVpn

	7. Disconnecting dVpn`

	EthWindowsModule10 = ` *1.Download & Installation steps of sentinel-Desktop Application*
	
	*step1*: Go to https://sentinel.co

	*step2*: Scroll down you will see a .exe file.
				Please download and install it.

	`

	EthWindowsModule20 = `
	*2.Wallet Creation on Ethereum*
						
	*step1*: Open sentinel App click create/restore wallet

	*step2*: Enter a Anonymous ID password to Create Account

	*step3*: Click create to create a new wallet
										(or)
					
	*select a keystore file if you want restore your previous account*
						  
	*step4*: Copy your wallet address & public_key and store
									the public_key securely.

	*step5*: Click on checklist and Go to Sentinel-MainNet Dashboard
					
	`
	EthWindowsModule30 = ` 
	*3.MainNet Sending & Receiving 'SENT' Tokens*

	*step1*: Type Recipient wallet Address

	*step2*: Type Amount 

	*step3*: select the gas price (maximum gas price will 
				result in faster transactions)

	*step4*: type your wallet password and click send
						
						`

	EthWindowsModule40 = `
	*4.TestNet Activation* 
						
	*step1*: Toggle to activate Ethereum TestNet

	*step2*: Get free test tokens by clicking get free test tokens

	*step3*: verify your balance at top left corner

	`

	EthWindowsModule50 = `
	*5.Connecting to dVpn*
						
	*step1*: click the globe Icon in the menu to get available vpn nodes

	*step2*: select a protocol for your connection

	*step4*: check your IP before connecting to a node

	*step3*: connect to one of those nodes by clicking on node

	*step4*: Do payment for connection

	*step5*: Try connecting Node again & you will get connected to node

	*step6*: Verify connection of node by checking your public IP.
						
	`
	EthWindowsModule60 = `
	*6.Disconnecting dVpn*
						
	*step1*: Click disconnect button to disconnect node

	*step2*: Add your Rating to Node

	*step3*: Check session details in the sessions section

	`

	TMWindowsModule40 = `
	*4.TestNet Activation* 
	
	*step1*: Toggle to activate TestNet

	*step2*: select Tendermint TestNet from dropdown list

	*step3*: Create a new Tendermint wallet account by providing
					ananymous name and password

	*step4*: Get free tokens by clicking get free tokens

	*step5*: verify your balance at top left corner

	`
	MobilewalletInstallMsg = "download the latest sentinel apk file from https://github.com/sentinel-official/sentinel/releases and install"
	MobileListOfMOdulesMsg = "module1\nmodule2\n module3"

	AndroidMobileListOfModulesMsg = `*Here is the list of Modules for Android Mobile wallet*

	1. Downloading & installation of sentinel-Mobile-wallet Application

	2. Wallet Creation on Ethereum

	3. Sentinel-MainNet

	4. TestNet Activation

	5. Receiving free test tokens

	6. Connecting to dVpn

	7. Disconnecting dVpn`
	IOSMobileListOfModulesMsg = "IOS version Currently Not Available"

	AndroidModule10 =` 
		*1.Download & Installation steps of sentinel Mobile wallet Application*

		*step1*: Go to https://github.com/sentinel-official/sentinel/releases

		*step2*: Find latest .apk file ,Download and install it.
	`

	AndroidModule20 = `
	*2.Wallet Creation on Ethereum*
						
	*step1*: Open sentinel App click create/restore wallet

	*step2*: Enter a Anonymous ID password to Create Account

	*step3*: Click create to create a new wallet
										(or)
					
	*select a keystore file if you want restore your previous account*
						  
	*step4*: Copy your wallet address & public_key and store
									the public_key securely.

	*step5*: Click on checklist and Go to Sentinel-MainNet Dashboard
					
	`
	AndroidModule30 =
		`
		*3.MainNet Sending & Receiving 'SENT' Tokens*

		*step1*: Type Recipient wallet Address

		*step2*: Type Amount 

		*step3*: select the gas price (maximum gas price will 
				result in faster transactions)

		*step4*: type your wallet password and click send
		`
	AndroidModule40 =
		`*4.TestNet Activation* 
		
		*step1*: Toggle to activate Ethereum TestNet

		*step2*: Get free test tokens by clicking get free test tokens

		*step3*: verify your balance at top left corner

		`
	AndroidModule50 =
		`
		*5.Connecting to dVpn*
						
		*step1*: click the globe Icon in the menu to get available vpn nodes

		*step2*: select a protocol for your connection

		*step4*: check your IP before connecting to a node

		*step3*: connect to one of those nodes by clicking on node

		*step4*: Do payment for connection

		*step5*: Try connecting Node again & you will get connected to node

		*step6*: Verify connection of node by checking your public IP.
						
		`
	AndroidModule60 =
		`
		*6.Disconnecting dVpn*
						
		*step1*: Click disconnect button to disconnect node

		*step2*: Add your Rating to Node

		*step3*: Check session details in the sessions section

		`
	LastModuleMsg = `                    *All Chapters Completed*`


	Socks5GreetingMsg = `Hey %s, welcome to the Sentinel Socks5 Proxy Bot for Telegram.
			Please select a blockchain network for payments to this bot.`

	Socks5EthereumMsg = "Ethereum is currently not available Please choose Tendermint Network"

	NodeAttachedAlready = "you already have a node assigned to your username. Please use /mynode to access it"

	CheckWalletOptionsError = "error while fetching user wallet address. in case you have not attached your wallet address, please share your wallet address again."
	Success                 = "Congratulations!! please click the button below to connect to the sentinel dVPN node and next time use /mynode to access this node"
	AskToSelectANode        = `Please select a node ID from the list below and reply in the format of
1 for Node 1, 2 for Node 2 and so on...`
	UserInfo = `Bandwidth Duration Left: <b>%0.0f days</b>
Ethereum Wallet Attached: <b>%s</b>`
	AskForEthWallet   = "Please share your ethereum wallet address that you want to use for transactions to this bot"
	AskForPayment     = "please send %s $SENTS to the following address and submit the transaction hash here: "
	AskForTMWallet    = "Please share your tendermint wallet address that you want to use for transactions to this bot"
	AskForBW          = "Please select how much bandwidth you need by clicking on one of the buttons below: "
	BWError           = "error while storing bandwidth price"
	NodeList          = "%s.) Location: %s\n User: %s \n Node wallet: %s"
	BWPeriods         = "you have opted for %s of unlimited bandwidth"
	Error             = "could not read user info"
	BWAttachmentError = "error occurred while adding user details for bandwidth requirements"
	ConnectMessage    = "please click on the button below to connect to Sentinel's SOCKS5 Proxy"
	NoEthNodes        = "no nodes available right now. please check again later or try our Tendermint network"
	NoTMNodes         = "no nodes available right now. please check again later or try our Ethereum network"
	InvalidOption     = "invalid option"
	HelpMsg           = `
		<b>here is the available commands list and their usage</b>

		1. /help - to get help about commands
		2. /start - to start sentinel bot
		3. /walkthrough - to get sentinel app walkthrough
		4. /tm {latestblock| validators} - to explore the sentinel tendermint testnet
		5. /sps - to get sentinel proxy service for telegram
		6. /restart_sps -  to restart the process of sps
		7. /sps_info - to get details about sps node
		8. /sps_wallet - to know attached wallet address
		9. /mynode - to get list of assigned sps proxy nodes
		10. /updates - to get updates about sentinel Network
	`

)