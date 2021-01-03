package mysql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/icza/gox/gox"
	"gorm.io/gorm"

	"github.com/ktakenaka/go-random/backend/app/domain/entity"
	db "github.com/ktakenaka/go-random/backend/testsupport/database"
)

func TestSampleRepository_Create(t *testing.T) {
	d, release := db.GetDB()
	defer release()

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		sample *entity.Sample
	}
	type wants struct {
		sample *entity.Sample
		isErr  bool
	}

	_fields := fields{DB: d.Session()}
	runFunc := func(t *testing.T, fields fields, args args, wants wants) {
		r := &SampleRepository{DB: fields.DB}
		got, err := r.Create(args.sample)

		if !cmp.Equal(*got, *wants.sample, cmpopts.IgnoreFields(*got, "CreatedAt", "UpdatedAt")) {
			t.Errorf("SampleRepository.Create() %v, want %v", got, wants.sample)
		}
		if (err != nil) != wants.isErr {
			t.Errorf("SampleRepository.Create() error = %v, wantErr %v", err, wants.isErr)
			return
		}
	}

	t.Run("valid entity", func(t *testing.T) {
		en := entity.Sample{
			Title:   "title",
			Content: gox.NewString("content"),
			UserID:  "dummyuserid",
		}
		runFunc(t, _fields, args{sample: &en}, wants{sample: &en, isErr: false})
	})
}
