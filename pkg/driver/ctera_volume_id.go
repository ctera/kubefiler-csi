package driver

import (
	"encoding/json"
	"errors"
)

type CteraVolumeID struct {
	FilerAddress string `json:"filer_address"`
	ShareName    string `json:"share_name"`
	Path         string `json:"path"`
}

func (c *CteraVolumeID) ToVolumeID() (*string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	ret := string(bytes)
	return &ret, nil
}

func getCteraVolumeIDFromVolumeID(volumeID string) (*CteraVolumeID, error) {
	if len(volumeID) == 0 {
		return nil, errors.New("volume ID missing in request")
	}

	var cteraVolumeID CteraVolumeID
	err := json.Unmarshal([]byte(volumeID), &cteraVolumeID)
	if err != nil {
		return nil, err
	}
	return &cteraVolumeID, nil
}
