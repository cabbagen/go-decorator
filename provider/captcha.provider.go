package provider

import (
	"time"
	"go-decorator/cache"
	"github.com/mojocn/base64Captcha"
)

type RedisCaptchaStore struct {
	Expiration   time.Duration
}

func (rsk RedisCaptchaStore) Set(id string, value string) {
	cache.GetRedisCacheInstance().Set(id, value, rsk.Expiration)
}

func (rsk RedisCaptchaStore) Get(id string, clear bool) string {
	value, error := cache.GetRedisCacheInstance().Get(id)

	if error != nil {
		return value
	}
	if clear {
		cache.GetRedisCacheInstance().Del(id)
	}
	return value
}

func (rsk RedisCaptchaStore) Verify(id, answer string, clear bool) bool {
	value := rsk.Get(id, clear)
	return value == answer
}

var defaultCaptchaInstance *base64Captcha.Captcha = base64Captcha.NewCaptcha(
	base64Captcha.DefaultDriverDigit,
	RedisCaptchaStore { time.Minute * 10 },
)

func GenerateCaptcha() (map[string]string, error) {
	var captchaInfo map[string]string = make(map[string]string)

	id, b64s, error := defaultCaptchaInstance.Generate()

	if error != nil {
		return nil, error
	}

	captchaInfo["b64s"] = b64s
	captchaInfo["captchaId"] = id

	return captchaInfo, nil
}

func ValidateCaptcha(id, answer string) bool {
	return defaultCaptchaInstance.Verify(id, answer, true)
}
