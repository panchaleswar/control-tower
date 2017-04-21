package config

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/engineerbetter/concourse-up/aws"
)

const terraformStateFileName = "terraform.tfstate"
const configBucketS3Region = "eu-west-1"
const configFilePath = "config.json"

// IClient is an interface for the config file client
type IClient interface {
	Load() (*Config, error)
	LoadOrCreate() (*Config, bool, error)
	StoreAsset(filename string, contents []byte) error
	HasAsset(filename string) (bool, error)
	LoadAsset(filename string) ([]byte, error)
}

// Client is a client for loading the config file  from S3
type Client struct {
	Project string
}

// StoreAsset stores an associated configuration file
func (client *Client) StoreAsset(filename string, contents []byte) error {
	return aws.WriteFile(client.configBucket(),
		filename,
		configBucketS3Region,
		contents,
	)
}

// LoadAsset loads an associated configuration file
func (client *Client) LoadAsset(filename string) ([]byte, error) {
	return aws.LoadFile(
		client.configBucket(),
		filename,
		configBucketS3Region,
	)
}

// HasAsset returns true if an associated configuration file exists
func (client *Client) HasAsset(filename string) (bool, error) {
	return aws.HasFile(
		client.configBucket(),
		filename,
		configBucketS3Region,
	)
}

// Load loads an existing config file from S3
func (client *Client) Load() (*Config, error) {
	configBytes, err := aws.LoadFile(
		client.configBucket(),
		configFilePath,
		configBucketS3Region,
	)
	if err != nil {
		return nil, err
	}

	conf := Config{}
	if err := json.Unmarshal(configBytes, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

// LoadOrCreate loads an existing config file from S3, or creates a default if one doesn't already exist
func (client *Client) LoadOrCreate() (*Config, bool, error) {
	defaultConfigBytes, err := generateDefaultConfig(
		client.Project,
		client.deployment(),
		configBucketS3Region,
	)
	if err != nil {
		return nil, false, err
	}

	if err = aws.EnsureBucketExists(client.configBucket(), configBucketS3Region); err != nil {
		return nil, false, err
	}

	configBytes, createdNewFile, err := aws.EnsureFileExists(
		client.configBucket(),
		configFilePath,
		configBucketS3Region,
		defaultConfigBytes,
	)
	if err != nil {
		return nil, false, err
	}

	conf := Config{}
	if err := json.Unmarshal(configBytes, &conf); err != nil {
		return nil, false, err
	}

	return &conf, createdNewFile, nil
}

func (client *Client) deployment() string {
	return fmt.Sprintf("concourse-up-%s", client.Project)
}

func (client *Client) configBucket() string {
	return fmt.Sprintf("%s-config", client.deployment())
}