package services

var (
	ItemService ItemServiceInterface = &itemService{}
)

type ItemServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemService struct {
}

func (s *itemService) GetItem() {

}

func (s *itemService) SaveItem() {

}
