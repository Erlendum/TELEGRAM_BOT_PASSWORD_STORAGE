package modes

import (
	"github.com/pkg/errors"
	"log"
	"src/config"
	"src/internal/repositories"
	"src/internal/repositories/redis_repositories"
	"src/internal/services"
	"src/internal/services/services_implementation"
	"src/internal/telegram"
)

type App struct {
	Config       config.Config
	repositories *appRepositoryFields
	services     *appServiceFields
	bot          *telegram.Bot
}

type appRepositoryFields struct {
	passwordRecordRepository repositories.PasswordRecordRepository
}

type appServiceFields struct {
	passwordRecordService services.PasswordRecordService
}

func (a *App) initRepositories() *appRepositoryFields {
	fields, err := redis_repositories.CreateRedisRepositoryFields("config.json", "./config")
	if err != nil {
		return nil
	}
	err = a.Config.ParseConfig("config.json", "./config")
	if err != nil {
		return nil
	}
	f := &appRepositoryFields{
		passwordRecordRepository: redis_repositories.NewPasswordRecordRedisRepository(fields),
	}

	return f
}

func (a *App) initBot() *telegram.Bot {
	botAPI, err := a.Config.Bot.Init()
	if err != nil {
		return nil
	}

	return telegram.NewBot(botAPI, a.services.passwordRecordService)

}

func (a *App) initServices(r *appRepositoryFields) *appServiceFields {
	u := &appServiceFields{
		passwordRecordService: services_implementation.NewPasswordRecordServiceImplementation(r.passwordRecordRepository),
	}

	return u
}

func (a *App) Init() error {
	a.repositories = a.initRepositories()
	if a.repositories == nil {
		return errors.Errorf("init repositories failed")
	}
	a.services = a.initServices(a.repositories)
	if a.services == nil {
		return errors.Errorf("init services failed")
	}

	a.bot = a.initBot()
	if a.bot == nil {
		return errors.Errorf("init bot failed")
	}
	return nil
}

func (a *App) Run() {
	err := a.Init()
	if err != nil {
		log.Fatal(err)
	}

	if true {
		err := a.bot.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

}
