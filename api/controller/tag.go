package controller

import (
	"fmt"
	"github.com/moira-alert/moira-alert"
	"github.com/moira-alert/moira-alert/api"
	"github.com/moira-alert/moira-alert/api/dto"
)

func GetAllTagsAndSubscriptions(database moira.Database) (*dto.TagsStatistics, *api.ErrorResponse) {
	//todo работает медлено
	tagsNames, err := database.GetTagNames()
	if err != nil {
		return nil, api.ErrorInternalServer(err)
	}

	tagsStatistics := dto.TagsStatistics{
		List: make([]dto.TagStatistics, 0, len(tagsNames)),
	}

	for _, tagName := range tagsNames {
		tagStat := dto.TagStatistics{}
		tagStat.TagName = tagName
		tagStat.Subscriptions, err = database.GetTagsSubscriptions([]string{tagName})
		if err != nil {
			return nil, api.ErrorInternalServer(err)
		}
		tagStat.Triggers, err = database.GetTagTriggerIds(tagName)
		if err != nil {
			return nil, api.ErrorInternalServer(err)
		}
		tagsStatistics.List = append(tagsStatistics.List, tagStat)
	}
	return &tagsStatistics, nil
}

func GetAllTags(database moira.Database) (*dto.TagsData, *api.ErrorResponse) {
	tagsNames, err := database.GetTagNames()
	if err != nil {
		return nil, api.ErrorInternalServer(err)
	}

	tagsData := &dto.TagsData{
		TagNames: tagsNames,
	}

	return tagsData, nil
}

func DeleteTag(database moira.Database, tagName string) (*dto.MessageResponse, *api.ErrorResponse) {
	triggerIds, err := database.GetTagTriggerIds(tagName)
	if err != nil {
		return nil, api.ErrorInternalServer(err)
	}

	if len(triggerIds) > 0 {
		return nil, api.ErrorInvalidRequest(fmt.Errorf("This tag is assigned to %v triggers. Remove tag from triggers first", len(triggerIds)))
	} else {
		if err = database.DeleteTag(tagName); err != nil {
			return nil, api.ErrorInternalServer(err)
		}
	}
	return &dto.MessageResponse{Message: "tag deleted"}, nil
}
