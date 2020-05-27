// This package is a copy of vsc code in package golang.org/x/tools/go/vcs,
// modified to custom git command.

package vcs

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Verbose enables verbose operation logging.
var Verbose bool

// ShowCmd controls whether VCS commands are printed.
var ShowCmd bool

// A Cmd describes how to use a version control system
// like Mercurial, Git, or Subversion.
type Cmd struct {
	Name string
	Cmd  string // name of binary to invoke command

	InitCmd      string // command to init a fresh repository
	AddAllCmd    string // command to add all files into an existing repository
	CommitAllCmd string // command to commit all files into an existing repository
	DiffStatCmd  string // command to get diff in repository

	LogCmd string // command to list repository changelogs in an XML format

	Scheme  []string
	PingCmd string
}

// vcsList lists the known version control systems
var vcsList = []*Cmd{
	vcsGit,
}

// ByCmd returns the version control system for the given
// command name (hg, git, svn, bzr).
func ByCmd(cmd string) *Cmd {
	for _, vcs := range vcsList {
		if vcs.Cmd == cmd {
			return vcs
		}
	}
	return nil
}

// vcsGit describes how to use Git.
var vcsGit = &Cmd{
	Name: "Git",
	Cmd:  "git",

	InitCmd:      "init {dir}",
	AddAllCmd:    "add -A",
	CommitAllCmd: "commit -a -m {message}",
	DiffStatCmd:  "diff --numstat --shortstat {start} {end}",

	Scheme: []string{"git", "https", "http", "git+ssh"},
}

func (v *Cmd) String() string {
	return v.Name
}

// run runs the command line cmd in the given directory.
// keyval is a list of key, value pairs.  run expands
// instances of {key} in cmd into value, but only after
// splitting cmd into individual arguments.
// If an error occurs, run prints the command line and the
// command's combined stdout+stderr to standard error.
// Otherwise run discards the command's output.
func (v *Cmd) run(dir string, cmd string, keyval ...string) error {
	_, err := v.run1(dir, cmd, keyval, true)
	return err
}

// runVerboseOnly is like run but only generates error output to standard error in verbose mode.
func (v *Cmd) runVerboseOnly(dir string, cmd string, keyval ...string) error {
	_, err := v.run1(dir, cmd, keyval, false)
	return err
}

// runOutput is like run but returns the output of the command.
func (v *Cmd) runOutput(dir string, cmd string, keyval ...string) ([]byte, error) {
	return v.run1(dir, cmd, keyval, true)
}

// run1 is the generalized implementation of run and runOutput.
func (v *Cmd) run1(dir string, cmdline string, keyval []string, verbose bool) ([]byte, error) {
	m := make(map[string]string)
	for i := 0; i < len(keyval); i += 2 {
		m[keyval[i]] = keyval[i+1]
	}
	args := strings.Fields(cmdline)
	for i, arg := range args {
		args[i] = expand(m, arg)
	}

	_, err := exec.LookPath(v.Cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Missing %s command.\n",
			v.Name)
		return nil, err
	}

	cmd := exec.Command(v.Cmd, args...)
	cmd.Dir = dir
	cmd.Env = envForDir(cmd.Dir)
	if ShowCmd {
		fmt.Printf("cd %s\n", dir)
		fmt.Printf("%s %s\n", v.Cmd, strings.Join(args, " "))
	}
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err = cmd.Run()
	out := buf.Bytes()
	if err != nil {
		if verbose || Verbose {
			fmt.Fprintf(os.Stderr, "# cd %s; %s %s\n", dir, v.Cmd, strings.Join(args, " "))
			os.Stderr.Write(out)
		}
		return nil, err
	}
	return out, nil
}

// Init create an empty Git repository or reinitialize an existing one
func (v *Cmd) Init(dir string) error {
	return v.run(".", v.InitCmd, "dir", dir)
}

// AddAll Add file contents to the index
func (v *Cmd) AddAll(dir string) error {
	return v.runVerboseOnly(dir, v.AddAllCmd)
}

// CommitAll Record changes to the repository
func (v *Cmd) CommitAll(dir string) error {
	return v.runVerboseOnly(dir, v.CommitAllCmd, "message", "auto commit")
}

