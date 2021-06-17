package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/yerassyldanay/invest/model"
	"testing"
)

func TestDocumentGet(t *testing.T) {
	var project = HelperGetAnyProject(t)
	var document = model.Document{ProjectId: project.Id, Step: project.Step}

	documents, err := document.OnlyGetDocumentsByStepsAndProjectId(project.Id, []interface{}{1}, model.GetDB())

	// check - step 1
	require.NoError(t, err)
	require.NotZero(t, len(documents))
	require.Condition(t, func() (bool) { return len(documents) > 0 })

	step1 := len(documents)

	documents, err = document.OnlyGetDocumentsByStepsAndProjectId(project.Id, []interface{}{2}, model.GetDB())

	// check - step 2
	require.NoError(t, err)
	require.NotZero(t, len(documents))
	require.Condition(t, func() (bool) { return len(documents) > 0 })

	step2 := len(documents)

	documents, err = document.OnlyGetDocumentsByProjectId(project.Id, model.GetDB())

	// check - step 1 & 2
	require.NoError(t, err)
	require.Equal(t, len(documents), step1 + step2)
}

