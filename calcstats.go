package main

import "fmt"

func ProcessOtherEnchantments(selection Selection) Computed_Stats {
	var allEnchantments []map[string]float64

	// Process all item slots
	items := []ItemSelect{
		selection.ItemSlot.Head,
		selection.ItemSlot.Chest,
		selection.ItemSlot.Foot,
		selection.ItemSlot.Hands,
		selection.ItemSlot.Pants,
		selection.ItemSlot.Back,
		selection.ItemSlot.Necklace,
		selection.ItemSlot.RingOne,
		selection.ItemSlot.RingTwo,
		selection.ItemSlot.WeaponOne,
		selection.ItemSlot.WeaponTwo,
	}

	for _, item := range items {

		enchantSlots := []struct {
			enchantType string
			value       string
		}{
			{item.Enchant.TypeU, item.Enchant.ValueU},
			{item.Enchant.TypeR, item.Enchant.ValueR},
			{item.Enchant.TypeE, item.Enchant.ValueE},
			{item.Enchant.TypeL, item.Enchant.ValueL},
			{item.Enchant.TypeQ, item.Enchant.ValueQ},
		}

		for _, slot := range enchantSlots {
			if slot.enchantType != "" && slot.value != "" {
				enchantMap := Enchantother(slot.enchantType, slot.value)
				if len(enchantMap) > 0 {
					allEnchantments = append(allEnchantments, enchantMap)
				}
			}
		}
	}
	fmt.Println(allEnchantments, "other")
	//fmt.Println(EnchantComputedOthers(allEnchantments), "other")
	// Compute the final stats using your existing function
	return EnchantComputedOthers(allEnchantments)
}

func ProcessBaseEnchantments(selection Selection) Stats {
	var allEnchantments []map[string]int

	items := []ItemSelect{
		selection.ItemSlot.Head,
		selection.ItemSlot.Chest,
		selection.ItemSlot.Foot,
		selection.ItemSlot.Hands,
		selection.ItemSlot.Pants,
		selection.ItemSlot.Back,
		selection.ItemSlot.Necklace,
		selection.ItemSlot.RingOne,
		selection.ItemSlot.RingTwo,
		selection.ItemSlot.WeaponOne,
		selection.ItemSlot.WeaponTwo,
	}

	for _, item := range items {
		enchantSlots := []struct {
			enchantType string
			value       string
		}{
			{item.Enchant.TypeU, item.Enchant.ValueU},
			{item.Enchant.TypeR, item.Enchant.ValueR},
			{item.Enchant.TypeE, item.Enchant.ValueE},
			{item.Enchant.TypeL, item.Enchant.ValueL},
			{item.Enchant.TypeQ, item.Enchant.ValueQ},
		}

		for _, slot := range enchantSlots {
			if slot.enchantType != "" && slot.value != "" {
				enchantMap := Enchantattrib(slot.enchantType, slot.value)
				if len(enchantMap) > 0 {
					allEnchantments = append(allEnchantments, enchantMap)
				}
			}
		}
	}
	fmt.Println(allEnchantments, "base")
	//fmt.Println(setEnchantStats(allEnchantments), "base")
	return setEnchantStats(allEnchantments)
}

/*
func GetAvailableHelmetEnchantments(selection Selection) []string {
	// Get the currently selected helmet enchantment
	selected := getSelectedHelmetEnchantment(selection)

	// If no enchantment is selected, return all
	if selected == "" {
		return helmetEnchantments
	}

	// Filter out the selected enchantment
	available := []string{}
	for _, enchant := range helmetEnchantments {
		if enchant != selected {
			available = append(available, enchant)
		}
	}

	return available
}

// Extract the currently selected helmet enchantment
func getSelectedHelmetEnchantment(selection Selection) string {
	helm := selection.ItemSlot.Head.Enchant

	// Check which enchantment type is selected in the helmet
	if helm.TypeU != "" {
		return helm.TypeU
	}
	if helm.TypeR != "" {
		return helm.TypeR
	}
	if helm.TypeE != "" {
		return helm.TypeE
	}
	if helm.TypeL != "" {
		return helm.TypeL
	}
	if helm.TypeQ != "" {
		return helm.TypeQ
	}
	if helm.TypeA != "" {
		return helm.TypeA
	}

	return "" // No enchantment selected
}

/*

func getUsedFieldNames(obj interface{}) []string {
	var names []string
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if fieldValue.Kind() == reflect.Struct {
			// Recursively check nested structs
			nestedNames := getUsedFieldNames(fieldValue.Interface())
			names = append(names, nestedNames...)
		} else if !fieldValue.IsZero() {
			// Only include fields that have values
			names = append(names, field.Name)
		}
	}
	return names
}

/*
func getAllFieldNamesFlat(slotenchant Enchantmentx) []string {
	var names []string

	t := reflect.TypeOf(slotenchant)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type

		if fieldType.Kind() == reflect.Struct {
			for j := 0; j < fieldType.NumField(); j++ {
				nestedField := fieldType.Field(j)
				names = append(names, nestedField.Name)
			}
		} else {
			names = append(names, field.Name)
		}
	}

	return names
}

/*
func removeBasefromList(slot string, list []string)[]string{
	Im := NewItemManager()

	itemlist := Im.ArmorsBySlot(slot)

	removeAttrs := make(map[string]bool)
	if selection. == 1 {
		removeAttrs["Strength"] = true
	}
	if itemtype.BaseAttribute.Vigor[1] == 1 {
		removeAttrs["Vigor"] = true
	}
	if itemtype.BaseAttribute.Agility[1] == 1 {
		removeAttrs["Agility"] = true
	}
	if itemtype.BaseAttribute.Dexterity[1] == 1 {
		removeAttrs["Dexterity"] = true
	}
	if itemtype.BaseAttribute.Will[1] == 1 {
		removeAttrs["Will"] = true
	}
	if itemtype.BaseAttribute.Knowledge[1] == 1 {
		removeAttrs["Knowledge"] = true
	}
	if itemtype.BaseAttribute.Resourcefulness[1] == 1 {
		removeAttrs["Resourcefulness"] = true
	}

	result := make([]string, 0, len(enchantmentlist))
	for _, enchant := range enchantmentlist {
		if !removeAttrs[enchant] {
			result = append(result, enchant)
		}
	}
	return result
}
*/
