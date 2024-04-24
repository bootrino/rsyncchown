// rsyncchown 
// Andrew Stuart 24 Apr 2024
// MIT License

// see the readme at https://github.com/bootrino/rsyncchown for explanation

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func sanitizePath(path string) string {
	path = filepath.Clean(path)
	if !filepath.IsAbs(path) {
		path, _ = filepath.Abs(path)
	}
	return path
}

func runRsync(args []string) {
	rsyncArgs := []string{} // to store filtered args for rsync
	for _, arg := range args {
		if !strings.HasPrefix(arg, "--chown=") {
			rsyncArgs = append(rsyncArgs, arg)
		}
	}

	fmt.Println("Running rsync...")
	rsyncCmd := exec.Command("rsync", rsyncArgs...)
	rsyncCmd.Stdout = os.Stdout
	rsyncCmd.Stderr = os.Stderr
	fmt.Println("Rsync command:", rsyncCmd.String())
	err := rsyncCmd.Run()
	if err != nil {
		fmt.Printf("Error executing rsync command: %v\n", err)
	}
	fmt.Println("--------------------------------------------------")
}

func runChown(args []string) {
	var user, group, target, sshHost string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--chown=") {
			chownArgs := strings.SplitN(arg[len("--chown="):], ":", 3)
			if len(chownArgs) == 3 {
				user = chownArgs[0]
				group = chownArgs[1]
				target = sanitizePath(chownArgs[2])
				sshHost = strings.Split(args[len(args)-1], ":")[0] // Extract SSH host from the last arg
				break
			} else {
				fmt.Println("Error: Incorrect format for --chown option. Please specify user:group:target_directory.")
				return
			}
		}
	}

	if user != "" && group != "" && target != "" {
		fmt.Println("Running chown...")
		chownCommand := fmt.Sprintf("sudo chown -R %s:%s %s", user, group, target)
		sshCmd := exec.Command("ssh", sshHost, chownCommand)
		fmt.Println("Chown command:", chownCommand)
		fmt.Println("SSH command:", sshCmd.String())
		output, err := sshCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing chown command via SSH: %v\n", err)
			fmt.Printf("Output: %s\n", string(output))
		}
		fmt.Println("--------------------------------------------------")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: [rsync options] [file] [destination]")
		return
	}

	fmt.Println("This program synchronizes files using rsync and optionally changes ownership using chown on the remote host.")
	fmt.Println("Usage of --chown option:")
	fmt.Println("  --chown=user:group:target_directory")
	fmt.Println("  This option specifies the user and group ownership to be applied on the target directory at the remote host.")
	fmt.Println("  It is important to note that this option is handled internally and is not passed to rsync.")

	runRsync(os.Args[1:])
	runChown(os.Args[1:])
}
