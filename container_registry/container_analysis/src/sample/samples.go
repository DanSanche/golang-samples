package main

import (
	"fmt"
	"golang.org/x/net/context"
	containeranalysis "cloud.google.com/go/devtools/containeranalysis/apiv1alpha1"
	containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1alpha1"
)

func deleteOccurrence(occurrenceName string) (error){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	req := &containeranalysispb.DeleteOccurrenceRequest{Name: occurrenceName}
	return c.DeleteOccurrence(ctx, req)
}


func createNote(noteId, projectId string) (*containeranalysispb.Note, error) {
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	parent := containeranalysis.ProjectPath(projectId)
	noteVulType := containeranalysispb.VulnerabilityType{}
	noteType := containeranalysispb.Note_VulnerabilityType{&noteVulType}
	note := containeranalysispb.Note{NoteType:&noteType}
	req := &containeranalysispb.CreateNoteRequest{Parent:parent, NoteId:noteId, Note:&note}
	return c.CreateNote(ctx, req)
}


func main() {
	fmt.Println("hello world")
	_, err := createNote("test3", "sanche-testing-project")
	if err != nil {
		fmt.Println(err)
	}
}
