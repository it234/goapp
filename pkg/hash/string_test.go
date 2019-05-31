package hash

import (
	"testing"
)

// 测试字符串hash

func Test_Md5String(t *testing.T) {
	val := Md5String("111111")
	if val != "96e79218965eb72c92a549dd5a330112" {
		t.Errorf("string md5值计算错误:%s", val)
	}
}

func Test_Sha1String(t *testing.T) {
	val := Sha1String("111111")
	if val != "3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d" {
		t.Errorf("string sha1值计算错误:%s", val)
	}
}

func Test_Sha256String(t *testing.T) {
	val := Sha256String("111111")
	if val != "bcb15f821479b4d5772bd0ca866c00ad5f926e3580720659cc80d39c9d09802a" {
		t.Errorf("string sha256值计算错误:%s", val)
	}
}

func Test_Sha512String(t *testing.T) {
	val := Sha512String("111111")
	if val != "b0412597dcea813655574dc54a5b74967cf85317f0332a2591be7953a016f8de56200eb37d5ba593b1e4aa27cea5ca27100f94dccd5b04bae5cadd4454dba67d" {
		t.Errorf("string sha512值计算错误:%s", val)
	}
}
