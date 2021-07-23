package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

func main() {
	posh := New()

	stdout, stderr, err := posh.Execute("New-Item -path \"$env:userprofile\" -Name \"logfileees\" -ItemType \"directory\"")
	stdout, stderr, err = posh.Execute("New-Item -path \"logfileees\" -Name \"logfileees.txt\" -ItemType \"file\"")
	stdout, stderr, err = posh.Execute("Get-ItemProperty HKLM:\\Software\\Wow6432Node\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\* | Select-Object DisplayName, DisplayVersion | Sort-Object -Property DisplayName -Unique | Format-Table -AutoSize > \"$env:userprofile\\logfileees\\Programas_Instalados.txt\" | Systeminfo > \"$env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt\"")
	stdout, stderr, err = posh.Execute("Get-Content -path \"$env:userprofile\\logfileees\\Programas_Instalados.txt\" | Add-Content -path \"$env:userprofile\\logfileees\\logfileees.txt\"")
	stdout, stderr, err = posh.Execute("Get-WmiObject -Class Win32_Processor -ComputerName . | Select-Object -Property \"name\" > \"$env:userprofile\\logfileees\\cpu.txt\"")
	stdout, stderr, err = posh.Execute("get-WMIobject Win32_LogicalDisk -Filter \"DeviceID = 'C:'\" | Select-Object -Property \"Size\" > \"$env:userprofile\\logfileees\\disk.txt\"")
	stdout, stderr, err = posh.Execute("(Get-Content -path \"$env:userprofile\\logfileees\\cpu.txt\" -TotalCount 6)[3] | Add-Content -path \"$env:userprofile\\logfileees\\logfileees.txt\"")
	stdout, stderr, err = posh.Execute("(Get-Content -path \"$env:userprofile\\logfileees\\disk.txt\" -TotalCount 6)[3] | Add-Content -path \"$env:userprofile\\logfileees\\logfileees.txt\"")
	stdout, stderr, err = posh.Execute("Add-Content -Path \"$env:userprofile\\logfileees\\logfileees.txt\" -value (Select-String -Path \"$env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt\" -Pattern \"Nome do host:\",\"Nome do sistema operacional:\",\"Memória física total:\")")
	fmt.Println(stdout)
	fmt.Println(stderr)

	if err != nil {
		fmt.Println(err)
	}
}
