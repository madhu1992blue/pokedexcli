package main
import "time"
import "github.com/madhu1992blue/pokedexcli/internal/pokeapi"
type config struct {
	pokeclient *pokeapi.Client
	NextLocationAreaURL *string
	PreviousLocationAreaURL *string
}
func main() {

	cfg := &config{
		pokeclient: pokeapi.NewClient("https://pokeapi.co/api/v2", 5 * time.Second, 5* time.Minute ),
	}
	startRepl(cfg)
}
