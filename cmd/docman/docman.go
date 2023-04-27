package docman

import (
	"docman/cfg"
	Init "docman/init"
	"docman/pkg/log"
	"docman/pkg/server"
	"os"

	"github.com/spf13/cobra"
)

var (
	serverConf = &cfg.Config.Server
	rootCmd    = &cobra.Command{
		Use:     "doc",
		Version: serverConf.Version,
		Short:   "A brief description of your application",
		Long:    `A longer description that spans multiple lines and likely contains examples and usage of using your application. For example:`,
		Run: func(cmd *cobra.Command, args []string) {
			dev := cmd.Flags().Changed("dev")
			prod := cmd.Flags().Changed("prod")
			env := cmd.Flags().Changed("Env")
			if dev {
				serverConf.Env = "dev"
			} else if prod {
				serverConf.Env = "prod"
			} else if env && serverConf.Env != "dev" && serverConf.Env != "prod" {
				log.Error("Env value must be dev or prod")
				os.Exit(1)
			} else {
				// default dev
			}

		},
	}
)

func init() {
	// 初始化相关配置或服务
	rootCmd.Flags().Bool("dev", false, "run in dev mode")
	rootCmd.Flags().Bool("prod", false, "run in prod mode")
	rootCmd.Flags().StringVarP(&serverConf.Env, "Env", "e", "dev", "set Env value")
	rootCmd.Flags().UintVarP(&serverConf.Port, "port", "p", 0, "set port value")
	// if user input -h or --help, print help info and exit
	rootCmd.Flags().BoolP("help", "h", false, "help info")

}
func init() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

// Execute docman程序入口
func Execute() {
	if err := Init.Init(); err != nil {
		log.Error("init failed", err.Error())
		os.Exit(1)
	}
	if err := server.Run(); err != nil {
		log.Error("server run failed", err.Error())
		os.Exit(1)
	}
}
