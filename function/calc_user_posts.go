package function

import (
	"context"

	"cloud.google.com/go/firestore"
)

// ユーザがPostを投稿した際に、ユーザの投稿について再計算を行う
func CalcUserPosts(store *firestore.Client, uid string) error {
	ctx := context.Background()
	_, _, err := store.Collection("test").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		return err
	}
	return nil
}
