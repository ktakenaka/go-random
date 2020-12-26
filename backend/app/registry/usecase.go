// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"github.com/google/wire"

	"github.com/ktakenaka/go-random/backend/app/config"
	"github.com/ktakenaka/go-random/backend/app/domain/repository"
	"github.com/ktakenaka/go-random/backend/app/domain/service"
	"github.com/ktakenaka/go-random/backend/app/external/database"
	"github.com/ktakenaka/go-random/backend/app/interface/adaptor/restclient"
	"github.com/ktakenaka/go-random/backend/app/interface/persistence/mysql"
	"github.com/ktakenaka/go-random/backend/app/usecase"
	infradb "github.com/ktakenaka/go-random/backend/pkg/infra/database"
)

func InitializeSampleUsecase() *usecase.SampleUsecase {
	wire.Build(SampleUsecaseSet)
	return &usecase.SampleUsecase{}
}

func InitializeSignInUsecase() *usecase.SignInUsecase {
	wire.Build(SignInUsecaseSet)
	return &usecase.SignInUsecase{}
}

var (
	SampleUsecaseSet = wire.NewSet(
		sampleRepositorySet,
		service.NewSampleService,
		usecase.NewSampleUsecase,
	)

	sampleRepositorySet = wire.NewSet(
		mysql.NewSampleRepository,
		database.MySQLConnection,
		wire.Bind(new(repository.SampleRepository), new(*mysql.SampleRepository)),
		wire.Bind(new(repository.DBConnection), new(*infradb.DB)),
	)

	SignInUsecaseSet = wire.NewSet(
		googleRepositorySet,
		userRepositorySet,
		usecase.NewSignInUsecase,
	)

	googleRepositorySet = wire.NewSet(
		config.GetGoogleOauthConfig,
		restclient.NewGoogleRepository,
		wire.Bind(new(repository.GoogleRepository), new(*restclient.GoogleRepository)),
	)

	userRepositorySet = wire.NewSet(
		mysql.NewUserRepository,
		database.MySQLConnection,
		wire.Bind(new(repository.UserRepository), new(*mysql.UserRepository)),
	)
)
