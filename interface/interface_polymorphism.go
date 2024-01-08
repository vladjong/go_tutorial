package gotutorial

import "errors"

// интерфейс для работы с хранилищем
//
//go:generate go run github.com/vektra/mockery/v2@v2.39.1 --name Storage
type Storage interface {
	User(id string) string
	Add(id, name string)
}

type Cache struct {
	storage map[string]string
}

func (c *Cache) User(id string) string {
	return c.storage[id]
}

func (c *Cache) Add(id, name string) {
	c.storage[id] = name
}

// здесь должна быть логика работы с Postgres
type Postgres struct{}

func (p *Postgres) User(id string) string {
	return "name"
}

func (p *Postgres) Add(id, name string) {
}

// бизнес-логика нашего приложения
type UseCase struct {
	repo Storage
}

// пример полиморфизма. Передаем параметр repo в виде интерфейса Storage
func New(repo Storage) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CheckUser(id string) error {
	user := u.repo.User(id)
	if user == "" {
		return errors.New("user don't exist")
	}

	return nil
}

func main() {
	cache := &Cache{
		storage: map[string]string{},
	}

	postgres := &Postgres{}

	// Иницилизация useCase и проверка работоспособности
	useCaseOne := New(cache)
	useCaseOne.CheckUser("1")

	useCaseTwo := New(postgres)
	useCaseTwo.CheckUser("1")
}
