package main

import (
	"fmt"
	"golang.org/x/net/context"
	containeranalysis "cloud.google.com/go/devtools/containeranalysis/apiv1alpha1"
	containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1alpha1"
)


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

func createOccurrence(imageUrl, parentNoteId, projectId string) (*containeranalysispb.Occurrence, error){
	return nil, nil
}
func updateNote(updated *containeranalysispb.Note, noteId, projectId string) (error){
	return nil
}

func updateOccurrence(updated *containeranalysispb.Occurrence, occurrenceName string) (error){
	return nil
}

func deleteNote(noteId, projectId string) (error){
	return nil
}

func deleteOccurrence(occurrenceName string) (error){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	req := &containeranalysispb.DeleteOccurrenceRequest{Name: occurrenceName}
	return c.DeleteOccurrence(ctx, req)
}

func getNote(noteId, projectId string) (*containeranalysispb.Note, error){
	return nil, nil
}

func getOccurrence(occurrenceName string) (*containeranalysispb.Occurrence, error){
	return nil, nil
}

func getDiscoveryInfo(imageUrl, projectId string) (error){
	return nil
}

func getOccurrencesForNote(noteId, projectId string) (int, error){
	return 0, nil
}

func getOccurrencesForImage(imageUrl, projectId string) (int, error){
	return 0, nil
}

func pubsub(subscriptionId string, timeout int, projectId string) (int, error){
	return 0, nil
}

func createOccurrenceSubscription(subscriptionId, projectId string) (error){
	return nil
}

func main() {
	fmt.Println("hello world")
	_, err := createNote("test3", "sanche-testing-project")
	if err != nil {
		fmt.Println(err)
	}
}
