package manager

import (
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type windowsManager struct {
	runner boshsys.CmdRunner
}

func NewWindowsManager(runner boshsys.CmdRunner) *windowsManager {
	return &windowsManager{runner: runner}
}

func (manager *windowsManager) SetPrimary(address string) error {
	stdout, _, _, err := manager.runner.RunCommand("powershell.exe", "/var/vcap/packages/dns-windows/bin/list-server-addresses.ps1")
	if err != nil {
		return bosherr.WrapError(err, "Executing list-server-addresses.ps1")
	}

	servers := strings.SplitN(stdout, "\r\n", 2)

	if servers[0] == address {
		return nil
	}

	_, _, _, err = manager.runner.RunCommand("powershell.exe", "/var/vcap/packages/dns-windows/bin/prepend-dns-server.ps1", address)
	if err != nil {
		return bosherr.WrapError(err, "Executing prepend-dns-server.ps1")
	}

	return nil
}

func (manager *windowsManager) Read() ([]string, error) {
	return nil, nil
}
