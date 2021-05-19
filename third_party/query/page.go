package query

func ToOffsetLimit(paginate *bool, page, size *int) (int, int) {
	offset := 0
	limit := 5

	if paginate != nil {
		if *paginate == true {
			return 0, 0
		}
	}

	if page != nil && size != nil {
		offset = *page * (*size - 1)
		limit = *size
	} else if page != nil && size == nil {
		offset = limit * (*page - 1)
	} else if size != nil && page == nil {
		limit = *size
	}

	return offset, limit
}
