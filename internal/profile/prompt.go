package profile

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/gabe565/gh-profile/internal/github"
	"os"
	"path/filepath"
)

func Prompt() (Profile, error) {
	conf := github.ConfigDir()

	files, err := os.ReadDir(filepath.Join(conf, "accounts"))
	if err != nil {
		return Profile{}, err
	}

	profiles := []string{"New Profile"}
	for _, file := range files {
		if file.IsDir() {
			profiles = append(profiles, file.Name())
		}
	}

	var answer string
	if err := survey.AskOne(&survey.Select{
		Message: "Choose a profile",
		Options: profiles,
	}, &answer, survey.WithValidator(survey.Required)); err != nil {
		return Profile{}, err
	}

	if answer == "New Profile" {
		if err := survey.AskOne(&survey.Input{
			Message: "Enter new profile name",
		}, &answer, survey.WithValidator(survey.Required)); err != nil {
			return Profile{}, err
		}
	}

	return New(answer), nil
}
