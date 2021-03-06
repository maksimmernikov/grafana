package commands

import (
	"github.com/fatih/color"
	"github.com/maksimmernikov/grafana/pkg/cmd/grafana-cli/logger"
	s "github.com/maksimmernikov/grafana/pkg/cmd/grafana-cli/services"
)

func upgradeCommand(c CommandLine) error {
	pluginsDir := c.PluginDirectory()
	pluginName := c.Args().First()

	localPlugin, err := s.ReadPlugin(pluginsDir, pluginName)

	if err != nil {
		return err
	}

	v, err2 := s.GetPlugin(pluginName, c.RepoDirectory())

	if err2 != nil {
		return err2
	}

	if ShouldUpgrade(localPlugin.Info.Version, v) {
		s.RemoveInstalledPlugin(pluginsDir, pluginName)
		return InstallPlugin(pluginName, "", c)
	}

	logger.Infof("%s %s is up to date \n", color.GreenString("✔"), pluginName)
	return nil
}
