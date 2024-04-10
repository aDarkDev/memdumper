package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetMemoryAddresses(data *[]byte) []string {
	r, _ := regexp.Compile(`([0-9a-f]{3,})-([0-9a-f]{3,})`)
	matches := r.FindAllStringSubmatch(string(*data), -1)
	lList := []string{}
	for _, match := range matches {
		address := "0x" + match[1] + " " + "0x" + match[2]
		lList = append(lList, address)
	}
	return lList
}

func DumpAddress(pid string, from string, to string) {
	command := fmt.Sprintf(`gdb --batch --pid %s -ex "dump memory %s-%s-%s.dump %s %s"`, pid, pid, from, to, from, to)
	cmd := exec.Command("bash", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[x] permission denied to dump \"" + from + "\" \"" + to + "\".")
		fmt.Println("sudo", os.Args[0], os.Args[1])
		os.Exit(1)
	}
	fmt.Printf("[+] dumped address from \"%s\" to \"%s\"\n", from, to)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(os.Args[0], "<pid>")
		os.Exit(1)
	}

	pid := os.Args[1]
	data, err := os.ReadFile("/proc/" + pid + "/maps")
	if err != nil {
		fmt.Println("[x] pid", pid, "not found.")
		os.Exit(1)
	}

	fmt.Print("\t\tMemory Dumper - v0.1 - aDarkDev\n\n")
	addresses := GetMemoryAddresses(&data)
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines)-1; i++{
		fmt.Println(i, lines[i])
	}

	fmt.Print("\nchoose number, example 1:9 mean 1 to 9 or just a single number: ")
	var input string
	fmt.Scan(&input)
	fmt.Println("")

	if strings.Contains(input, ":") {
		splited := strings.Split(input, ":")
		start, err := strconv.Atoi(splited[0])
		if err != nil {
			fmt.Println("[x] input is not int.")
			os.Exit(1)
		}
		end, err := strconv.Atoi(splited[1])
		if err != nil {
			fmt.Println("[x] input is not int.")
			os.Exit(1)
		}

		for i := start; i <= end; i++ {
			spl := strings.Split(addresses[i], " ")
			from := spl[0]
			to := spl[1]
			DumpAddress(pid, from, to)
		}
	} else {
		index, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("[x] input is not int.")
			os.Exit(1)
		}

		spl := strings.Split(addresses[index], " ")
		from := spl[0]
		to := spl[1]
		DumpAddress(pid, from, to)
	}

	fmt.Println("\nDumped Successfully.")
}
