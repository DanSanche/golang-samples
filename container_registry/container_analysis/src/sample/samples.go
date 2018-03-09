package main

import (
	"fmt"
	"golang.org/x/net/context"
	containeranalysis "cloud.google.com/go/devtools/containeranalysis/apiv1alpha1"
	containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1alpha1"
)

func deleteOccurrence(occurrenceName string){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
		fmt.Println(err)
	}
	req := &containeranalysispb.DeleteOccurrenceRequest{Name: occurrenceName}
	err = c.DeleteOccurrence(ctx, req)
	if err != nil {
		// TODO: Handle error.
		fmt.Println(err)
	}
}

func createNote(noteId, projectId string){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
		fmt.Println(err)
	}
	parent := containeranalysis.ProjectPath(projectId)
	noteVulType := containeranalysispb.VulnerabilityType{}
	noteType := containeranalysispb.Note_VulnerabilityType{&noteVulType}
	note := containeranalysispb.Note{NoteType:&noteType}
	req := &containeranalysispb.CreateNoteRequest{Parent:parent, NoteId:noteId, Note:&note}
	_, err = c.CreateNote(ctx, req)
	if err != nil {
		// TODO: Handle error.
		fmt.Println(err)
	}
}


func main() {
	fmt.Println("hello world")
	createNote("sadfadsfasdadsfasd", "sanche-testing-project")
}
