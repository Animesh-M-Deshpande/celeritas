package main

import "fmt"

func doMigrate(arg2, arg3 string) error {

	dsn := getDSN()

	//fmt.Printf("calling up migration, dsn=%s", dsn)
	//fmt.Println("")

	switch arg2 {
	case "up":
		fmt.Println("calling up migration, dsn:" + dsn)
		err := cel.MigrateUp(dsn)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" {
			err := cel.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := cel.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}

	case "reset":

		err := cel.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = cel.MigrateUp(dsn)
		if err != nil {
			return err
		}

	default:
		showHelp()

	}

	return nil
}
