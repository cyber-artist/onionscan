package protocol

import (
	"fmt"

	"github.com/cyber-artist/onionscan/config"
	"github.com/cyber-artist/onionscan/report"

	"net/http"

	"github.com/cyber-artist/onionscan/spider"
	"github.com/cyber-artist/onionscan/utils"
)

type HTTPProtocolScanner struct {
	Client *http.Client
}

func (hps *HTTPProtocolScanner) ScanProtocol(hiddenService string, osc *config.OnionScanConfig, report *report.OnionScanReport) {

	// HTTP
	osc.LogInfo(fmt.Sprintf("Checking %s http(80)\n", hiddenService))
	conn, err := utils.GetNetworkConnection(hiddenService, 80, osc.TorProxyAddress, osc.Timeout)
	if conn != nil {
		conn.Close()
	}
	if err != nil {
		osc.LogInfo("Failed to connect to service on port 80\n")
		report.WebDetected = false
	} else {
		osc.LogInfo("Found potential service on http(80)\n")
		report.WebDetected = true
		wps := new(spider.OnionSpider)
		wps.Crawl(report.HiddenService, osc, report)
	}
}
