package worker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ntt360/pmon2/app"
	"github.com/ntt360/pmon2/app/executor"
	"github.com/ntt360/pmon2/app/model"
	"github.com/ntt360/pmon2/client/service"
)

func Start(processFile string, flags *model.ExecFlags) (string, error) {
	// prepare params
	file, err := os.Stat(processFile)
	if os.IsNotExist(err) || file.IsDir() {
		return "", errors.New(fmt.Sprintf("%s not exist", processFile))
	}

	// get run process user
	runUser, err := GetProcUser(flags)
	if err != nil {
		return "", nil
	}

	name := flags.Name
	// get process file name
	if len(name) <= 0 {
		name = filepath.Base(processFile)
	}

	// checkout process name whether exist
	if app.Db().First(&model.Process{}, "name = ?", name).Error == nil {
		return "", fmt.Errorf("process name: %s already exist, please set other name by --name", name)
	}

	// start process
	process, err := executor.Exec(processFile, flags.Log, name, flags.Args, flags.Dir, runUser, !flags.NoAutoRestart)
	if err != nil {
		return "", err
	}
	process.CreatedAt = time.Now()
	process.UpdatedAt = time.Now()

	// waiting process state
	var stat = service.NewProcStat(process).Wait()

	// return process data
	return service.AddData(stat)
}
