package encryption

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/CocaineCong/secret"
	"gorm.io/gorm/schema"
)

// MoneyEncryptSerializer 适合有固定key的字段
// eg:
// Money    string		`gorm:"column:money;serializer:pfm"`
type MoneyEncryptSerializer struct {
}

func init() {
	schema.RegisterSerializer("money", MoneyEncryptSerializer{})
}

func (MoneyEncryptSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)
	if dbValue != nil {
		var val string
		switch v := dbValue.(type) {
		case []byte:
			val = string(v)
		case string:
			val = v
		default:
			return fmt.Errorf("failed to decrypt value: %v", dbValue)
		}
		aesEncrypt, _ := secret.NewAesEncrypt("", "", "", secret.AesEncrypt128, secret.AesModeTypeCTR)
		var plain string
		if val != "" {
			plain = aesEncrypt.SecretDecrypt(plain)
		}
		err = json.Unmarshal([]byte(plain), fieldValue.Interface())
		if err != nil {
			err = json.Unmarshal([]byte(strconv.Quote(plain)), fieldValue.Interface())
		}
	}

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

func (MoneyEncryptSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (r interface{}, err error) {
	switch fieldValue.(type) {
	case string:
		value := fieldValue.(string)
		if value == "" {
			return value, nil
		}
		aesEncrypt, _ := secret.NewAesEncrypt("", "", "", secret.AesEncrypt128, secret.AesModeTypeCTR)
		return aesEncrypt.SecretEncrypt(value), nil

	default:
		return fieldValue, nil
	}
}
