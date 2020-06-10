package services2

var (
	BeerService beerServiceInterface = &beerService{}
)

type beerService struct {}

type beerServiceInterface interface {
	GimmeABeer()
}

func (bs *beerService) GimmeABeer() {

}
