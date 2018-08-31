package main
import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const cmdRoot = "core"

func main() {
	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath("./viper")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
		os.Exit(1)
	}

	//test json 文件
	fmt.Println(viper.Get("id"))
	fmt.Println(viper.GetBool("isKeyAttribute"))

	// test yaml 文件
	//environment := viper.GetBool("security.enabled")
	//fmt.Println("security.enabled:", environment)
	//
	//fullstate := viper.GetString("statetransfer.timeout.fullstate")
	//fmt.Println("statetransfer.timeout.fullstate:", fullstate)
	//
	//abcdValue := viper.GetString("peer.abcd")
	//fmt.Println("peer.abcd:", abcdValue)
}




