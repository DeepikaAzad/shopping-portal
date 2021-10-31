package commands

import (
	"fmt"
	"log"

	"github.com/DeepikaAzad/go-to-do-app/go-server/database"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, sqlDB := database.Connection()
			defer sqlDB.Close()
			tx := db.Begin()
			for _, migrate := range database.AutoMigrate(tx) {
				if err := migrate.Run(tx); err != nil {
					tx.Rollback()
					log.Panic(errors.Wrap(err, "[Migrate] Running migration `"+migrate.TableName+"` failed with error: "))
				}
			}
			tx.Commit()
			fmt.Printf("Migration completed!\n")
			return nil
		},
	}
}
