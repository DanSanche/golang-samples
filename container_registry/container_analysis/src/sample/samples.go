package main

import (
	"fmt"
	"sync"
	"time"

	containeranalysis "cloud.google.com/go/devtools/containeranalysis/apiv1alpha1"
	pubsub "cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
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
	note := containeranalysispb.Note{NoteType: &noteType}
	req := &containeranalysispb.CreateNoteRequest{Parent: parent, NoteId: noteId, Note: &note}
	return c.CreateNote(ctx, req)
}

// [END create_note]

// [START create_occurrence]
//Creates and returns a new occurrence
func CreateOccurrence(imageUrl, parentNoteId, projectId string) (*containeranalysispb.Occurrence, error) {
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	parent := containeranalysis.NotePath(projectId, parentNoteId)
	project := containeranalysis.ProjectPath(projectId)
	vulDetails := containeranalysispb.Occurrence_VulnerabilityDetails{new(containeranalysispb.VulnerabilityType_VulnerabilityDetails)}
	occurrence := containeranalysispb.Occurrence{NoteName: parent, ResourceUrl: imageUrl, Details: &vulDetails}
	req := &containeranalysispb.CreateOccurrenceRequest{Parent: project, Occurrence: &occurrence}
	return c.CreateOccurrence(ctx, req)
}

// [END create_occurrence]

// [START update_note]
//Makes an update to an existing note
func UpdateNote(updated *containeranalysispb.Note, noteId, projectId string) error {
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	noteName := containeranalysis.NotePath(projectId, noteId)
	req := &containeranalysispb.UpdateNoteRequest{Name: noteName, Note: updated}
	_, err = c.UpdateNote(ctx, req)
	return err
}

// [END update_note]

// [START update_occurrence]
//Makes an update to an existing occurrence
func UpdateOccurrence(updated *containeranalysispb.Occurrence, occurrenceName string) error {
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	req := &containeranalysispb.UpdateOccurrenceRequest{Name: occurrenceName, Occurrence: updated}
	_, err = c.UpdateOccurrence(ctx, req)
	return err
}

// [END update_occurrence]

// [START delete_note]
//Deletes an existing note from the project
func DeleteNote(noteId, projectId string) error {
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
func DeleteOccurrence(occurrenceName string) error {
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
func GetNote(noteId, projectId string) (*containeranalysispb.Note, error) {
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
func GetOccurrence(occurrenceName string) (*containeranalysispb.Occurrence, error) {
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	req := &containeranalysispb.GetOccurrenceRequest{occurrenceName}
	return c.GetOccurrence(ctx, req)
}

// [END get_occurrence]

// [START discovery_info]
//Prints the Discovery occurrence created for a specified image
//This occurrence contains information about the initial scan on the image
func GetDiscoveryInfo(imageUrl, projectId string) error {
	filterStr := `kind="DISCOVERY" AND resourceUrl="` + imageUrl + `"`
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return err
	}
	project := containeranalysis.ProjectPath(projectId)
	req := &containeranalysispb.ListOccurrencesRequest{Parent: project, Filter: filterStr}
	iterator := c.ListOccurrences(ctx, req)
	var complete error
	var occ *containeranalysispb.Occurrence
	for complete == nil {
		occ, complete = iterator.Next()
		fmt.Println(occ)
	}
	return nil
}

// [END discovery_info]

// [START occurrences_for_note]
//Retrieves all the occurrences associated with a specified note
func GetOccurrencesForNote(noteId, projectId string) (int, error) {
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return -1, err
	}
	noteName := containeranalysis.NotePath(projectId, noteId)
	req := &containeranalysispb.ListNoteOccurrencesRequest{Name: noteName}
	iterator := c.ListNoteOccurrences(ctx, req)
	var complete error
	count := -1
	for complete == nil {
		// we can do something with the retrieved occurence here
		// for this sample, we will just count them
		_, complete = iterator.Next()
		count = count + 1
	}

	return count, nil
}

// [END occurrences_for_note]

// [START occurrences_for_image]
//Retrieves all the occurrences associated with a specified image
func GetOccurrencesForImage(imageUrl, projectId string) (int, error) {
	filterStr := `resourceUrl="` + imageUrl + `"`
	ctx := context.Background()
	c, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return -1, err
	}
	project := containeranalysis.ProjectPath(projectId)
	req := &containeranalysispb.ListOccurrencesRequest{Parent: project, Filter: filterStr}
	iterator := c.ListOccurrences(ctx, req)
	var complete error
	count := -1
	for complete == nil {
		// we can do something with the retrieved occurence here
		// for this sample, we will just count them
		_, complete = iterator.Next()
		count = count + 1
	}

	return count, nil
}

// [END occurrences_for_image]

// [START pubsub]
//Handle incoming occurrences using a pubsub subscription
func Pubsub(subscriptionId string, timeout int, projectId string) (int, error) {
	ctx := context.Background()
	toctx, _ := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	var mu sync.Mutex
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return -1, err
	}
	sub := client.Subscription(subscriptionId)
	count := 0

	// listen on the subscription until the context times out
	err = sub.Receive(toctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		count = count + 1
		fmt.Printf("Message %d: %q\n", count, string(msg.Data))
		msg.Ack()
		mu.Unlock()
	})
	if err != nil {
		return -1, err
	}
	fmt.Println(count)
	return count, nil
}

//Creates and returns a pubsub subscription listening to the occurrence topic.
//This topic provides updates when occurrences are modified
func CreateOccurrenceSubscription(subscriptionId, projectId string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return err
	}

	topicId := "resource-notes-occurrences-v1alpha1"
	topic := client.Topic(topicId)
	config := pubsub.SubscriptionConfig{Topic: topic}
	_, err = client.CreateSubscription(ctx, subscriptionId, config)
	return err
}

// [END pubsub]
