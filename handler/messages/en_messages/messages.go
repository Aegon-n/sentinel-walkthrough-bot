package en_messages

var (
	LangSelectMsg = map[string]string{"English":"Please choose your language from the list below: ","Russian":"Пожалуйста, выберите ваш язык из списка ниже"}
	LangChosenMsg = "LangChosenMsg"
	SelectwalkthroughMsg = "Please select /help to know commands available for the Sentinel Network bot"
	SelectSpsMessage = "To start using Sentinel's SOCKS5 Proxy Service, use the /sps command"
	WelcomeGreetMsg = `Hello @%s,

*Welcome to the Sentinel Network*`
	WalkthroughGreetMsg = `Hey @%s,

*Welcome to the Sentinel Network walkthrough*`
	ExitMsg = "\n\n\n\tThank you for using the Sentinel app walkthrough guide on the Sentinel Network bot\t\n\n\n"
	AppSelectMsg = "Choose App from the list below:"
	DesktopOSSelectMsg = "*Sentinel -> Desktop*\n\nChoose OS from the list below: "
	MobileOSSelectMsg = "*Sentinel -> Mobile*\n\nChoose OS from the list below: "

	LinuxNetworkSelectMsg = "*Sentinel -> Desktop-> Linux*\n\nChoose a network from the list below: "
	WindowsNetworkSelectMsg = "*Sentinel -> Desktop-> Windows*\n\nChoose a network from the list below: "
	MacNetworkSelectMsg = "*Sentinel -> Desktop -> Mac*\n\nChoose a network from the list below: "
	AndroidNetworkSelectMsg = "*Sentinel -> Desktop->Android*\n\nChoose a network from the list below: "
	IOSNetworkSelectMsg = "*Sentinel -> Desktop->IOS*\n\nChoose a network from the list below: "
	SelectUpdateBlog = "Select a button below for getting the latest updates"
	EthWinListOfModulesMsg = `*Here are the list of modules the for Ethereum TestNet*

	1. Downloading & installation of Sentinel Desktop Client

	2. Wallet Creation on Ethereum (ETH)

	3. Sentinel Mainnet

	4. TestNet Activation

	5. Receiving free test tokens

	6. Connecting to dVPN

	7. Disconnecting dVPN`

	EthWindowsModule10 = ` *1. Download & Installation steps of Sentinel Desktop Client*
	
	*Step 1*: Go to https://Sentinel.co

	*Step 2*: Scroll down and you will see the option to select your operating system and download an executable for that. Please download and install the '.exe' file`

	EthWindowsModule20 = `
	*2. Wallet Creation on Ethereum*
						
	*Step 1*: Open Sentinel App click create/restore wallet

	*Step 2*: Enter a Anonymous ID password to Create Account

	*Step 3*: Click create to create a new wallet
										(or)
					
	*Select a keystore file if you want restore your previous account*
						  
	*Step 4*: Copy your wallet address & public_key and store
									the public_key securely.

	*Step 5*: Click on checklist and Go to Sentinel-MainNet Dashboard
					
	`
	EthWindowsModule30 = ` 
	*3. MainNet Sending & Receiving 'SENT' Tokens*

	*Step 1*: Type Recipient wallet Address

	*Step 2*: Type Amount 

	*Step 3*: Select the gas price (maximum gas price will 
				result in faster transactions)

	*Step 4*: Type your wallet password and click send
						
						`

	EthWindowsModule40 = `
	*4. TestNet Activation* 
						
	*Step 1*: Toggle to activate Ethereum TestNet

	*Step 2*: Get free test tokens by clicking get free test tokens

	*Step 3*: verify your balance at top left corner

	`

	EthWindowsModule50 = `
	*5. Connecting to dVPN*
						
	*Step 1*: click the globe Icon in the menu to get available vpn nodes

	*Step 2*: select a protocol for your connection

	*Step 4*: check your IP before connecting to a node

	*Step 3*: connect to one of those nodes by clicking on node

	*Step 4*: Do payment for connection

	*Step 5*: Try connecting Node again & you will get connected to node

	*Step 6*: Verify connection of node by checking your public IP.
						
	`
	EthWindowsModule60 = `
	*6. Disconnecting dVPN*
						
	*Step 1*: Click disconnect button to disconnect node

	*Step 2*: Add your Rating to Node

	*Step 3*: Check session details in the sessions section

	`

	TMWindowsModule40 = `
	*4. TestNet Activation* 
	
	*Step 1*: Toggle to activate TestNet

	*Step 2*: Select Tendermint TestNet from dropdown list

	*Step 3*: Create a new Tendermint wallet account by providing
					ananymous name and password

	*Step 4*: Get free tokens by clicking get free tokens

	*Step 5*: Verify your balance at top left corner

	`
	MobilewalletInstallMsg = "download the latest Sentinel apk file from https://github.com/Sentinel-official/Sentinel/releases and install"
	MobileListOfMOdulesMsg = "module1\nmodule2\n module3"

	AndroidMobileListOfModulesMsg = `*Here is the list of Modules for Android Mobile wallet*

	1. Downloading & installation of Sentinel-Mobile-wallet Application

	2. Wallet Creation on Ethereum

	3. Sentinel mainnet

	4. TestNet Activation

	5. Receiving free test tokens

	6. Connecting to dVPN

	7. Disconnecting dVPN`
	IOSMobileListOfModulesMsg = "IOS version Currently Not Available"

	AndroidModule10 =` 
		*1.Download & Installation steps of Sentinel Mobile wallet Application*

		*Step 1*: Go to https://github.com/Sentinel-official/Sentinel/releases

		*Step 2*: Find latest .apk file ,Download and install it.
	`

	AndroidModule20 = `
	*2.Wallet Creation on Ethereum*
						
	*Step 1*: Open Sentinel App click create/restore wallet

	*Step 2*: Enter a Anonymous ID password to Create Account

	*Step 3*: Click create to create a new wallet
										(or)
					
	*select a keystore file if you want restore your previous account*
						  
	*Step 4*: Copy your wallet address & public_key and store
									the public_key securely.

	*Step 5*: Click on checklist and Go to Sentinel-MainNet Dashboard
					
	`
	AndroidModule30 =
		`
		*3.MainNet Sending & Receiving 'SENT' Tokens*

		*Step 1*: Type Recipient wallet Address

		*Step 2*: Type Amount 

		*Step 3*: Select the gas price (maximum gas price will 
				result in faster transactions)

		*Step 4*: Type your wallet password and click send
		`
	AndroidModule40 =
		`*4.TestNet Activation* 
		
		*Step 1*: Toggle to activate Ethereum TestNet

		*Step 2*: Get free test tokens by clicking get free test tokens

		*Step 3*: Verify your balance at top left corner

		`
	AndroidModule50 =
		`
		*5.Connecting to dVPN*
						
		*Step 1*: Click the globe Icon in the menu to get available vpn nodes

		*Step 2*: Select a protocol for your connection

		*Step 4*: Check your IP before connecting to a node

		*Step 3*: Connect to one of those nodes by clicking on node

		*Step 4*: Do payment for connection

		*Step 5*: Try connecting Node again & you will get connected to node

		*Step 6*: Verify connection of node by checking your public IP.
						
		`
	AndroidModule60 =
		`
		*6.Disconnecting Sentinel dVPN*
						
		*Step 1*: Click disconnect button to disconnect node

		*Step 2*: Add your Rating to Node

		*Step 3*: Check session details in the sessions section

		`
	LastModuleMsg = `                    *All Chapters Completed*`


	Socks5GreetingMsg = `Hey %s, welcome to the Sentinel Socks5 Proxy Bot for Telegram.
			Please select a blockchain network for payments to this bot.`

	Socks5EthereumMsg = "Ethereum is currently not available. Please choose the Tendermint Network"

	NodeAttachedAlready = "You already have a node assigned to your username. Please use /mynode to access it"

	CheckWalletOptionsError = "Error while fetching user wallet address. in case you have not attached your wallet address, please share your wallet address again."
	Success                 = "Congratulations!! Please click the button below to connect to the Sentinel dVPN node and next time use /mynode to access this node"
	AskToSelectANode        = `Please select a node ID from the list below and reply in the format of
1 for Node 1, 2 for Node 2 and so on...`
	UserInfo = `Bandwidth Duration Left: <b>%0.0f days</b>
Ethereum Wallet Attached: <b>%s</b>`
	AskForEthWallet   = "Please share your ethereum wallet address that you want to use for transactions to this bot"
	AskForPayment     = "Please send %s $SENTS to the following address and submit the transaction hash here: "
	AskForTMWallet    = "Please share your tendermint wallet address that you want to use for transactions to this bot"
	AskForBW          = "Please select how much bandwidth you need by clicking on one of the buttons below: "
	BWError           = "error while storing bandwidth price"
	NodeList          = "%s.) Location: %s\n User: %s \n Node wallet: %s"
	BWPeriods         = "You have opted for %s of unlimited bandwidth"
	Error             = "could not read user info"
	BWAttachmentError = "error occurred while adding user details for bandwidth requirements"
	ConnectMessage    = "Please click on the button below to connect to Sentinel's SOCKS5 Proxy"
	NoEthNodes        = "No nodes available right now. please check again later or try our Tendermint network"
	NoTMNodes         = "No nodes available right now. please check again later or try our Ethereum network"
	InvalidOption     = "This is an invalid option, please use /help to check out all commands"
	HelpMsg           = `
		<b>Here are the available commands and their utility</b>

		1. /help - to get help about commands
		2. /start - start the Sentinel bot
		3. /about - to know about Sentinel Network
		4. /sps - to get Sentinel proxy service for telegram
		5. /mynode - to get list of assigned sps proxy nodes
		6. /restart_sps -  to restart the process of sps
		7. /disconect_proxy - to disconnect the proxy
		8. /downloads - to get download links for Sentinel clients
		9. /guides - Guides to participate in Sentinel Network
		10. /stats - to Know dVPN statistics of Sentinel Network
		11. /restart - to restart the Sentinel Network Bot
	`

)
