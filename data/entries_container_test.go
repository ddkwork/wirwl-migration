package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wirwl/log"
)

func TestThatEntriesContainerHasDataAfterLoading(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	err := container.LoadData()
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, GetExampleComicEntries(), container.entries[comicsEntryType])
	assert.Equal(t, GetExampleVideoEntries(), container.entries[videoEntryType])
	assert.Equal(t, GetExampleMusicEntries(), container.entries[musicEntryType])
}

func TestThatEntriesContainerCreatesOutputDataFile(t *testing.T) {
	testDbPath, cleanup := getTempDbPath()
	defer cleanup()
	container := NewEntriesContainer(NewSampleTestDataProvider(testDbPath))
	err := container.SaveData()
	if err != nil {
		log.Fatal(err)
	}
	assert.FileExists(t, testDbPath)
}

func TestThatEntriesContainerReturnsAnErrorOnDataLoadFailure(t *testing.T) {
	container := NewEntriesContainer(NewAlwaysFailingProvider())
	err := container.LoadData()
	assert.NotNil(t, err)
}

func TestThatEntriesContainerReturnsAnErrorOnDataSaveFailuer(t *testing.T) {
	container := NewEntriesContainer(NewAlwaysFailingProvider())
	err := container.SaveData()
	assert.NotNil(t, err)
}

func TestThatAddingNewEntryTypeWorks(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	typeToAdd := EntryType{
		Name:                  "added entry",
		CompletionElementName: "test element",
		ImageQuery:            "entry query",
	}
	_ = container.AddEntryType(typeToAdd)
	assert.NotNil(t, container.entries[typeToAdd])
}

func TestThatErrorIsReturnedWhenTryingToAddEntryTypeWithTheSameName(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	typeToAdd := EntryType{
		Name:                  "added entry",
		CompletionElementName: "test element",
		ImageQuery:            "entry query",
	}
	typeToAdd2 := EntryType{
		Name:                  "added entry",
		CompletionElementName: "test element2",
		ImageQuery:            "entry query2",
	}
	_ = container.AddEntryType(typeToAdd)
	err := container.AddEntryType(typeToAdd2)
	assert.Contains(t, err.Error(), "Entry type with name 'added entry' already exists")
}

func TestThatErrorIsReturnedWhenTryingToAddEntryTypeWithEmptyName(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	typeToAdd := EntryType{
		Name:                  "",
		CompletionElementName: "test element",
		ImageQuery:            "entry query",
	}
	err := container.AddEntryType(typeToAdd)
	assert.Contains(t, err.Error(), "Cannot add entry type with an empty name")
}

func TestThatWhenEntryTypeExistsItIsRemoved(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	typeToAdd := EntryType{
		Name:                  "test type",
		CompletionElementName: "test element",
		ImageQuery:            "entry query",
	}
	_ = container.AddEntryType(typeToAdd)
	_ = container.DeleteEntryType("test type")
	assert.Equal(t, 0, len(container.entries))
}

func TestThatWhenEntryTypeDoesNotExistErrorIsReturned(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	err := container.DeleteEntryType("non existing entry type")
	assert.Contains(t, err.Error(), "Cannot delete an entry type with name 'non existing entry type' as there is no such type")
}

func TestThatItIsPossibleToUpdateEntryType(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	typeToAdd := EntryType{
		Name:                  "test type",
		CompletionElementName: "test element",
		ImageQuery:            "entry query",
	}
	_ = container.AddEntryType(typeToAdd)
	typeToUpdateWith := EntryType{
		Name:                  "new type",
		CompletionElementName: "another element",
		ImageQuery:            "some other query",
	}
	_ = container.UpdateEntryType("test type", typeToUpdateWith)
	assert.Nil(t, container.entries[typeToAdd])
	assert.NotNil(t, container.entries[typeToUpdateWith])
}

func TestThatErrorIsReturnedWhenTryingToUpdateTypeToTypeWithEmptyName(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	typeToAdd := EntryType{
		Name:                  "test type",
		CompletionElementName: "test element",
		ImageQuery:            "entry query",
	}
	_ = container.AddEntryType(typeToAdd)
	typeToUpdateWith := EntryType{
		Name:                  "",
		CompletionElementName: "another element",
		ImageQuery:            "some other query",
	}
	err := container.UpdateEntryType("test type", typeToUpdateWith)
	assert.Contains(t, err.Error(), "Cannot update entry type with name 'test type' to type with an empty name")
	assert.NotNil(t, container.entries[typeToAdd])
	assert.Nil(t, container.entries[typeToUpdateWith])
}

