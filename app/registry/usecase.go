// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"github.com/google/wire"
	"github.com/ktakenaka/go-random/app/domain/repository"
	"github.com/ktakenaka/go-random/app/domain/service"
	"github.com/ktakenaka/go-random/app/external/database"
	"github.com/ktakenaka/go-random/app/interface/persistence/mysql"
	"github.com/ktakenaka/go-random/app/usecase"
)

func InitializeSampleUsecase() *usecase.SampleUsecase {
	wire.Build(SampleUsecaseSet)
	return &usecase.SampleUsecase{}
}

var (
	SampleUsecaseSet = wire.NewSet(
		database.MySQLConnection,
		sampleRepositorySet,
		service.NewSampleService,
		usecase.NewSampleUsecase,
	)
)

var (
	sampleRepositorySet = wire.NewSet(
		mysql.NewSampleRepository,
		mysql.NewTransactionManager,
		wire.Bind(new(repository.SampleRepository), new(*mysql.SampleRepository)),
		wire.Bind(new(repository.TransactionManager), new(*mysql.TransactionManager)),
	)
)
