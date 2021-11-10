package contract

type MeowUsecase interface {
	Create(body string) (int, error)
}
