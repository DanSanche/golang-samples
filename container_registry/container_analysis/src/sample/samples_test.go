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
	tryLimit int
	sleepTime int
}

//Run before each test. Creates a set of useful variables
func setup(t *testing.T) (TestVariables){
	//get test name
	value := reflect.ValueOf(*t)
	name := value.FieldByName("name").String()
	t.Log("SETUP " + name)
	//get current timestamp
	timestamp:= strconv.Itoa(int(time.Now().Unix()))
	//set how many times to retry network tasks
	tryLimit := 20
	sleepTime := 1

	//create variables used by tests
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	noteId := "note-" + timestamp + name
	imageUrl := "www." + timestamp + name + ".com"
	noteObj, _ := sample.CreateNote(noteId, projectId)


	v := TestVariables{name, noteId, imageUrl, projectId, noteObj, tryLimit, sleepTime}
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

//test equality between two values
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%s != %s", a, b)					
	}
}


func TestCreateNote(t *testing.T){
	v := setup(t)

	newNote, err := sample.GetNote(v.noteId, v.projectId)
	if err != nil{
		t.Fatal(err)
	} else if newNote != nil {
		assertEqual(t, newNote.Name, v.noteObj.Name)
	} else {
		t.Error("both outputs == nil")
	}

	teardown(t, v)
}

func TestDeleteNote(t *testing.T){
	v := setup(t)
	
	err := sample.DeleteNote(v.noteId, v.projectId)
	if err != nil{
		t.Error(err)
	}
	deleted, err := sample.GetNote(v.noteId, v.projectId)
	if deleted != nil || err == nil {
		t.Error("GetNote succeeded after DeleteNote")
	}
	
	teardown(t, v)
}

func TestUpdateNote(t *testing.T){
	v := setup(t)

	description := "updated"
	v.noteObj.ShortDescription = description
	err := sample.UpdateNote(v.noteObj, v.noteId, v.projectId)
	if err != nil {
		t.Error(err)
	}
	updated, err := sample.GetNote(v.noteId, v.projectId)
	if err != nil{
		t.Fatal(err)
	} else if updated != nil {
		assertEqual(t, updated.ShortDescription, description)
	} else {
		t.Error("Could not find updated note. No error returned")
	}

	teardown(t, v)
}

func TestCreateOccurrence(t *testing.T){
	v := setup(t)
	
	created, err := sample.CreateOccurrence(v.imageUrl, v.noteId, v.projectId)
	if err != nil {
		t.Error(err)
	} else if err == nil && created == nil {
		t.Error("Both CreateOccurrence outputs == nil")
	} else {
		retrieved, err := sample.GetOccurrence(created.Name)
		if err != nil{
			t.Error(err)
		} else if retrieved != nil {
			assertEqual(t, retrieved.Name, created.Name)
		} else {
			t.Error("Could not find updated note. No error returned")
		}
	}

	teardown(t, v)
}

func TestDeleteOccurrence(t *testing.T){
	v := setup(t)

	created, err := sample.CreateOccurrence(v.imageUrl, v.noteId, v.projectId)
	if err != nil {
		t.Error(err)
	} else if err == nil && created == nil {
		t.Error("Both CreateOccurrence outputs == nil")
	} else {
		err = sample.DeleteOccurrence(created.Name)	
		if err != nil {
			t.Error(err)
		}
		deleted, err := sample.GetOccurrence(created.Name)
		if deleted != nil || err == nil {
			t.Error("GetOccurrence succeeded after DeleteOccurrence")
		}
	}

	teardown(t, v)
}

func TestUpdateOccurrence(t *testing.T){
	v := setup(t)
	

	created, err := sample.CreateOccurrence(v.imageUrl, v.noteId, v.projectId)
	if err != nil {
		t.Error(err)
	} else if err == nil && created == nil {
		t.Error("Both CreateOccurrence outputs == nil")
	} else {
		newType := "updated"
		vul := new(containeranalysispb.VulnerabilityType_VulnerabilityDetails)
		vul.Type = newType
		vulDetails := containeranalysispb.Occurrence_VulnerabilityDetails{VulnerabilityDetails:vul}
		occurrence := containeranalysispb.Occurrence{NoteName: created.NoteName, ResourceUrl: created.ResourceUrl, Details:&vulDetails}
		sample.UpdateOccurrence(&occurrence, created.Name)
		retrieved, err := sample.GetOccurrence(created.Name)
		if err != nil{
			t.Error(err)
		} else if retrieved != nil {
			retrievedDetails := retrieved.GetVulnerabilityDetails()
			assertEqual(t, retrievedDetails.Type, newType)
		} else {
			t.Error("Could not find updated occurrence. No error returned")
		}
	}
	teardown(t, v)
}

func TestOccurrencesForImage(t *testing.T){
	v := setup(t)
	newCount := 0
	tries := 0
	
	origCount, err := sample.GetOccurrencesForImage(v.imageUrl, v.projectId)
	if err != nil {
		t.Fatal(err)
	}
	created, _ := sample.CreateOccurrence(v.imageUrl, v.noteId, v.projectId)
	for newCount != 1 && tries < v.tryLimit{
		newCount, _ = sample.GetOccurrencesForImage(v.imageUrl, v.projectId)
		tries = tries + 1
		time.Sleep(time.Second*time.Duration(v.sleepTime))
	}
	assertEqual(t, origCount, 0)
	assertEqual(t, newCount, 1)

	// clean up
	sample.DeleteOccurrence(created.Name)
	teardown(t, v)
}

func TestOccurrencesForNote(t *testing.T){
	v := setup(t)
	newCount := 0
	tries := 0
	
	origCount, err := sample.GetOccurrencesForNote(v.noteId, v.projectId)
	if err != nil {
		t.Fatal(err)
	}
	created, _ := sample.CreateOccurrence(v.imageUrl, v.noteId, v.projectId)
	for newCount != 1 && tries < v.tryLimit{
		newCount, _ = sample.GetOccurrencesForNote(v.noteId, v.projectId)
		tries = tries + 1
		time.Sleep(time.Second*time.Duration(v.sleepTime))
	}
	assertEqual(t, origCount, 0)
	assertEqual(t, newCount, 1)

	// clean up
	sample.DeleteOccurrence(created.Name)
	teardown(t, v)
}

func TestPubSub(t *testing.T){
	v := setup(t)
	t.Errorf("failed")
	teardown(t, v)
}
