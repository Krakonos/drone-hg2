package main

import (
	"fmt"
	"os"
	"os/user"
	"os/exec"
	"path/filepath"
	"io/ioutil"
)

type (
	Repo struct {
		Owner   string
		Name    string
		Link    string
		Avatar  string
		Branch  string
		Private bool
		Trusted bool
	}

	Build struct {
		Number   int
		Workspace string
		Event    string
		Status   string
		Deploy   string
		Created  int64
		Started  int64
		Finished int64
		Link     string
	}

	Commit struct {
		Remote  string
		Sha     string
		Ref     string
		Link    string
		Branch  string
		Message string
		Author  Author
	}

	Author struct {
		Name   string
		Email  string
		Avatar string
	}

	Config struct {
		// plugin-specific parameters and secrets
	}

	Netrc struct {
		Machine  string
		Username    string
		Password string

	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
		Netrc  Netrc
	}
)

func (p Plugin) Exec() error {
	// plugin logic goes here
	fmt.Println("Arguments are: ", p);

	// First, write hgrc for mercurial
	err := writeHgrc(p.Netrc.Machine, p.Netrc.Username, p.Netrc.Password)
	if err != nil {
		return err
	}

	// Next write the netrc for git (in case we have some git subrepositories
	err = writeNetrc(p.Netrc.Machine, p.Netrc.Username, p.Netrc.Password)
	if err != nil {
		return err
	}

	// This array will store a sequence of commands to be executed
	var cmds []*exec.Cmd

	cmds = append(cmds, exec.Command("hg", "clone",   p.Commit.Remote, "-r", p.Commit.Sha, p.Build.Workspace));
	cmds = append(cmds, exec.Command("hg", "update", p.Commit.Sha));

	// Execute all the commands.
	for _, cmd := range cmds {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		trace(cmd)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Command failed!")
			return err
		} else {
			fmt.Println("Command successful!")
		}
	}

	fmt.Println("Cloning done!")

	return nil
}

var hgrcFile = `
[auth]
drone.prefix = %s
drone.username = %s
drone.password = %s
drone.schemes = https
`

func writeHgrc(Machine string, Username string, Password string) error {
	out := fmt.Sprintf(
		hgrcFile,
		Machine, // TODO this may require adding http(s) prefix
		Username,
		Password,
	)
	home := "/root"
	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}

	path := filepath.Join(home, ".hgrc")
	return ioutil.WriteFile(path, []byte(out), 0600)
}
