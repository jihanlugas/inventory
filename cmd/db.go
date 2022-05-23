package cmd

import (
	"fmt"
	"github.com/jihanlugas/inventory/db"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Command related with database",
	Long: `With this command you can
	create : create new database if not exists
	destroy: delete database 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify specific database command")
	},
}

var dbCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database table",
	Long:  `Create database table`,
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbCreateCmd)
}

func create() {
	var err error

	dbpool := db.Initialize()
	defer dbpool.Close()

	sqlContent, err := ioutil.ReadFile("./docs/sql/db.sql")
	if err != nil {
		panic(err)
	}
	sql := string(sqlContent)

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	tx, err := conn.Begin(ctx)
	if err != nil {
		panic(err)
	}
	defer db.DeferHandleTransaction(ctx, tx)

	exec, err := tx.Exec(ctx, sql)
	if err != nil {
		fmt.Println("err ", err)
		panic(err)
	}

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		panic(err)
	}

	fmt.Println("exec ", exec)
	fmt.Println("success")
}
