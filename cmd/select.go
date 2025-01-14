package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/thuongtruong109/gouse"
	"golang.org/x/exp/slices"
)

type option struct {
	Name  string
	Items []map[int8]string
}

var options = []option{
	{
		Name: "album",
		Items: []map[int8]string{
			{1: "Create new album"},
			{2: "Get all albums"},
			{3: "Get album info by id"},
			{4: "Delete album by id"},
			{5: "Update album by id"},

			{6: "Add track to album"},
			{7: "Get all tracks of album"},
			{8: "Delete track from album"},
		},
	},
	{
		Name: "artist",
		Items: []map[int8]string{
			{9: "Create new artist"},
			{10: "Get all artists"},
			{11: "Get artist by id"},
			// {12: "Get all albums of artist"},
			// {13: "Get all songs of artist"},
			{12: "Delete artist by id"},
			{13: "Update artist by id"},
		},
	},
	{
		Name: "genre",
		Items: []map[int8]string{
			{14: "Create new genre"},
			{15: "Get all genres"},
			{16: "Get genre by id"},
			{17: "Delete genre by id"},
			{18: "Update genre by id"},
		},
	},
	{
		Name: "track",
		Items: []map[int8]string{
			{19: "Create new track"},
			{20: "Get all tracks"},
			{21: "Get track by id"},
			{22: "Delete track by id"},
			{23: "Update track by id"},
		},
	},
	{
		Name: "playlist",
		Items: []map[int8]string{
			{24: "Create new playlist"},
			{25: "Get all playlists"},
			{26: "Get playlist by id"},
			{27: "Delete playlist by id"},
			{28: "Update playlist by id"},

			{29: "Add track to playlist"},
			{30: "Get all tracks of playlist"},
			{31: "Delete track of playlist"},
		},
	},
}

func (d *Delivery) Run() int8 {
	var scopes []string

	for _, option := range options {
		scopes = append(scopes, fmt.Sprintf("%s%v", gouse.SCOPE_SYM, option.Name))
	}

	var qs = []*survey.Question{
		{
			Name: "scope",
			Prompt: &survey.Select{
				Message: "Choose a scope that you want:",
				Options: scopes,
			},
		},
	}

	answers := struct {
		Scopes string `survey:"scope"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	type Choice struct {
		Key   int8
		Value string
	}

	var choices []Choice
	var choiceValue []string

	for _, option := range options {
		if strings.Contains(answers.Scopes, option.Name) {
			for _, item := range option.Items {
				for k, v := range item {
					choiceValue = append(choiceValue, gouse.DOT_SYM+v)
					getResult := &Choice{
						Key:   k,
						Value: v,
					}
					choices = append(choices, *getResult)
				}
			}
			break
		}
	}

	qs2 := []*survey.Question{
		{
			Name: "options",
			Prompt: &survey.Select{
				Message: "Choose a options:",
				Options: choiceValue,
			},
		},
	}

	answers2 := struct {
		Options string `survey:"options"`
	}{}

	err2 := survey.Ask(qs2, &answers2)
	if err2 != nil {
		fmt.Println(err2.Error())
		return -1
	}

	idx := slices.IndexFunc(choices, func(c Choice) bool {
		return strings.Contains(answers2.Options, c.Value)
	})

	return int8(choices[idx].Key)
}
