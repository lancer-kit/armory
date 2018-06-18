package uis

import (
	"database/sql"
	"errors"
	"time"

	"gitlab.inn4science.com/vcg/go-common/ith/ams"
)

type (
	RequestState struct {
		Id          int64             `json:"id" db:"id"`
		UserId      int64             `json:"-" db:"user_id"`
		State       ams.RequestStatus `json:"state" db:"state"`
		Data        string            `json:"data,omitempty" db:"data"`
		CreatedAt   *SqlTime          `json:"createdAt,omitempty" db:"created_at"`
		ProcessedAt *SqlTime          `json:"processedAt,omitempty" db:"processed_at"`
	}

	RegistrationStatus struct {
		UserId       int64             `json:"userId"`
		IsRegistered bool              `json:"isRegistered"`
		Status       ams.AccountStatus `json:"status"`
	}

	SqlTime int64
)

const SqlTimeLayout = "2006-01-02 15:04:05"

func (t *SqlTime) ToTime() time.Time {
	if t == nil {
		return time.Time{}
	}
	return time.Unix(int64(*t), 0)
}

func (t *SqlTime) String() string {
	return t.ToTime().Format(SqlTimeLayout)
}

func (t *SqlTime) FromString(s string) error {
	tmp, err := time.Parse(SqlTimeLayout, s)
	if err != nil {
		return err
	}
	*t = SqlTime(tmp.UTC().Unix())
	return nil
}
func (t *SqlTime) Scan(src interface{}) error {
	switch src.(type) {
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("AccountStatus: can't scan column data into int64")
		}

		*t = SqlTime(ni.Int64)
		return nil
	}
	return errors.New("SqlTime: invalid type")
}
