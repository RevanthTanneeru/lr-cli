package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MakeNowJust/heredoc"
	"github.com/loginradius/lr-cli/api"
	"github.com/loginradius/lr-cli/request"

	"github.com/loginradius/lr-cli/cmdutil"
	"github.com/loginradius/lr-cli/config"

	"github.com/spf13/cobra"
)

var field int

<<<<<<< HEAD
=======
/*lr add schema --feild 1
Enter the Display Name (About): About You
Is Required (Y/n): Y
Do you want to set Advance Configuiration for this feild(Y/n): Yes
Select Field Type*/
type Schema struct {
	Display          string `json:"Display"`
	Enabled          bool   `json:"Enabled"`
	IsMandatory      bool   `json:"IsMandatory"`
	Parent           string `json:"Parent"`
	ParentDataSource string `json:"ParentDataSource"`
	Permission       string `json:"Permission"`
	Name             string `json:"name"`
	Rules            string `json:"rules"`
	Status           string `json:"status"`
	Type             string `json:"type"`
}

type Result struct {
	Data []Schema `json:"Data"`
}

>>>>>>> Implement registration schema commands in LoginRadius CLI
func NewschemaCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "schema",
		Short:   "delete schema",
		Long:    `This commmand deletes schema fields`,
		Example: heredoc.Doc(`$ lr delete schema --fieldname <fieldname>`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if field == 0 {
				return &cmdutil.FlagError{Err: errors.New("`fieldname` is required argument")}
			}
			// if fieldName == "emailid" || fieldName == "password" {
			// 	return &cmdutil.FlagError{Err: errors.New("EmailId and Password fields cannot be deleted")}
			// }
			return delete(field)

		},
	}

	fl := cmd.Flags()
	fl.IntVarP(&field, "fieldname", "f", 0, "field name")

	return cmd
}

func delete(Field int) error {
	res, err := api.GetSites()
	if err != nil {
		return err
	}
	if res.Productplan.Name == "free" {
		fmt.Println("Kindly Upgrade the plan to enable this command for your app")
		return nil
	}

	var url string
	conf := config.GetInstance()

	resultResp1, err := api.GetFields("active")

	if Field > len(resultResp1.Data) {
		fmt.Println("please run 'lr get schema -all' first. Please enter the field number accordingly")
		return nil
	}

	if resultResp1.Data[Field-1].Name == "emailid" || resultResp1.Data[Field-1].Name == "password" {
		fmt.Println("EmailId and Password fields cannot be deleted")
		return nil
	}

	resultResp1.Data[Field-1].Enabled = false

	body, _ := json.Marshal(resultResp1)
	url = conf.AdminConsoleAPIDomain + "/platform-configuration/default-fields?"

	var resultResp api.ResultResp
	resp, err := request.Rest(http.MethodPost, url, nil, string(body))
	err = json.Unmarshal(resp, &resultResp)
	if err != nil {
		return err
	}
	fmt.Println("The field has been sucessfully deleted")
	return nil
}
