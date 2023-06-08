package repository

type Sqlite struct{}

func NewRepo() *Sqlite {
	return &Sqlite{}
}
