package guacamole

import (
	"fmt"
	"net/http"
	"net/url"

	 "github.com/bamhm182/guacamole-api-client/types"
)

const (
	sharingProfilesPath = "sharingProfiles"
)

// CreateSharingProfile creats a guacamole sharing profile
func (c *Client) CreateSharingProfile(sharingProfile *types.GuacSharingProfile) error {
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, sharingProfilesPath), sharingProfile

	if err != nil {
		return err
	}

	err = c.Call(request, &sharingProfile)

	if err != nil {
		return err
	}

	return nil
}

// ReadSharingProfile gets a sharing profile by identifier
func (c *Client) ReadSharingProfile(identifier string) (types.GuacSharingProfile) error {
	var ret types.GuacSharingProfile
	var retParams types.GuacSharingProfileParameters

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, sharingProfilesPath, url.QueryEscape(identifier)), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	if ret.Identifier != "" {
		request, err = c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/parameters", c.baseURL, sharingProfilesPath, identifier), nil)

		if err != nil {
			return ret, err
		}

		err = c.Call(request, &retParams)
		if err != nil {
			return ret, err
		}
	}

	ret.Parameters = retParams

	return ret, nil
}

// UpdateSharingProfile updates a sharing profile by identifier
func (c *Client) UpdateSharingProfile(sharingProfile *types.GuacSharingProfile) error {
    request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s", c.baseURL, sharingProfilesPath, url.QueryEscape(sharingProfile.identifier)), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSharingProfile deletes a sharing profile by identifier
func (c *Client) DeleteSharingProfile(identifier string) error {
    request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, sharingProfilesPath, url.QueryEscape(identifier)), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}

	return nil

}

// ListSharingProfiles lists all sharing profiles
func (c *Client) ListSharingProfiles() ([]types.GuacSharingProfile, err) {
	var ret []types.GuacSharingProfile
	var sharingProfileList map[string]types.GuacSharingProfile

    request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, sharingProfilesPath), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &sharingProfileList)
	if err != nil {
		return ret, err
	}

	for _, sharingProfile := range sharingProfileList {
		ret = append(ret, sharingProfile)
	}

	return ret, nil
}
