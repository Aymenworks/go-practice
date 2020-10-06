package repository

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Aymenworks/go-practice/entity"
)

type repo struct{}

func NewFirestorePostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client object %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to add post in firestore db %v", err)
		return nil, err
	}

	return post, nil
}

const (
	projectID      = "go-practice-posts"
	collectionName = "posts"
)

func (*repo) GetAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client object %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Printf("Failed to iterate next the list of posts", err)
			break
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	if err != nil {
		log.Fatalf("Failed to add post in firestore db %v", err)
		return nil, err
	}

	return posts, nil
}

func (*repo) FindByID(id string) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client object %v", err)
		return nil, err
	}

	fmt.Println("query", id)

	defer client.Close()

	doc, err := client.Doc(collectionName + "/" + id).Get(ctx)
	if err != nil {
		fmt.Println("error", err)
		// TODO: Handle error.
	} else {
		fmt.Println("yayyyy", doc.Data())
	}

	post := entity.Post{
		ID:    doc.Data()["ID"].(int64),
		Title: doc.Data()["Title"].(string),
		Text:  doc.Data()["Text"].(string),
	}

	fmt.Println("post", post)

	return &post, nil
}
