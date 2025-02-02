package monitor

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

const (
	UA              = "neighborhood/9.50.2 (iPhone; iOS 15.4.1; Scale/3.00)"
	CITY            = "0101"
	API_VERSION     = "9.50.2"
	BUILD_VERSION   = "1232"
	TIME_OUT        = 10 * time.Second
	BOOKABLE        = "可预约"
	NOTICE_BOOKABLE = "可以预约啦"
	NOTICE_TITLE    = "运力监控"
)

var Conf = new(config)

type config struct {
	Mode      int    `mapstructure:"mode"`
	Rate      uint   `mapstructure:"rate"`
	StationId string `mapstructure:"station_id"`
	Longitude string `mapstructure:"longitude"`
	Latitude  string `mapstructure:"latitude"`
	Bark      struct {
		Id    []string `mapstructure:"id"`
		Sound string   `mapstructure:"sound"`
	}
}

func InnitConfig(path string) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("fatal error reading configuration file: %s \n", err)
		panic(fmt.Errorf("fatal error reading configuration file: %s \n", err))
	}
	// Unmarshal configuration
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("unmarshal configuration failed, err: %s \n", err)
		panic(fmt.Errorf("unmarshal configuration failed, err: %s \n", err))
	}

	if Conf.StationId == "" {
		fmt.Println("rquire station_id")
		if Conf.Latitude == "" || Conf.Longitude == "" {
			panic(fmt.Errorf("请至少填写station_id和经纬度的其中一项"))
		}
		Conf.StationId = GetStationId()
		if Conf.StationId == "" {
			panic(fmt.Errorf("获取站点信息失败"))
		}

	}
	if len(Conf.Bark.Id) == 0 {
		fmt.Println("require bark_id")
		panic(fmt.Errorf("require bark_id"))
	}
	if Conf.Mode == 0 {
		fmt.Println("当前为本机模式运行")
	} else if Conf.Mode == 1 {
		fmt.Println("当前为GitHub Action模式运行")
	}
	if Conf.Rate == 0 {
		Conf.Rate = 3600
	}
}
