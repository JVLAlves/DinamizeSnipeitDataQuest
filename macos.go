package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// this is a comment

var linhas []string
var infos []string

func quit() {

	cmd := exec.Command(`osascript`, "-s", "h", "-e", `quit app "terminal"`)

	stderr, err := cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	// osascript -e 'tell application "Terminal" to do script "echo hello"'
	cmd := exec.Command(`osascript`, "-s", "h", "-e", `tell application "Terminal" to do script "uname -n > HOSTNAME.txt"`)
	// cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command(`osascript`, "-s", "h", "-e", `tell application "Terminal" to do script "sysctl -a |grep machdep.cpu.brand_string |awk '{print $2,$3,$4}' > CPU.txt
	"`)

	stderr, err = cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ = ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command(`osascript`, "-s", "h", "-e", `tell application "Terminal" to do script "hostinfo |grep memory |awk '{print $4,$5}'  > MEMORIA.txt"`)

	stderr, err = cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ = ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command(`osascript`, "-s", "h", "-e", `tell application "Terminal" to do script "diskutil list |grep disk0s2 |awk '{print $5,$6}' >> HOSTNAME.txt"`)

	stderr, err = cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ = ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command(`osascript`, "-s", "h", "-e", `tell application "Terminal" to do script "sw_vers -productVersion > SO.txt"`)

	stderr, err = cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ = ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	slurp, _ = ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	defer quit()

	file, err := os.Open("DISK.txt")
	if err != nil {
		log.Fatal(err)
	}

	bla, err := ioutil.ReadAll(file)
	fmt.Println("1. Dentro do filescanner" + string(bla))

	file.Seek(0, 0)
	fileScanner := bufio.NewScanner(file)
	linhas = []string{}

	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
		fmt.Printf("linhas = %#v\n", linhas)
	}
	fmt.Println("2. Dentro do filescanner" + fileScanner.Text())
	if fileScanner.Err() != nil {
		log.Fatalf("Erro SCAN: %v", fileScanner.Err().Error())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[0])
	file.Close()

	file, err = os.Open("HOSTNAME.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[1])
	file.Close()

	file, err = os.Open("CPU.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[2])
	file.Close()

	file, err = os.Open("MEMORIA.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[3])
	file.Close()

	file, err = os.Open("SO.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[4])
	file.Close()

	fmt.Println(infos)

	cmd = exec.Command(`osascript`, "-s", "h", "-e", `quit app "terminal"`)

	stderr, err = cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ = ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

}
