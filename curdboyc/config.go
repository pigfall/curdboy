package curdboyc

type Config struct{
	entSchemaDirPath string
	targetDirPath string
	entTargetDirPath string
}

func NewConfig(entSchemaDirPath string,targetDirPath string,entTargetDirPath string) *Config{
	return &Config{
		entSchemaDirPath: entSchemaDirPath,
		targetDirPath: targetDirPath,
		entTargetDirPath: entTargetDirPath,
	}
}
