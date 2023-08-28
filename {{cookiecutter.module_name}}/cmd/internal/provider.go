package internal

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"
	"io/ioutil"
)

type EnvVariables struct {
	LogLevel               string `envconfig:"log_level" required:"true"`
	Application            string `envconfig:"application" required:"true"`
	Project                string `envconfig:"project" required:"true"`
	ProjectGroup           string `envconfig:"project_group" required:"true"`
	Version                string `envconfig:"version" required:"true"`
	CustomAttributesPrefix string `envconfig:"attributes_prefix" required:"true"`
	AWSCABundle            string `envconfig:"AWS_CA_BUNDLE"`
}

type Provider struct {
	awsConfig    aws.Config
	envVariables EnvVariables
}

func NewProvider() *Provider {
	var envVariables EnvVariables
	if err := envconfig.Process("", &envVariables); err != nil {
		panic(fmt.Sprintf("cannot load config: %v", err))
	}
	cfg := loadAwsConfig(envVariables.AWSCABundle)
	otelaws.AppendMiddlewares(&cfg.APIOptions)
	return &Provider{cfg, envVariables}
}

func (p *Provider) ProvideLoggerConfig() log.Configuration {
	return log.NewConfiguration(p.envVariables.LogLevel, p.envVariables.Application, p.envVariables.Project, p.envVariables.ProjectGroup, p.envVariables.Version, p.envVariables.CustomAttributesPrefix)
}

func loadAwsConfig(awsCABundlePath string, extraLoadOptions ...func(*config.LoadOptions) error) aws.Config {
	var loadOptions []func(*config.LoadOptions) error
	loadOptions = append(loadOptions, extraLoadOptions...)
	if len(awsCABundlePath) > 0 {
		loadOptions = append(loadOptions, func(options *config.LoadOptions) error {
			file, err := ioutil.ReadFile(awsCABundlePath)
			if err != nil {
				return err
			}
			options.CustomCABundle = bytes.NewReader(file)
			return nil
		})
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), loadOptions...)
	if err != nil {
		panic(errors.Wrap(err, "Failed to load default aws config"))
	}
	return cfg
}
