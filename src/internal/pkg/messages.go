package pkg

const (
	CommandStart  = "start"
	CommandSet    = "set"
	CommandGet    = "get"
	CommandDelete = "del"

	MsgStart = `/set \- добавить пароль для сервиса
	/get \- получить пароль для сервиса
	/del \- удалить пароль для сервиса`
	MsgSet            = "Введите название сервиса и пароль через пробел"
	MsgGet            = "Введите название сервиса"
	MsgDel            = "Введите название сервиса"
	MsgUnknownCommand = "Неизвестная команда"

	MsgSuccessSet = "Установлен пароль *%s* для сервиса *%s*"
	MsgSuccessDel = "Пароль для сервиса *%s* удален"

	MsgErrorSetInvalidArgs = `Некорректный формат записи`
	MsgErrorSetTooLong     = `Длина параметров слишком большая`
	MsgErrorGet            = "Пароль для данного сервиса не найден"
	MsgErrorDel            = "Запись для данного сервиса не найдена"
	MsgRepositoryError     = "Внутренняя ошибка базы данных"
	MsgErrorUnknown        = `Неизвестная ошибка`
)
