package segments

import (
	"errors"
	"fmt"	
	"strings"
  	
	"github.com/jandedobbeleer/oh-my-posh/src/platform"
	"github.com/jandedobbeleer/oh-my-posh/src/properties"
)

type Godot struct {
    props   properties.Properties
    env     platform.Environment

    GodotVersion string
}

const (
    //NewProp enables something
    NewProp properties.Property = "newprop"
)

func (g *Godot) Enabled() bool {
	godotVersion, err := g.GetGodotVersion()
    return godotVersion != "" && err == nil
}

func (g *Godot) GetGodotVersion() (version string, err error) {
	projectFile, err := g.env.HasParentFilePath("project.godot")
	if err != nil {
		g.env.Debug("No project.godot file found")
		return
	}

	projectFilePath := projectFile.Path
	projectFileContent := g.env.FileContent(projectFilePath)
	firstLine := strings.Split(projectFileContent, "\n")[0]

	versionKey := "config_version = "
	versionKeyIndex := strings.Index(firstLine, versionKey)

	if versionKeyIndex == -1 {
		err := errors.New("project.godot is missing the 'config_version' key")
		return "", err
	}

	versionValueIndex := versionKeyIndex + len(versionKey)
	configVersion := firstLine[versionValueIndex:]
	configVersion = strings.TrimSpace(configVersion)

	var godotVersionsByConfigVersion = map[string]string {
		"4": "3.x",
		"5": "4.x",
	}

	godotVersion, found := godotVersionsByConfigVersion[configVersion]

	if found {
		return godotVersion, nil
	}

	g.env.Debug(fmt.Sprintf("Godot version %s is not supported", configVersion))
	err = errors.New("Godot version for this project is not supported")
	return "", err 
}

func (g *Godot) Template() string {
    return " {{.Text}} world "
}

func (g *Godot) Init(props properties.Properties, env platform.Environment) {
    g.props = props
    g.env = env
}