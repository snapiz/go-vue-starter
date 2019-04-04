package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/boilingcore"
	"github.com/volatiletech/sqlboiler/drivers"
	"github.com/volatiletech/sqlboiler/importers"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "db:models",
		Short: "Generate models from database schema",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := cgo.NewDB("", false)

			if err != nil {
				log.Fatal(err)
			}

			defer db.Close()

			driverName := "psql"
			driverPath := fmt.Sprintf("%s/bin/sqlboiler-%s", cgo.GetGOPath(), driverName)

			drivers.RegisterBinary(driverName, driverPath)

			cmdConfig := &boilingcore.Config{
				DriverName: driverName,
				OutFolder:  "src/models",
				PkgName:    "models",
				//Debug:      true,
				AddGlobal: true,
				//AddPanic:         viper.GetBool("add-panic-variants"),
				NoContext: true,
				NoTests:   true,
				//NoHooks:          viper.GetBool("no-hooks"),
				//NoRowsAffected:   viper.GetBool("no-rows-affected"),
				//NoAutoTimestamps: viper.GetBool("no-auto-timestamps"),
				//Wipe:             viper.GetBool("wipe"),
				//StructTagCasing:  strings.ToLower(viper.GetString("struct-tag-casing")), // camel | snake
				//TemplateDirs:     viper.GetStringSlice("templates"),
				//Tags:             viper.GetStringSlice("tag"),
				//Replacements:     viper.GetStringSlice("replace"),
				//Aliases:          boilingcore.ConvertAliases(viper.Get("aliases")),
				//TypeReplaces:     boilingcore.ConvertTypeReplace(viper.Get("types")),
			}

			if cmdConfig.Debug {
				fmt.Fprintln(os.Stderr, "using driver:", driverPath)
			}

			// Configure the driver
			cmdConfig.DriverConfig = map[string]interface{}{
				"host":      "localhost",
				"whitelist": strings.Split(os.Getenv("WHITELIST_TABLES"), ","),
				//"blacklist": viper.GetStringSlice(driverName + ".blacklist"),
			}

			sdbSource := strings.Split(cgo.GetDBSource(), " ")
			for _, value := range sdbSource {
				svalue := strings.Split(value, "=")
				if svalue[0] == "password" {
					svalue[0] = "pass"
				}
				cmdConfig.DriverConfig[svalue[0]] = svalue[1]
			}
			cmdConfig.Imports = importers.NewDefaultImports()
			cmdState, err := boilingcore.New(cmdConfig)
			if err != nil {
				log.Fatal(err)
			}
			err = cmdState.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	})
}
