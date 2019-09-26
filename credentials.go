package vcodeHMAC

import (
	"bufio"
	"os"
	"strings"
)

func getCredentials(fileString string) ([2]string, error) {
	var credentials [2]string

	// Modification by securityRelic@github
	// Here's my quick hack for use in docker pipeline tool (e.g., bitbucket)
	// and secure the secure.  If "fileString" is empty, look for OS environment
	// variable else error.  Need to find a way to may this a bit cleaner!

	if fileString == "" {

		key, ok := os.LookupEnv("VERACODE_API_KEY_ID")
		if ok == false {
			return credentials, errors.New("missing VERACODE_API_KEY_ID")
		}
		shhh, ok := os.LookupEnv("VERACODE_API_KEY_SECRET")
		if ok == false {
			return credentials, errors.New("missing VERACODE_API_KEY_SECRET")
		}

		credentials[0], credentials[1] = key, shhh
		return credentials, nil

	}

	file, err := os.Open(fileString)
	if err != nil {
		return credentials, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//We remove spaces to account for discrepancies in user configuration of creds file
		if strings.Contains(scanner.Text(), "veracode_api_key_id") {
			removeSpaces := strings.Replace(scanner.Text(), " ", "", -1)
			credentials[0] = strings.Replace(removeSpaces, "veracode_api_key_id=", "", -1)
		} else if strings.Contains(scanner.Text(), "veracode_api_key_secret") {
			removeSpaces := strings.Replace(scanner.Text(), " ", "", -1)
			credentials[1] = strings.Replace(removeSpaces, "veracode_api_key_secret=", "", -1)
		}
	}

	if err := scanner.Err(); err != nil {
		return credentials, err
	}

	return credentials, nil
}
