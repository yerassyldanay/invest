package utils

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

func Convert_string_to_hash(str string) (string, error) {
	bt, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err == nil {
		str = string(bt)
	}
	return str, err
}

func generateResultId() (string, error) {
	const alphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	k := make([]byte, 7)
	for i := range k {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphaNum))))
		if err != nil {
			return "", err
		}
		k[i] = alphaNum[idx.Int64()]
	}
	return string(k), nil
}

// GenerateId generates a unique key to represent the result
// in the database
//func (r *Result) GenerateId(tx *gorm.DB) error {
//	// Keep trying until we generate a unique key (shouldn't take more than one or two iterations)
//	for {
//		rid, err := generateResultId()
//		if err != nil {
//			return err
//		}
//		r.RId = rid
//		err = tx.Table("results").Where("r_id=?", r.RId).First(&Result{}).Error
//		if err == gorm.ErrRecordNotFound {
//			break
//		}
//	}
//	return nil
//}
