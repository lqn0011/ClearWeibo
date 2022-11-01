package conf

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/joho/godotenv"
)

//注意已经存在的环境变量优先于.env内的变量
func LoadConfig(file string, v interface{}) (err error) {
	dir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		return
	}

	dotenv := filepath.Join(dir, ".env")
	if _, err := os.Stat(dotenv); err == nil {
		err = godotenv.Load(dotenv) //如果要.env优先于环境变量可改为Overload
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
	} else {
		// file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	content = []byte(os.ExpandEnv(string(content)))
	err = yaml.Unmarshal(content, v)
	if err != nil {
		return
	}

	return
}

func MustLoadConfig(file string, v interface{}) {
	err := LoadConfig(file, v)
	if err != nil {
		panic(err)
	}
}
