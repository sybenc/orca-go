package conf

import (
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper 实例冲突
	"os"
	"time"
)

// viper 库实例
var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 1. 初始化 Viper 库
	viper = viperlib.New()

	// 2. 配置类型，设置为 yaml
	viper.SetConfigType("yaml")

	// 3. 添加配置文件路径
	viper.AddConfigPath(".") // 可以根据项目结构调整路径

	// 4. 自动加载环境变量，支持环境变量覆盖
	viper.AutomaticEnv()

	// 5. 初始化配置函数 map
	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 用于加载指定环境的配置文件
func InitConfig(env string) {
	// 加载 YAML 配置文件
	loadYamlConfig(env)
	// 注册配置信息
	loadConfig()
}

// loadConfig 将动态配置加载到 Viper 中
func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

// loadYamlConfig 加载 YAML 文件配置
func loadYamlConfig(envSuffix string) {
	configFileName := "./config.yaml" // 默认的配置文件

	// 支持不同环境的配置文件，如 config.testing.yaml
	if len(envSuffix) > 0 {
		filepath := "config." + envSuffix + ".yaml"
		if _, err := os.Stat(filepath); err == nil {
			configFileName = filepath
		}
	}

	// 设置配置文件名（包含路径）
	viper.SetConfigFile(configFileName)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err) // 配置文件读取失败时，抛出错误
	}

	// 监控配置文件变化并自动重新加载
	viper.WatchConfig()
}

// 通用的内部获取配置值函数
func internalGet(path string, defaultValue ...interface{}) interface{} {
	// 如果配置不存在，则返回默认值
	if !viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

func GetUint64(path string, defaultValue ...interface{}) uint64 {
	return cast.ToUint64(internalGet(path, defaultValue...))
}

func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

func GetTime(path string, defaultValue ...interface{}) time.Time {
	return cast.ToTime(internalGet(path, defaultValue...))
}
func GetDuration(path string, defaultValue ...interface{}) time.Duration {
	return cast.ToDuration(internalGet(path, defaultValue...))
}
func GetStringSlice(path string, defaultValue ...interface{}) []string {
	return cast.ToStringSlice(internalGet(path, defaultValue...))
}
func GetStringMap(path string, defaultValue ...interface{}) map[string]interface{} {
	return cast.ToStringMap(internalGet(path, defaultValue...))
}
func GetStringMapString(path string, defaultValue ...interface{}) map[string]string {
	return cast.ToStringMapString(internalGet(path, defaultValue...))
}
