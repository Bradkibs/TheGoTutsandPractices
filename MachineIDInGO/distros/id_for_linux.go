package distros

const (
	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"
)

func machineID() (string, error) {
	id, err := ReadFile(dbusPath)
	if err != nil {
		id, err = ReadFile(dbusPathEtc)
	}
	if err != nil {
		return "", err
	}
	return Trim(string(id)), nil
}
