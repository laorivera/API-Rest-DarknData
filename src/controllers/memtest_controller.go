package controllers

import (
	"builder/src/models"
	"fmt"
	"runtime"
	"testing"
)

// /// mem test ia script

func printMemoryUsage(step string) {

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	fmt.Printf("[%s] Memory: %.2f MB\n", step, float64(m.Alloc)/1024/1024)
}

func TestMemoryUsage(t *testing.T) {

	runtime.GC()
	printMemoryUsage("START - Before loading items")

	_ = NewItemManager()

	printMemoryUsage("AFTER - Items loaded")

	// Do some operations to see memory in action
	im := NewItemManager()
	selection := models.Selection{
		ItemSlot: models.ItemSlot{
			Head:  models.ItemSelect{Name: "Armet"},
			Chest: models.ItemSelect{Name: "Doublet"},
			Foot:  models.ItemSelect{Name: "LightfootBoots"},
			Hands: models.ItemSelect{Name: "RivetedGloves"},
			Pants: models.ItemSelect{Name: "LooseTrousers"},
			Back:  models.ItemSelect{Name: "RadiantCloak"},
		},
	}

	_ = im.ArmorsByName(selection)
	printMemoryUsage("AFTER - One armor lookup")

	_ = im.WeaponsByName(selection)
	printMemoryUsage("AFTER - Weapon lookup")

	_ = im.AccesoriesByName(selection)
	printMemoryUsage("AFTER - Accessory lookup")
}
