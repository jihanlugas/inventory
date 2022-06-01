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
	drop: delete database 
	`,
}

var dbCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database table",
	Long:  `Create database table`,
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

var dbDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop database table",
	Long:  `Drop database table`,
	Run: func(cmd *cobra.Command, args []string) {
		drop()
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbCreateCmd)
	dbCmd.AddCommand(dbDropCmd)
}

func create() {
	var err error

	dbpool := db.Initialize()
	defer dbpool.Close()

	sqlContent, err := ioutil.ReadFile("./docs/sql/create.sql")
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

	_, err = tx.Exec(ctx, sql)
	if err != nil {
		fmt.Println("err ", err)
		panic(err)
	}

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		panic(err)
	}

	fmt.Println("success create")
}

func drop() {
	var err error

	dbpool := db.Initialize()
	defer dbpool.Close()

	sqlContent, err := ioutil.ReadFile("./docs/sql/drop.sql")
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

	_, err = tx.Exec(ctx, sql)
	if err != nil {
		fmt.Println("err ", err)
		panic(err)
	}

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		panic(err)
	}

	fmt.Println("success drop")
}
