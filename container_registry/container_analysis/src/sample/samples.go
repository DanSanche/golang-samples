package main

import (
	"fmt"
	"golang.org/x/net/context"
	containeranalysis "cloud.google.com/go/devtools/containeranalysis/apiv1alpha1"
	containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1alpha1"
)

// [START create_note]
//Creates and returns a new note
func CreateNote(noteId, projectId string) (*containeranalysispb.Note, error) {
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
// [END create_note]

// [START create_occurrence]
//Creates and returns a new occurrence
func CreateOccurrence(imageUrl, parentNoteId, projectId string) (*containeranalysispb.Occurrence, error){
	return nil, nil
}
// [END create_occurrence]

// [START update_note]
//Makes an update to an existing note
func UpdateNote(updated *containeranalysispb.Note, noteId, projectId string) (error){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	noteName := containeranalysis.NotePath(projectId, noteId)
	req := &containeranalysispb.UpdateNoteRequest{Name: noteName, Note:updated}
	_, err = c.UpdateNote(ctx, req)
	return err
}
// [END update_note]

// [START update_occurrence]
//Makes an update to an existing occurrence
func UpdateOccurrence(updated *containeranalysispb.Occurrence, occurrenceName string) (error){
	return nil
}
// [END update_occurrence]


// [START delete_note]
//Deletes an existing note from the project
func DeleteNote(noteId, projectId string) (error){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	noteName := containeranalysis.NotePath(projectId, noteId)
	req := &containeranalysispb.DeleteNoteRequest{Name: noteName}
	return c.DeleteNote(ctx, req)
}
// [END delete_note]

// [START delete_occurrence]
//Deletes an existing occurrence from the project
func DeleteOccurrence(occurrenceName string) (error){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	req := &containeranalysispb.DeleteOccurrenceRequest{Name: occurrenceName}
	return c.DeleteOccurrence(ctx, req)
}
// [END delete_occurrence]

// [START get_note]
//Retrieves a note based on its noteId and projectId
func GetNote(noteId, projectId string) (*containeranalysispb.Note, error){
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	noteName := containeranalysis.NotePath(projectId, noteId)
	req := &containeranalysispb.GetNoteRequest{noteName}
	return c.GetNote(ctx, req)
}
// [END get_note]

// [START get_occurrence]
//Retrieves an occurrence based on its name
func GetOccurrence(occurrenceName string) (*containeranalysispb.Occurrence, error){
	return nil, nil
}
// [END get_occurrence]

// [START discovery_info]
//Prints the Discovery occurrence created for a specified image
//This occurrence contains information about the initial scan on the image
func GetDiscoveryInfo(imageUrl, projectId string) (error){
	return nil
}
// [END discovery_info]

// [START occurrences_for_note]
//Retrieves all the occurrences associated with a specified note
func GetOccurrencesForNote(noteId, projectId string) (int, error){
	return 0, nil
}
// [END occurrences_for_note]

// [START occurrences_for_image]
//Retrieves all the occurrences associated with a specified image
func GetOccurrencesForImage(imageUrl, projectId string) (int, error){
	return 0, nil
}
// [END occurrences_for_image]

// [START pubsub]
//Handle incoming occurrences using a pubsub subscription
func Pubsub(subscriptionId string, timeout int, projectId string) (int, error){
	return 0, nil
}

//Creates and returns a pubsub subscription listening to the occurrence topic.
//This topic provides updates when occurrences are modified
func CreateOccurrenceSubscription(subscriptionId, projectId string) (error){
	return nil
}
// [END pubsub]

func main() {
	fmt.Println("hello world")
	_, err := CreateNote("test3", "sanche-testing-project")
	if err != nil {
		fmt.Println(err)
	}
}
