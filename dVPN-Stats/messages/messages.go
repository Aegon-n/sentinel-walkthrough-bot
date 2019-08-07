package messages

const (
	StatsHomeMsg = "Hey @%s,\n*Welcome to the Sentinel Network dVPN Statistics*"
	ChooseOption = "please choose an option from the list below:"
	UnableToGetStats = "Unable to get statistics"
	StatsMsg = `*Sentinel dVPN - Stats*
	1. Current Active Nodes : *%d nodes*
	2. Daily Average Active Nodes : *%d nodes*
	3. Current Active Sessions : *%d sessions (connected users)*
	4. Daily Average Active Sessions : *%d sessions (connected users)*
	5. Data Consumed in the last 24 hours : *%.2fGB*
	6. Total Data exchanged on the Sentinel network : *%.2fTB*`
	ActiveNodesListMsg = "Here it is: *List Of Active dVPN - Nodes*"
	NodeList = "%s.) %s, %s\n     Speed: %.2f Mbps\n     CPU-Load: %.2f%s"
)