package test

import (
	"MyPIPE/domain/model"
	"testing"
)

func TestCommentID(t *testing.T) {
	trueCases := []struct {
		CommentID uint64
	}{
		{CommentID: 10},
	}

	falseCases := []struct {
		CommentID uint64
	}{
		{CommentID: 0},
	}

	for _, trueCase := range trueCases {
		_, newCommentErr := model.NewCommentID(trueCase.CommentID)
		if newCommentErr != nil {
			t.Error("CommentID TrueCase Error.")
		}
	}

	for _, falseCase := range falseCases {
		_, newCommentErr := model.NewCommentID(falseCase.CommentID)
		if newCommentErr == nil {
			t.Error("CommentID FalseCase Error.")
		}
	}
}

func TestCommentBody(t *testing.T) {
	trueCases := []struct {
		CommentBody string
	}{
		{CommentBody: "こんにちは"},
	}

	falseCases := []struct {
		CommentBody string
	}{
		{CommentBody: ""},
		{CommentBody: "11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111" +
			"11111111111111111111111111111111111111111111111111"},
	}

	for _, trueCase := range trueCases {
		_, newCommentErr := model.NewCommentBody(trueCase.CommentBody)
		if newCommentErr != nil {
			t.Error("CommentBody TrueCase Error.")
		}
	}

	for _, falseCase := range falseCases {
		_, newCommentErr := model.NewCommentBody(falseCase.CommentBody)
		if newCommentErr == nil {
			t.Error("CommentBody FalseCase Error.")
		}
	}
}
