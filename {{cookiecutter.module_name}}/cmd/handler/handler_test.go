package handler

import (
	"context"
	"fmt"
	"github.com/Ryanair/gofrlib-test/frContainers"
	"github.com/Ryanair/gofrlib-test/testDefaults"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

type HandlerIntegrationSuite struct {
	frContainers.BaseLocalstackIntegrationSuite
}

func TestHandlerIntegrationSuite(t *testing.T) {
	suite.Run(t, &HandlerIntegrationSuite{
		BaseLocalstackIntegrationSuite: frContainers.NewBaseLocalstackIntegrationSuite(func() (testcontainers.Container, string) {
			return frContainers.RunLocalstackServices(context.Background(), "...")
		}),
	})
}

func (suite *HandlerIntegrationSuite) SetupTest() {
	fmt.Println("Before test")
}

func (suite *HandlerIntegrationSuite) TearDownTest() {
	fmt.Println("After test")
}

func (suite *HandlerIntegrationSuite) Test_RenameMe() {
	// WHEN
	err := New(testDefaults.LoggerConfig).Handle(context.Background())

	// THEN
	suite.NoError(err)
}
