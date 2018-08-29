package spellcheck

import (
	"log"

	"github.com/client9/gospell"
	"github.com/elaugier/lpdp/pkg/config"
)

//Spellcheck ...
func Spellcheck(word string) bool {
	configuration, err := config.Get()
	if err != nil {
		//logger.Fatal(err)
	}
	gs, err := gospell.NewGoSpell(configuration.GetString("gospell.aff"), configuration.GetString("gospell.dic"))
	if err != nil {
		log.Fatalf("Unable to load dictionary: %s", err)
	}
	return gs.Spell(word)
}
