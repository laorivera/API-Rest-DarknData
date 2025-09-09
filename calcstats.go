package main

func ProcessSelectionEnchantments(selection Selection) Computed_Stats {
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

	// Compute the final stats using your existing function
	return EnchantComputedOthers(allEnchantments)
}

func ProcessBaseAttributeEnchantments(selection Selection) Stats {
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

	return setEnchantStats(allEnchantments)
}
