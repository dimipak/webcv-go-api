package config

type Storage struct {
	Filesystem   string `env:"FILESYSTEM"`
	S3BucketName string `env:"S3_BUCKET_NAME"`
	S3Region     string `env:"S3_REGION"`
	S3KeyId      string `env:"S3_KEY_ID"`
	S3SecretKey  string `env:"S3_SECRET_KEY"`
}

func (s *Storage) setValues() {
	envEncode(s)
}
