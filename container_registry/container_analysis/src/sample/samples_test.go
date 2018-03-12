package main

import (
	"os"
	"testing"
	"reflect"
	"time"
	"strconv"
	sample "sample"
	containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1alpha1"
)

type TestVariables struct {
	testName string
	noteId string
	imageUrl string
	projectId string
	noteObj *containeranalysispb.Note
}

//Run before each test. Creates a set of useful variables
func setup(t *testing.T) (TestVariables){
	//get test name
	value := reflect.ValueOf(*t)
	name := value.FieldByName("name").String()
	t.Log("SETUP " + name)
	//get current timestamp
	timestamp:= strconv.Itoa(int(time.Now().Unix()))

	//create variables used by tests
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	noteId := "note-" + timestamp + name
	imageUrl := "www." + timestamp + name + ".com"
	noteObj, _ := sample.CreateNote(noteId, projectId)

	v := TestVariables{name, noteId, imageUrl, projectId, noteObj}
	return v
}

//Run after each test
//Removes any unneeded resources allocated for test
func teardown(t *testing.T, v TestVariables) {
	t.Log("TEARDOWN " + v.testName)
	err := sample.DeleteNote(v.noteId, v.projectId)
	if err != nil {
		t.Log(err)
	}
}


func TestCreateNote(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestDeleteNote(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestUpdateNote(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestCreateOccurrence(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestDeleteOccurrence(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestUpdateOccurrence(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestOccurrencesForImage(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestOccurrencesForNote(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}

func TestPubSub(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}
