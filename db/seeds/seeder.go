package seeds

import (
	"gorm.io/gorm"
	"log"
	"reflect"
)

type Seed struct {
	db *gorm.DB
}

func seed(s Seed, seedMethod string) {
	m := reflect.ValueOf(s).MethodByName(seedMethod)

	if !m.IsValid() {
		log.Fatal("No method called ", seedMethod)
	}

	log.Println("Seeding"+seedMethod, "...")
	m.Call(nil)
	log.Println("Seeding "+seedMethod, "successful")
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeders...")
		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			seed(s, method.Name)
		}
	}

	for _, item := range seedMethodNames {
		seed(s, item)
	}
}
