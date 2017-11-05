package imagedock

import (
	minio "github.com/minio/minio-go"
)

func Connect() {
	//log.Debug("hey")
	mc, err := minio.New("localhost",
		"IK8XNAH6XL1CZK6TBP93",
		"Y+2i4Pqn8eu2kmliJNiGlvqXnti3BWyeR2Ra2Ugu",
		false)
	mc.BucketExists("blah")
	if err != nil {
	}
}
