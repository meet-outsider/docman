package docman

import (
	"docman/cfg"
	Init "docman/init"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	server  = &cfg.Config.Server
	rootCmd = &cobra.Command{
		Use:     "doc",
		Version: server.Version,
		Short:   "A brief description of your application",
		Long:    `A longer description that spans multiple lines and likely contains examples and usage of using your application. For example:`,
		Run: func(cmd *cobra.Command, args []string) {
			dev := cmd.Flags().Changed("dev")
			prod := cmd.Flags().Changed("prod")
			env := cmd.Flags().Changed("Env")
			if dev {
				server.Env = "dev"
			} else if prod {
				server.Env = "prod"
			} else if env && server.Env != "dev" && server.Env != "prod" {
				fmt.Println("Env args error")
				os.Exit(1)
			} else {
				//todo##
			}

		},
	}
)

func init() {
	// 初始化相关配置或服务
	rootCmd.Flags().Bool("dev", false, "run in dev mode")
	rootCmd.Flags().Bool("prod", false, "run in prod mode")
	rootCmd.Flags().StringVarP(&server.Env, "Env", "e", "dev", "set Env value")
	rootCmd.Flags().UintVarP(&server.Port, "port", "p", 0, "set port value")

}
func init() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Execute docman程序入口
func Execute() {
	if err := Init.Init(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("docman Env:%s Listening port: %d\n", server.Env, server.Port)
}
