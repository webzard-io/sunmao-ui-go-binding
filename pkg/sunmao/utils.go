package sunmao

func componentAttachedToSlot(builder BaseComponentBuilder) bool {
	for _, t := range builder.ValueOf().Traits {
		if t.Type == "core/v1/slot" {
			return true
		}
	}
	return false
}