func TestThatErrorIsReturnWhenTryingToUpdateNonExistentType(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	typeToUpdateWith := EntryType{
		Name:                  "some name",
		CompletionElementName: "another element",
		ImageQuery:            "some other query",
	}
	err := container.UpdateEntryType("test type", typeToUpdateWith)
	assert.Contains(t, err.Error(), "Cannot update entry type 'test type' as no such type exists")
	assert.Nil(t, container.entries[typeToUpdateWith])
}

func TestThatItIsPossibleToRetrieveExistingEntryType(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	addedType := EntryType{
		Name:                  "test type",
		CompletionElementName: "test element",
		ImageQuery:            "entry query",
	}
	_ = container.AddEntryType(addedType)
	retrievedType, _ := container.EntryTypeWithName("test type")
	assert.Equal(t, addedType, retrievedType)
}

func TestThatErrorIsReturnedWhenTryingToRetrieveNonExistentEntryType(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	_, err := container.EntryTypeWithName("test type")
	assert.Contains(t, err.Error(), "Cannot retrieve entry type with name 'test type' as such entry type doesn't exist")
}

func TestThatEntriesGroupedByTypeAreProperlyReturned(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	_ = container.LoadData()
	groupedEntries := container.EntriesGroupedByType()
	assert.Equal(t, GetExampleComicEntries(), groupedEntries[comicsEntryType])
	assert.Equal(t, GetExampleVideoEntries(), groupedEntries[videoEntryType])
	assert.Equal(t, GetExampleMusicEntries(), groupedEntries[musicEntryType])
}

func TestThatChangeCallbackFunctionIsCalledOnEveryChangeForEveryListener(t *testing.T) {
	function1Called := false
	function2Called := false
	function1 := func() {
		function1Called = true
	}
	function2 := func() {
		function2Called = true
	}
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	container.SubscribeToChanges(function1)
	container.SubscribeToChanges(function2)

	err := container.AddEntryType(comicsEntryType)
	if err != nil {
		log.Fatal(err)
	}
	assert.True(t, function1Called)
	assert.True(t, function2Called)
	function1Called = false
	function2Called = false

	err = container.UpdateEntryType(comicsEntryType.Name, videoEntryType)
	if err != nil {
		log.Fatal(err)
	}
	assert.True(t, function1Called)
	assert.True(t, function2Called)
	function1Called = false
	function2Called = false

	err = container.DeleteEntryType(videoEntryType.Name)
	if err != nil {
		log.Fatal(err)
	}
	assert.True(t, function1Called)
	assert.True(t, function2Called)
	function1Called = false
	function2Called = false
}

func TestThatItIsPossibleToUpdateEntryTypeWhenBothUpdatedAndEntryToUpdateHaveTheSameName(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	err := container.AddEntryType(comicsEntryType)
	if err != nil {
		log.Fatal(err)
	}
	err = container.UpdateEntryType(comicsEntryType.Name, EntryType{
		Name:                  comicsEntryType.Name,
		CompletionElementName: "updated name",
		ImageQuery:            "updated query",
	})
	if err != nil {
		log.Fatal(err)
	}
	updatedEntryType, err := container.EntryTypeWithName(comicsEntryType.Name)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, comicsEntryType.Name, updatedEntryType.Name)
	assert.Equal(t, "updated name", updatedEntryType.CompletionElementName)
	assert.Equal(t, "updated query", updatedEntryType.ImageQuery)
}

func TestThatEntryTypeIsNotRemovedWhenItIsUpdatedWithTheExactSameDataAsItHasCurrently(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	err := container.AddEntryType(comicsEntryType)
	if err != nil {
		log.Fatal(err)
	}
	err = container.UpdateEntryType(comicsEntryType.Name, comicsEntryType)
	if err != nil {
		log.Fatal(err)
	}
	updatedEntryType, err := container.EntryTypeWithName(comicsEntryType.Name)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotNil(t, updatedEntryType)
}

func TestThatContainerReturnsCorrectAmountOfEntryTypes(t *testing.T) {
	container := NewEntriesContainer(NewSampleTestDataProvider(""))
	err := container.LoadData()
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 3, container.AmountOfTypes())
}
