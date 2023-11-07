//go:build !test && !secretcli

package types

// create func Asset
func Asset(name string) ([]byte, error) {
	return []byte("random"), nil
}

func GetApiKey() ([]byte, error) {
	apiKeyFile, err := Asset("api_key.txt")
	if err != nil {
		return nil, err
	}

	return apiKeyFile, nil
}

func GetSpid() ([]byte, error) {
	apiKeyFile, err := Asset("spid.txt")
	if err != nil {
		return nil, err
	}

	return apiKeyFile, nil
}
