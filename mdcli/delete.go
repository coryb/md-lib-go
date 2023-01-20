package mdcli

import (
	"os"
	"path/filepath"

	mdlib "github.com/spinnaker/md-lib-go"
)

// Delete is a command line interface for removing the management
// of a delivery config.  All management history will be removed.
func Delete(opts *CommandOptions) error {
	configPath := filepath.Join(opts.ConfigDir, opts.ConfigFile)
	if _, err := os.Stat(configPath); err != nil {
		return err
	}

	cli := mdlib.NewClient(
		mdlib.WithBaseURL(opts.BaseURL),
		mdlib.WithHTTPClient(opts.HTTPClient),
	)

	mdProcessor := mdlib.NewDeliveryConfigProcessor(
		mdlib.WithDirectory(opts.ConfigDir),
		mdlib.WithFile(opts.ConfigFile),
		mdlib.WithLogger(opts.Logger),
	)

	err := mdProcessor.Delete(cli)
	if err != nil {
		return err
	}

	opts.Logger.Noticef("OK")
	return nil
}
