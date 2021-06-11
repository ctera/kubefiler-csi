package driver

import (
	"encoding/json"
	"errors"
)

type CteraVolumeId struct {
	FilerAddress string `json:"filer_address"`
	ShareName  	 string `json:"share_name"`
	Path 		 string `json:"path"`
}

func (c *CteraVolumeId) ToVolumeId() (*string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	ret := string(bytes)
	return &ret, nil
}

func getCteraVolumeIdFromVolumeId(volumeId string) (*CteraVolumeId, error){
	if len(volumeId) == 0 {
		return nil, errors.New("Volume ID missing in request")
	}

	var cteraVolumeId CteraVolumeId
	err := json.Unmarshal([]byte(volumeId), &cteraVolumeId)
	if err != nil {
		return nil, err
	}
	return &cteraVolumeId, nil
}
