package tools

import "github.com/speps/go-hashids/v2"

// HashID encodes values into hashid, ref: https://hashids.org/
func HashID(values ...int64) (string, error) {
	hd := hashids.NewData()
	hd.Salt = SALT

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	id, err := h.EncodeInt64(values)
	if err != nil {
		return "", err
	}

	return id, nil
}
