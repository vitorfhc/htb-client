package consts

type HTBPath string

const (
	HTBPathListRetiredLabMachines HTBPath = "/api/v4/machine/list/retired/paginated"
	HTBPathListActiveLabMachines  HTBPath = "/api/v4/machine/paginated"
	HTBPathVPNServers             HTBPath = "/api/v4/connections/servers"
	HTBPathActiveLabMachine       HTBPath = "/api/v4/machine/active"
	HTBPathSpawnLabMachine        HTBPath = "/api/v4/vm/spawn"
	HTBPathTerminateLabMachine    HTBPath = "/api/v4/vm/terminate"
)

const HTBHost = "https://www.hackthebox.eu"
