package config

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Configuration struct {
	BaseURL      string `json:"base_url"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// // TODO: Create functions to save/load tokens from secure storage
// func LoadTokens() Tokens {
// 	// Placeholder implementation
// 	// On error, assume tokens are not available so create empty Tokens struct

// 	return Tokens{}
// }

// func StoreTokens(tokens Tokens) error {

// 	// Read existing configuration file to get base_url, or should I just use current config?
// 	//var cfg Configuration

// 	// Get config name from Viper
// 	cfgFile := viper.ConfigFileUsed()
// 	if cfgFile == "" {
// 		return errors.New("no configuration file found to store tokens")
// 	}

// 	//fileType := viper.

// 	// Placeholder implementation
// 	return nil
// }
