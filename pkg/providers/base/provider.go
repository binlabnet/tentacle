package base

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
	"github.com/analogj/tentacle/pkg/utils"
	"github.com/analogj/tentacle/pkg/errors"
	"github.com/analogj/tentacle/pkg/credentials"
	"encoding/json"
)

type Provider struct {
	Alias          string
	ProviderConfig map[string]interface{}

	//Global Options
	DebugMode	bool
	OutputMode	string
}

func (p *Provider) Authenticate() error {
	return errors.NotImplementedError("Authenticate function not implemented")
}

func (p *Provider) Get(queryData map[string]string) error {
	return errors.NotImplementedError("Get function not implemented")
}

func (p *Provider) List(queryData map[string]string)  ([]credentials.SummaryInterface, error) {
	return nil, errors.NotImplementedError("List function not implemented")
}


//utility/helper functions

var reservedFlags = []string{ "output", "debug" }

func  (p *Provider) CommandProcessFlagsToQueryData(c *cli.Context) map[string]string {


	queryData := map[string]string{}
	for _, flagName := range c.FlagNames() {
		//skip over reserved flags.
		if utils.SliceIncludes(reservedFlags, flagName){
			continue
		}

		queryData[flagName] = c.String(flagName)
	}

	if p.DebugMode{
		fmt.Printf("DEBUG: Query data: %#v\n", queryData)
	}

	return queryData
}

func (p *Provider) CommandProcessGlobalFlags(c *cli.Context) error {


	p.DebugMode = c.Bool("debug")
	p.OutputMode = c.String("output")

	if p.DebugMode {
		fmt.Println("DEBUG: enabled debug mode")
	}

	return nil
}

func (p *Provider) CommandPrintCredentials(c *cli.Context, commandType string, credentialData interface{}, credentialError error) error {

	if credentialError != nil {
		return credentialError
	}
	switch commandType {
	case "get":
		return PrintCredential(p.OutputMode, credentialData.(credentials.GenericInterface))
	case "list":
		return PrintCredentials(p.OutputMode, credentialData.([]credentials.SummaryInterface))
	default:
		return errors.InvalidArgumentsError(fmt.Sprintf("command type is invalid (%v)", commandType))
	}

	return nil
}

func PrintCredential(outputMode string, credential credentials.GenericInterface) error {

	switch outputMode {
	case "json":
		secret, err := credential.ToJsonString()
		if err != nil {
			return err
		}
		fmt.Println(secret)
	case "yaml":
		secret, err := credential.ToYamlString()
		if err != nil {
			return err
		}
		fmt.Println(secret)
	case "raw":
		secret, err := credential.ToRawString()
		if err != nil {
			return err
		}
		fmt.Println(secret)
	case "table":
		secret, err := credential.ToTableString()
		if err != nil {
			return err
		}
		fmt.Println(secret)

	default:
		return errors.InvalidArgumentsError(fmt.Sprintf("output mode is invalid (%v)", outputMode))
	}
	return nil
}

func PrintCredentials(outputMode string, credentials []credentials.SummaryInterface) error {

	switch outputMode {
	case "json":
		jsonBytes, err := json.MarshalIndent(credentials,"", "    ")
		if err != nil {
			return err
		}
		fmt.Printf(string(jsonBytes))
	case "raw":
		return errors.InvalidArgumentsError("output mode raw is unsupported for list")
	case "table":
		//TODO: print a table.
		//fmt.Println(credential.ToTableString())
	default:
		return errors.InvalidArgumentsError(fmt.Sprintf("output mode is invalid (%v)", outputMode))
	}
	return nil
}