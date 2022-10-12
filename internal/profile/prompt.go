package profile

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/gabe565/gh-profile/internal/github"
	"os"
	"path/filepath"
)

func List() ([]string, error) {
	conf := github.ConfigDir()

	files, err := os.ReadDir(filepath.Join(conf, "accounts"))
	if err != nil {
		return []string{}, err
	}

	profiles := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			profiles = append(profiles, file.Name())
		}
	}

	return profiles, nil
}

func Prompt() (Profile, error) {
	profiles, err := List()
	if err != nil {
		return Profile{}, err
	}

	var answer string
	if err := survey.AskOne(&survey.Select{
		Message: "Choose a profile",
		Options: profiles,
	}, &answer, survey.WithValidator(survey.Required)); err != nil {
		return Profile{}, err
	}

	return New(answer), nil
}
