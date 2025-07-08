package dao

import (
	"context"
)

const (
	DURATION     = 300 // 5分钟
	SENDDuration = 60
)

func (d *Dao) StoreVerificationCode(ctx context.Context, key string, value string) error {
	err := d.rc.SetexCtx(ctx, "user:"+key+";EmailCode", value, DURATION)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) StoreSendCodeTime(ctx context.Context, key string) error {
	err := d.rc.SetexCtx(ctx, "user:"+key+";EmailTime", "ok", SENDDuration)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetSendCodeTime(ctx context.Context, key string) (string, error) {
	result, err := d.rc.GetCtx(ctx, "user:"+key+";EmailTime")
	if err != nil {
		return "", err
	}
	return result, nil
}

func (d *Dao) GetVerificationCode(ctx context.Context, key string) (string, error) {
	result, err := d.rc.GetCtx(ctx, "user:"+key+";EmailCode")
	if err != nil {
		return "", err
	}
	return result, nil
}
