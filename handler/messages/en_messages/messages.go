package en_messages

var (
	LangSelectMsg = map[string]string{"English":"Please choose your language from the list below: ","Russian":"Пожалуйста, выберите ваш язык из списка ниже"}
	LangChosenMsg = "Successfully chosen %s Language"
	SelectwalkthroughMsg = "Please select /walkthrough for sentinel App walkthrough "
	WelcomeGreetMsg = "Hey %s , Welcome to sentinel App walkthrough session."
	ExitMsg = "\n\n\n\tThank you for using sentinel app walkthrough session\t\n\n\n"
	AppSelectMsg = "Choose App from the list below: "
	DesktopOSSelectMsg = "*Sentinel->Desktop*\n\nChoose OS from the list below: "
	MobileOSSelectMsg = "*Sentinel->Mobile*\n\nChoose OS from the list below: "

	LinuxNetworkSelectMsg = "*Sentinel->Desktop->Linux*\n\nChoose a network from the list below: "
	WindowsNetworkSelectMsg = "*Sentinel->Desktop->Windows*\n\nChoose a network from the list below: "
	MacNetworkSelectMsg = "*Sentinel->Desktop->Mac*\n\nChoose a network from the list below: "
	AndroidNetworkSelectMsg = "*Sentinel->Desktop->Android*\n\nChoose a network from the list below: "
	IOSNetworkSelectMsg = "*Sentinel->Desktop->IOS*\n\nChoose a network from the list below: "

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

	LastModuleMsg = `                    *All Chapters Completed*                 `
)