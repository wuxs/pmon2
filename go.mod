module github.com/ntt360/pmon2

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/goinbox/gomisc v1.2.0 // indirect
	github.com/goinbox/shell v1.0.1
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/ntt360/gracehttp v1.4.6
	github.com/olekukonko/tablewriter v0.0.4
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/struCoder/pidusage v0.1.3
	gopkg.in/fsnotify.v1 v1.4.9 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace gopkg.in/fsnotify.v1 v1.4.9 => github.com/fsnotify/fsnotify v1.4.9
