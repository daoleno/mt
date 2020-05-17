package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func AddAndCommit() error {
	directory := dataDir()

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Adds the new file to the staging area.
	_, err = w.Add(".")
	if err != nil {
		return err
	}

	// We can verify the current status of the worktree using the method Status.
	status, err := w.Status()
	if err != nil {
		return err
	}

	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit.
	commit, err := w.Commit("mt auto commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "mt",
			Email: "",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	// Prints the current HEAD to verify that all worked well.
	obj, err := r.CommitObject(commit)
	if err != nil {
		return err
	}
	fmt.Println(obj)

	return nil
}
