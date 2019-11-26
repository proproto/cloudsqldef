package main

import (
	"fmt"
	"log"
	"os"

	"github.com/howeyc/gopass"
	"github.com/jessevdk/go-flags"
	"github.com/proproto/cloudsqldef"
	"github.com/proproto/cloudsqldef/adapter"
	"github.com/proproto/cloudsqldef/adapter/mysql"
	"github.com/proproto/cloudsqldef/schema"
)

// Return parsed options and schema filename
// TODO: Support `sqldef schema.sql -opt val...`
func parseOptions(args []string) (adapter.Config, *sqldef.Options) {
	var opts struct {
		User     string `short:"u" long:"user" description:"MySQL user name" value-name:"user_name" default:"root"`
		Password string `short:"p" long:"password" description:"MySQL user password, overridden by $MYSQL_PWD" value-name:"password"`
		Host     string `short:"h" long:"host" description:"Host to connect to the MySQL server" value-name:"host_name" default:"127.0.0.1"`
		Port     uint   `short:"P" long:"port" description:"Port used for the connection" value-name:"port_num" default:"3306"`
		Socket   string `short:"S" long:"socket" description:"The socket file to use for connection" value-name:"socket"`
		Prompt   bool   `long:"password-prompt" description:"Force MySQL user password prompt"`
		File     string `long:"file" description:"Read schema SQL from the file, rather than stdin" value-name:"sql_file" default:"-"`
		DryRun   bool   `long:"dry-run" description:"Don't run DDLs but just show them"`
		Export   bool   `long:"export" description:"Just dump the current schema to stdout"`
		SkipDrop bool   `long:"skip-drop" description:"Skip destructive changes such as DROP"`
		Help     bool   `long:"help" description:"Show this help"`
	}

	parser := flags.NewParser(&opts, flags.None)
	parser.Usage = "[options] db_name"
	args, err := parser.ParseArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Print("No database is specified!\n\n")
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("Multiple databases are given: %v\n\n", args)
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	database := args[0]

	options := sqldef.Options{
		SqlFile:  opts.File,
		DryRun:   opts.DryRun,
		Export:   opts.Export,
		SkipDrop: opts.SkipDrop,
	}

	password, ok := os.LookupEnv("MYSQL_PWD")
	if !ok {
		password = opts.Password
	}

	if opts.Prompt {
		fmt.Printf("Enter Password: ")
		pass, err := gopass.GetPasswd()
		if err != nil {
			log.Fatal(err)
		}
		password = string(pass)
	}

	config := adapter.Config{
		DbName:   database,
		User:     opts.User,
		Password: password,
		Host:     opts.Host,
		Port:     int(opts.Port),
		Socket:   opts.Socket,
	}
	return config, &options
}

func main() {
	config, options := parseOptions(os.Args[1:])

	database, err := mysql.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	sqldef.Run(schema.GeneratorModeMysql, database, options)
}
