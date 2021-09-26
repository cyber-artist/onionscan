package protocol

import (
	"github.com/cyber-artist/onionscan/config"
	"github.com/cyber-artist/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
