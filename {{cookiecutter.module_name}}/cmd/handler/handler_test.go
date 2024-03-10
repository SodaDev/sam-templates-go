package handler

import (
	"context"
	"fmt"
	"github.com/Ryanair/gofrlib-test/frContainers"
	"github.com/Ryanair/gofrlib-test/testDefaults"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HandlerIntegrationSuite struct {
	frContainers.BaseLocalstackIntegrationSuite
	handler *LambdaHandler
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
	suite.handler = New(testDefaults.LoggerConfig)
}

func (suite *HandlerIntegrationSuite) TearDownTest() {
	fmt.Println("After test")
}

func (suite *HandlerIntegrationSuite) Test_RenameMe() {
	// WHEN
	err := suite.handler.Handle(context.Background())

	// THEN
	suite.NoError(err)
}
