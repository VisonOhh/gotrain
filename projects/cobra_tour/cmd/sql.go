package cmd

import (
	"log"

	"github.com/gotrain/projects/cobra_tour/internal/sql2struct"
	"github.com/spf13/cobra"
)

func init() {
	sqlCmd.AddCommand(sqlCmdStruct)

	sqlCmdStruct.Flags().StringVarP(&username, "username", "", "", "pls input db username")
	sqlCmdStruct.Flags().StringVarP(&password, "password", "", "", "pls input db password")
	sqlCmdStruct.Flags().StringVarP(&host, "host", "", "", "pls input db host")
	sqlCmdStruct.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "pls input db charset")
	sqlCmdStruct.Flags().StringVarP(&dbType, "type", "", "", "pls input db type")
	sqlCmdStruct.Flags().StringVarP(&dbName, "db", "", "", "pls input db name")
	sqlCmdStruct.Flags().StringVarP(&tableName, "table", "", "", "pls input db table name")
}

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "convert db structure to go structure",
	Long:  "convert db structure to go structure",
}

var sqlCmdStruct = &cobra.Command{
	Use:   "struct",
	Short: "mysql conversion",
	Long:  "mysql conversion",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}

		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}
