package api

import (
	"fmt"
	"reflect"

	"github.com/ElrondNetwork/elrond-proxy-go/api/address"
	"github.com/ElrondNetwork/elrond-proxy-go/api/heartbeat"
	"github.com/ElrondNetwork/elrond-proxy-go/api/transaction"
	"github.com/ElrondNetwork/elrond-proxy-go/api/vmValues"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

type validatorInput struct {
	Name      string
	Validator validator.Func
}

// Start will boot up the api and appropriate routes, handlers and validators
func Start(elrondProxyFacade ElrondProxyHandler, port int) error {
	ws := gin.Default()
	ws.Use(cors.Default())

	err := registerValidators()
	if err != nil {
		return err
	}
	registerRoutes(ws, elrondProxyFacade)

	return ws.Run(fmt.Sprintf(":%d", port))
}

func registerRoutes(ws *gin.Engine, elrondProxyFacade ElrondProxyHandler) {
	addressRoutes := ws.Group("/address")
	addressRoutes.Use(WithElrondProxyFacade(elrondProxyFacade))
	address.Routes(addressRoutes)

	txRoutes := ws.Group("/transaction")
	txRoutes.Use(WithElrondProxyFacade(elrondProxyFacade))
	transaction.Routes(txRoutes)

	getValuesRoutes := ws.Group("/vm-values")
	getValuesRoutes.Use(WithElrondProxyFacade(elrondProxyFacade))
	vmValues.Routes(getValuesRoutes)

	heartbeatRoutes := ws.Group("/heartbeat")
	heartbeatRoutes.Use(WithElrondProxyFacade(elrondProxyFacade))
	heartbeat.Routes(heartbeatRoutes)
}

func registerValidators() error {
	validators := []validatorInput{
		{Name: "skValidator", Validator: skValidator},
	}
	for _, validatorFunc := range validators {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			err := v.RegisterValidation(validatorFunc.Name, validatorFunc.Validator)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// skValidator validates a secret key from user input for correctness
func skValidator(
	_ *validator.Validate,
	_ reflect.Value,
	_ reflect.Value,
	_ reflect.Value,
	_ reflect.Type,
	_ reflect.Kind,
	_ string,
) bool {
	return true
}

// WithElrondProxyFacade middleware will set up an ElrondFacade object in the gin context
func WithElrondProxyFacade(elrondProxyFacade ElrondProxyHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("elrondProxyFacade", elrondProxyFacade)
		c.Next()
	}
}
