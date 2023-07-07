package newspace

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const (
	main   = "main"
	master = "master"
)

var ErrSpaceExists = errors.New("space already exists")

func (m Model) Clone() (string, error) {
	branch := m.Branch
	if branch == "" {
		branch = main
	}

	spaceName, err := m.prepareSpace(branch)
	if err != nil {
		return "", fmt.Errorf("error preparing space: %w", err)
	}

	// try cloning to branch
	_, err = m.tryClone(spaceName, branch)

	if err != nil {
		rep, err := m.tryCloneDefaultBranch(spaceName)
		if err != nil {
			return spaceName, fmt.Errorf("error cloning default branch: %w", err)
		}

		err = m.checkoutBranch(rep, branch)
		if err != nil {
			return spaceName, fmt.Errorf("error checking out branch: %w", err)
		}
	}

	return spaceName, nil
}

func (m Model) checkoutBranch(rep *git.Repository, branch string) error {
	err := rep.CreateBranch(&config.Branch{
		Name: branch,
	})
	if err != nil {
		return fmt.Errorf("error creating branch: %w", err)
	}

	wt, err := rep.Worktree()
	if err != nil {
		return fmt.Errorf("error generating worktree: %w", err)
	}
	return wt.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
		Create: true,
	})
}

func (m Model) tryCloneDefaultBranch(spacename string) (rep *git.Repository, err error) {
	return git.PlainClone(spacename, false, &git.CloneOptions{
		URL: m.Repo,
		Auth: &http.BasicAuth{
			Username: m.Conf.GithubUsername,
			Password: m.Conf.GithubToken,
		},
	})
}

func (m Model) tryClone(spaceName, branch string) (*git.Repository, error) {
	return git.PlainClone(spaceName, false, &git.CloneOptions{
		URL:           m.Repo,
		ReferenceName: plumbing.ReferenceName(branch),
		Auth: &http.BasicAuth{
			Username: m.Conf.GithubUsername,
			Password: m.Conf.GithubToken,
		},
	})
}

func (m Model) prepareSpace(branch string) (string, error) {
	parts := strings.Split(m.Repo, "/")
	repoName := parts[len(parts)-1]
	repoName = strings.Replace(repoName, ".git", "", 1)
	ownerName := parts[len(parts)-2]

	return fmt.Sprintf("%s/%s/%s-%s", m.Conf.SpacesDirectory, ownerName, repoName, branch), nil
}
