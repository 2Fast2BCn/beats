package status

import (
	"io"
	"encoding/json"

	"github.com/elastic/beats/libbeat/common"
	"io/ioutil"
	"github.com/elastic/beats/libbeat/logp"
)

// Map body to MapStr
func eventMapping(m *MetricSet, body io.ReadCloser, hostname string, metricset string) (common.MapStr, error) {
	// Nginx upstream status sample:
	// {"servers": {
	//	"total": 4,
	//	"generation": 1,
	//	"server": [
	//		{"index": 0, "upstream": "websphere", "name": "10.60.42.8:9080", "status": "up", "rise": 366, "fall": 0, "type": "http", "port": 0, "since": "Fri, 01 Sep 2017 13:30:25 UTC", "elapse": 6},
	//		{"index": 1, "upstream": "websphere", "name": "10.60.42.8:9081", "status": "up", "rise": 2560, "fall": 0, "type": "http", "port": 0, "since": "Fri, 01 Sep 2017 13:30:25 UTC", "elapse": 6},
	//		{"index": 2, "upstream": "websphere", "name": "10.60.42.9:9080", "status": "up", "rise": 1861, "fall": 0, "type": "http", "port": 0, "since": "Fri, 01 Sep 2017 13:30:25 UTC", "elapse": 6},
	//		{"index": 3, "upstream": "websphere", "name": "10.60.42.9:9081", "status": "up", "rise": 764, "fall": 0, "type": "http", "port": 0, "since": "Fri, 01 Sep 2017 13:30:25 UTC", "elapse": 6}
	//	]
	// }}
	//
	type Server struct {
		Index   int         `json:"index"`
		Upstream string     `json:"upstream"`
		Name	 string     `json:"name"`
		Status	 string     `json:"status"`
		Rise	 int        `json:"rise"`
		Fall	 int        `json:"fall"`
		Type	 string     `json:"type"`
		Port	 int        `json:"port"`
		Since	 int        `json:"since"`
		Elapse	 int        `json:"elapse"`
	}

	type Servers struct {
		Total   	int     	`json:"total"`
		Up			int     	`json:"up"`
		Down		int     	`json:"down"`
		Generation	int			`json:"generation"`
		Server		[]Server	`json:"server"`
	}

	type Response struct {
		Servers	Servers    `json:"servers"`
	}

	dat := common.MapStr{}
	byteBody, err := ioutil.ReadAll(body)
	if err := json.Unmarshal([]byte(byteBody), &dat); err != nil {
		panic(err)
	}
	if err != nil {
		logp.Err("Error closing: %v", err)
	}

	return dat, nil
}
