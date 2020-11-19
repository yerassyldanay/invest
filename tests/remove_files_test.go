package tests

import (
	"context"
	"invest/model"
	"testing"
	"time"
)

// remove files in the directory called /documents/analysis/*
func TestRemoveFileLeftAfterAnalysis(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 20)
	model.Remove_files_left_after_analysis_periodically(ctx)
	defer cancel()
}
