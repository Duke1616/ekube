package v1

import (
	"ekube/conf"
	"encoding/base64"
	"fmt"
	"github.com/infraboard/mcube/crypto/cbc"
	"strings"
)

func (x *Cluster) EncryptKubeConf(key string) error {
	// 判断文本是否已经加密
	if strings.HasPrefix(x.Spec.KubeConfig, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text has ciphered")
	}

	cipherText, err := cbc.Encrypt([]byte(x.Spec.KubeConfig), []byte(key))
	if err != nil {
		return err
	}

	base64Str := base64.StdEncoding.EncodeToString(cipherText)
	x.Spec.KubeConfig = fmt.Sprintf("%s%s", conf.CIPHER_TEXT_PREFIX, base64Str)
	return nil
}

func (x *Cluster) DecryptKubeConf(key string) error {
	// 判断文本是否已经是明文
	if !strings.HasPrefix(x.Spec.KubeConfig, conf.CIPHER_TEXT_PREFIX) {
		return nil
	}

	base64CipherText := strings.TrimPrefix(x.Spec.KubeConfig, conf.CIPHER_TEXT_PREFIX)

	cipherText, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return err
	}

	planText, err := cbc.Decrypt([]byte(cipherText), []byte(key))
	if err != nil {
		return err
	}

	x.Spec.KubeConfig = string(planText)
	return nil
}
