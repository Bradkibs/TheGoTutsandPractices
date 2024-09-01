package distros

var (
	MachineID string
	Error     error
)

func init() {
	MachineID, Error = machineID()
}
