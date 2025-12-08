/*
 * Anasazi Precision Engineering LLC CONFIDENTIAL
 *
 * Unpublished Copyright (c) 2025 Anasazi Precision Engineering LLC. All Rights Reserved.
 *
 * Proprietary to Anasazi Precision Engineering LLC and may be covered by patents, patents
 * in process, and trade secret or copyright law. Dissemination of this information or
 * reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Anasazi Precision Engineering LLC.
 */
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
