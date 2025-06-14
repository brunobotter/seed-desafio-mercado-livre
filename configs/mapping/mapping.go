package mapping

type Config struct {
	DB DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	Host                 string            `mapstructcture:"host"`
	Name                 string            `mapstructcture:"name"`
	User                 string            `mapstructcture:"user"`
	Pass                 string            `mapstructcture:"pass"`
	Port                 int               `mapstructcture:"port"`
	Params               map[string]string `mapstructcture:"params"`
	MaxLifeTimeInMinutes int               `mapstructcture:"life_time"`
	MaxIdleConns         int               `mapstructcture:"idle_conn"`
	MaxOpenConns         int               `mapstructcture:"max_open_conn"`
}
