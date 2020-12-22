package common

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

// Config is a global config using in this proj.
type Config struct {
	ServerAddr   string `string:"server" json:"server"`
	ServerPortal int    `int:"portal" json:"portal"`
	LogPath      string `string:"log_path" json:"log_path"`

	DownloadConfig *DownloadConfig `DownloadConfig:"download_config" json:"download_config"`
}

// DownloadConfig is using in running time
type DownloadConfig struct {
	DownloadDir  string `long:"download-directory" json:"download_directory"`
	Aria2Address string `json:"aria2_address"`
	Aria2Secert  string `json:"aria2_secert"`
}

// LoadConfig will load ~/.myz_torrent_config.json or generate a new config.
func LoadConfig() (*Config, error) {
	var p int
	flag.IntVar(&p, "p", 8080, "")

	var s string
	flag.StringVar(&s, "s", "0.0.0.0", "")

	var d string
	flag.StringVar(&d, "d", "~/myz_torrent_download/", "")

	var c string
	flag.StringVar(&c, "c", "", "")

	var l string
	flag.StringVar(&l, "l", "", "")

	var aria2Address string
	flag.StringVar(&aria2Address, "aria2-address", "", "")

	var aria2Secert string
	flag.StringVar(&aria2Secert, "aria2-secert", "", "")

	flag.Parse()

	var conf = &Config{
		ServerAddr:   s,
		ServerPortal: p,
		LogPath:      l,

		DownloadConfig: &DownloadConfig{
			DownloadDir:  d,
			Aria2Address: aria2Address,
			Aria2Secert:  aria2Secert,
		},
	}

	if len(c) > 0 {
		if bs, err := ioutil.ReadFile(c); err != nil {
			return nil, err
		} else if err := json.Unmarshal(bs, conf); err != nil {
			return nil, err
		}
	}

	return conf, nil
}
