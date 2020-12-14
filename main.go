package main

import (
	"flag"
	"fmt"
	"github.com/ca-gip/gensshconfig/internal/services"
	"github.com/ca-gip/gensshconfig/internal/utils"
	"net"
	"os"
	"path/filepath"
)

func main() {

	// TODO : Add options for ignored folder
	inventory := flag.String("inventory", "", "Inventory folder to extract clusters host")
	bastion := flag.String("bastion", "", "IP Address of the bastion")
	user := flag.String("user", "", "User for the bastion")

	flag.Parse()
	checkRequiredFlag()

	inventoryPath, err := filepath.Abs(*inventory)
	if err != nil {
		fmt.Printf("counld not find %s", inventory)
		os.Exit(1)
	}

	bastionIP := net.ParseIP(*bastion)
	if bastionIP == nil {
		fmt.Println("--bastion should be an IP Addr")
		os.Exit(1)
	}

	config := services.NewSSHConfig(
		inventoryPath,
		utils.DefaultIgnoredFolders,
		*user,
		"k8s_config",
		&services.Host{Hostname: "bastion", Addr: bastionIP})

	err = config.FindCluster()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = config.BuildClusterInventory()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.Render()

}

func checkRequiredFlag() {
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range utils.RequiredArgs {
		if !seen[req] {
			_, _ = fmt.Fprintf(os.Stderr, "missing required --%s argument\n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}
}
