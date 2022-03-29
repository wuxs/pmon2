package logs

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ntt360/pmon2/app"
	"github.com/ntt360/pmon2/app/model"
	"github.com/nxadm/tail"
	"github.com/spf13/cobra"
)

var Follow bool

var Cmd = &cobra.Command{
	Use:     "logs",
	Short:   "output running process log",
	Example: "sudo pmon2 logs [id or name]",
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "--help" {

		}
		cmdRun(args)
	},
}

func init() {
	Cmd.Flags().BoolVarP(&Follow, "follow", "f", false, "follow")

}

func cmdRun(args []string) {
	val := args[0]

	var process model.Process
	err := app.Db().Find(&process, "name = ? or id = ?", val, val).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			app.Log.Fatal("pmon2 run err: %v", err)
		}

		// not found
		app.Log.Errorf("process %s not exist", val)
		return
	}

	t, err := tail.TailFile(process.Log, tail.Config{Follow: Follow})
	if err != nil {
		app.Log.Errorf("tail log %s error: %s", val, err.Error())
		return
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