// DiffStat Show changes between commits, commit and working tree, etc
func (v *Cmd) DiffStat(dir string, start, end string) ([]byte, error) {
	if len(start) != 0 {
		start = "HEAD@{" + start + "}"
	} else {
		start = "HEAD"
	}
	if len(end) != 0 {
		end = "HEAD@{" + end + "}"
	} else {
		end = "HEAD"
	}

	out, err := v.runOutput(dir, v.DiffStatCmd, "start", start, "end", end)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FromDir inspects dir and its parents to determine the
// version control system and code repository to use.
// On return, root is the import path
// corresponding to the root of the repository.
func FromDir(dir, srcRoot string) (vcs *Cmd, root string, err error) {
	// Clean and double-check that dir is in (a subdirectory of) srcRoot.
	dir = filepath.Clean(dir)
	srcRoot = filepath.Clean(srcRoot)
	if len(dir) <= len(srcRoot) || dir[len(srcRoot)] != filepath.Separator {
		return nil, "", fmt.Errorf("directory %q is outside source root %q", dir, srcRoot)
	}

	var vcsRet *Cmd
	var rootRet string

	origDir := dir
	for len(dir) > len(srcRoot) {
		for _, vcs := range vcsList {
			if _, err := os.Stat(filepath.Join(dir, "."+vcs.Cmd)); err == nil {
				root := filepath.ToSlash(dir[len(srcRoot)+1:])
				// Record first VCS we find, but keep looking,
				// to detect mistakes like one kind of VCS inside another.
				if vcsRet == nil {
					vcsRet = vcs
					rootRet = root
					continue
				}
				// Allow .git inside .git, which can arise due to submodules.
				if vcsRet == vcs && vcs.Cmd == "git" {
					continue
				}
				// Otherwise, we have one VCS inside a different VCS.
				return nil, "", fmt.Errorf("directory %q uses %s, but parent %q uses %s",
					filepath.Join(srcRoot, rootRet), vcsRet.Cmd, filepath.Join(srcRoot, root), vcs.Cmd)
			}
		}

		// Move to parent.
		ndir := filepath.Dir(dir)
		if len(ndir) >= len(dir) {
			// Shouldn't happen, but just in case, stop.
			break
		}
		dir = ndir
	}

	if vcsRet != nil {
		return vcsRet, rootRet, nil
	}

	return nil, "", fmt.Errorf("directory %q is not using a known version control system", origDir)
}

// RepoRoot represents a version control system, a repo, and a root of
// where to put it on disk.
type RepoRoot struct {
	VCS *Cmd

	// Repo is the repository URL, including scheme.
	Repo string

	// Root is the import path corresponding to the root of the
	// repository.
	Root string
}

// validateRepoRoot returns an error if repoRoot does not seem to be
// a valid URL with scheme.
func validateRepoRoot(repoRoot string) error {
	url, err := url.Parse(repoRoot)
	if err != nil {
		return err
	}
	if url.Scheme == "" {
		return errors.New("no scheme")
	}
	return nil
}

// errNoMatch is returned from matchGoImport when there's no applicable match.
var errNoMatch = errors.New("no import match")

// pathPrefix reports whether sub is a prefix of s,
// only considering entire path components.
func pathPrefix(s, sub string) bool {
	// strings.HasPrefix is necessary but not sufficient.
	if !strings.HasPrefix(s, sub) {
		return false
	}
	// The remainder after the prefix must either be empty or start with a slash.
	rem := s[len(sub):]
	return rem == "" || rem[0] == '/'
}

// expand rewrites s to replace {k} with match[k] for each key k in match.
func expand(match map[string]string, s string) string {
	for k, v := range match {
		s = strings.Replace(s, "{"+k+"}", v, -1)
	}
	return s
}

// noVCSSuffix checks that the repository name does not
// end in .foo for any version control system foo.
// The usual culprit is ".git".
func noVCSSuffix(match map[string]string) error {
	repo := match["repo"]
	for _, vcs := range vcsList {
		if strings.HasSuffix(repo, "."+vcs.Cmd) {
			return fmt.Errorf("invalid version control suffix in %s path", match["prefix"])
		}
	}
	return nil
}
