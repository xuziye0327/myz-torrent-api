package aria2

import (
	"encoding/json"
	"fmt"
	"myz-torrent-api/common"
)

// Aria2 rpc
type Aria2 struct {
	secret    string
	rpcClient *common.RPCClient
}

// NewAria2 creates an Aria2 client without secret
func NewAria2(url string) *Aria2 {
	return NewAria2WithSecret(url, "")
}

// NewAria2WithSecret creates an Aria2 client with secret
func NewAria2WithSecret(url, secret string) *Aria2 {
	if len(secret) > 0 {
		secret = fmt.Sprintf("token:%s", secret)
	}
	return &Aria2{
		secret:    secret,
		rpcClient: common.NewRPCClient(url),
	}
}

// AddURI adds a new download.
// uris is an array of HTTP/FTP/SFTP/BitTorrent URIs (strings) pointing to the same resource.
func (a *Aria2) AddURI(uris ...string) error {
	return a.handlerWithFixParams("aria2.addUri", nil, uris)
}

// TellStatus returns the progress of the download denoted by gid (string).
func (a *Aria2) TellStatus(gid string) (stat Status, err error) {
	return stat, a.handler("aria2.tellStatus", &stat, gid)
}

// TellAllJob returns a list of all downloads.
func (a *Aria2) TellAllJob(numWaiting, numStopped int) (stats []Status, err error) {
	methods := []Method{
		{
			MethodName: "aria2.tellActive",
			Params:     a.secretParams(),
		},
		{
			MethodName: "aria2.tellWaiting",
			Params:     a.secretParams(0, numWaiting),
		},
		{
			MethodName: "aria2.tellStopped",
			Params:     a.secretParams(0, numStopped),
		},
	}
	var allStats [][][]Status
	if err := a.handler("system.multicall", &allStats, methods); err != nil {
		return nil, err
	}
	for _, ss := range allStats {
		for _, s := range ss {
			stats = append(stats, s...)
		}
	}
	return stats, nil
}

// TellActive returns a list of active downloads.
func (a *Aria2) TellActive() (stats []Status, err error) {
	return stats, a.handlerWithFixParams("aria2.tellActive", &stats)
}

// TellWaiting returns a list of waiting downloads, including paused ones.
func (a *Aria2) TellWaiting(offset, num int) (stats []Status, err error) {
	return stats, a.handlerWithFixParams("aria2.tellWaiting", &stats, offset, num)
}

// TellStopped returns a list of stopped downloads.
func (a *Aria2) TellStopped(offset, num int) (stats []Status, err error) {
	return stats, a.handlerWithFixParams("aria2.tellStopped", &stats, offset, num)
}

// GetGlobalStat returns global statistics such as the overall download and upload speeds.
func (a *Aria2) GetGlobalStat() (stat GlobalStatistics, err error) {
	return stat, a.handlerWithFixParams("aria2.getGlobalStat", &stat)
}

// GetVersion returns the version of aria2 and the list of enabled features.
func (a *Aria2) GetVersion() error {
	return a.handlerWithFixParams("aria2.getVersion", nil)
}

func (a *Aria2) handler(method string, ret interface{}, params ...interface{}) error {
	if resp, err := a.rpcClient.Call(method, params...); err != nil {
		return err
	} else if resp.Error != nil {
		return fmt.Errorf("%s error: %s", method, common.JsonToString(resp.Error))
	} else if ret != nil {
		bs, _ := json.Marshal(resp.Result)
		json.Unmarshal(bs, &ret)
	}

	return nil
}

func (a *Aria2) handlerWithFixParams(method string, ret interface{}, params ...interface{}) error {
	return a.handler(method, ret, a.secretParams(params...)...)
}

func (a *Aria2) secretParams(params ...interface{}) []interface{} {
	if len(a.secret) > 0 {
		return append([]interface{}{a.secret}, params...)
	}
	return params
}
