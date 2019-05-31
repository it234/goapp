package hash

import (
	"testing"
)

// 测试字节数组hash

func Test_Md5Byte(t *testing.T) {
	val := Md5Byte([]byte("111111"))
	if val != "96e79218965eb72c92a549dd5a330112" {
		t.Errorf("byte md5值计算错误:%s", val)
	}
}

func Test_Sha1Byte(t *testing.T) {
	val := Sha1Byte([]byte("111111"))
	if val != "3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d" {
		t.Errorf("byte sha1值计算错误:%s", val)
	}
}

func Test_Sha256Byte(t *testing.T) {
	val := Sha256Byte([]byte("111111"))
	if val != "bcb15f821479b4d5772bd0ca866c00ad5f926e3580720659cc80d39c9d09802a" {
		t.Errorf("byte sha256值计算错误:%s", val)
	}
}

func Test_Sha512Byte(t *testing.T) {
	val := Sha512Byte([]byte("111111"))
	if val != "b0412597dcea813655574dc54a5b74967cf85317f0332a2591be7953a016f8de56200eb37d5ba593b1e4aa27cea5ca27100f94dccd5b04bae5cadd4454dba67d" {
		t.Errorf("byte sha512值计算错误:%s", val)
	}
}
