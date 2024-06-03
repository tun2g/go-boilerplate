package utils

import (
	dto "fist-app/src/shared/dto"
)

func GeneratePaginationResult(total int, data []interface{}, options *dto.PageOptionsDto) (*dto.PageDto) {
	pageMetaDto := dto.NewPageMetaDto(options, total)
	return dto.NewPageDto(data, pageMetaDto)
}
