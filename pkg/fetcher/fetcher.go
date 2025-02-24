package fetcher

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"git.cryptic.systems/volker.raschek/dcmerge/pkg/domain/dockerCompose"
	"gopkg.in/yaml.v3"
)

func Fetch(urls ...string) ([]*dockerCompose.Config, error) {
	dockerComposeConfigs := make([]*dockerCompose.Config, 0)

	for _, rawURL := range urls {
		dockerComposeURL, err := url.Parse(rawURL)
		if err != nil {
			return nil, err
		}

		switch {
		case dockerComposeURL.Scheme == "http" || dockerComposeURL.Scheme == "https":
			dockerComposeConfig, err := getDockerComposeViaHTTP(dockerComposeURL.String())
			if err != nil {
				return nil, err
			}

			dockerComposeConfigs = append(dockerComposeConfigs, dockerComposeConfig)
		case dockerComposeURL.Scheme == "file":
			fallthrough
		default:
			dockerComposeConfig, err := readDockerComposeFromFile(dockerComposeURL.Path)
			if err != nil {
				return nil, err
			}

			dockerComposeConfigs = append(dockerComposeConfigs, dockerComposeConfig)
		}
	}

	return dockerComposeConfigs, nil
}

var ErrorPathIsDir error = errors.New("Path is a directory")

func getDockerComposeViaHTTP(url string) (*dockerCompose.Config, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received unexpected HTTP-Statuscode %v", resp.StatusCode)
	}

	dockerCompose := dockerCompose.NewConfig()

	yamlDecoder := yaml.NewDecoder(resp.Body)
	err = yamlDecoder.Decode(&dockerCompose)
	if err != nil {
		return nil, err
	}

	return dockerCompose, nil
}

func readDockerComposeFromFile(name string) (*dockerCompose.Config, error) {
	fileStat, err := os.Stat(name)
	switch {
	case err != nil:
		return nil, err
	case fileStat.IsDir():
		return nil, fmt.Errorf("%w: %s", ErrorPathIsDir, name)
	}

	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dockerCompose := dockerCompose.NewConfig()

	yamlDecoder := yaml.NewDecoder(file)
	err = yamlDecoder.Decode(&dockerCompose)
	if err != nil {
		return nil, err
	}

	return dockerCompose, nil
}
