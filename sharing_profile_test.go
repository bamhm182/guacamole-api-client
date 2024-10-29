//go: build all || unittests || specific
// +build all unittests specific

package guacamole

import (
	"fmt"
	"os"
	"strings"
	"testing"
	
	"github.com/bamhm182/guacamole-api-client/types"
)

var (
	sharingProfilesConfig = Config{
        URL:                    os.Getenv("GUACAMOLE_URL"),
        Username:               os.Getenv("GUACAMOLE_USERNAME"),
        Password:               os.Getenv("GUACAMOLE_PASSWORD"),
        Token:                  os.Getenv("GUACAMOLE_TOKEN"),
        DataSource:             os.Getenv("GUACAMOLE_DATA_SOURCE"),
        DisableTLSVerification: true,
    }
    testSharingProfile = types.GuacSharingProfile{
        Name:                        "Test Sharing Profile",
        PrimaryConnectionIdentifier: "1592",
    }
)

func TestListSharingProfiles(t *testing.T) {
	if os.Getenv("GUACAMOLE_COOKIES") != "" {
        sharingProfilesConfig.Cookies = make(map[string]string)
        for _, e := range strings.Split(os.Getenv("GUACAMOLE_COOKIES"), ",") {
            cookie_split := strings.Split(e, "=")
            sharingProfilesConfig.Cookies[cookie_split[0]] = cookie_split[1]
        }
    }
    client := New(sharingProfilesConfig)

    err := client.Connect()
    if err != nil {
        t.Errorf("Error %s connecting to guacamole with config %+v", err, sharingProfilesConfig)
    }

    _, err = client.ListSharingProfiles()
    if err != nil {
        t.Errorf("Error %s listing sharing profiles with client %+v", err, client)
    }
}

func TestCreateSharingProfile(t *testing.T) {
	client := New(sharingProfilesConfig)

	err := client.Connect()
	if err != nil {
        t.Errorf("Error %s connecting to guacamole with config %+v", err, sharingProfilesConfig)
	}

	err = client.CreateSharingProfile(&testSharingProfile)
	if err != nil {
		t.Errorf("Error %s creating sharing profile: %s with client %+v", err, testSharingProfile.Name, client)
	}
}

func TestReadSharingProfile(t *testing.T) {
	client := New(sharingProfilesConfig)

	err := client.Connect()
	if err != nil {
        t.Errorf("Error %s connecting to guacamole with config %+v", err, sharingProfilesConfig)
	}

	err = client.ReadSharingProfile(fmt.Sprintf("%s", testSharingProfile.Identifier))
	if err != nil {
		t.Errorf("Error %s reading sharing profile: %s with client %+v", err, testSharingProfile.Name, client)
	}
}

func TestUpdateSharingProfile(t *testing.T) {
	client := New(sharingProfilesConfig)

	err := client.Connect()
	if err != nil {
        t.Errorf("Error %s connecting to guacamole with config %+v", err, sharingProfilesConfig)
	}

	testSharingProfile.Name = "New Name"
	testSharingProfile.Parameters.ReadOnly = "true"

	err = client.UpdateSharingProfile(&testSharingProfile)
	if err != nil {
		t.Errorf("Error %s updating sharing profile: %s with client %+v", err, testSharingProfile.Name, client)
	}
}

func TestDeleteSharingProfile(t *testing.T) {
	client := New(sharingProfilesConfig)

	err := client.Connect()
	if err != nil {
        t.Errorf("Error %s connecting to guacamole with config %+v", err, sharingProfilesConfig)
	}

	err = client.DeleteSharingProfile(testSharingProfile.Identifier)
	if err != nil {
		t.Errorf("Error %s deleting sharing profile: %s with client %+v", err, testSharingProfile.Name, client)
	}
}
