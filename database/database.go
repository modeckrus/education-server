package database

import (
	"context"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"

	cstorage "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	fstorage "firebase.google.com/go/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

//DB ...
type DB struct {
	Client           *mongo.Client
	Firebase         *firebase.App
	Storage          *fstorage.Client
	Bucket           *cstorage.BucketHandle
	signedURLOptions cstorage.SignedURLOptions
}

//Connect to the mongodb and firebase
func Connect(addr string) *DB {
	var err error
	var client *mongo.Client

	client, err = mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	print("Disconnecting...")
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	fmt.Println("Firebase Initialize")

	opt := option.WithCredentialsFile("/home/modeck/go/src/education/serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing app: %v", err))
	}
	storage, err := app.Storage(ctx)
	if err != nil {
		log.Fatal(err)
	}
	bucket, err := storage.Bucket("education-modeck.appspot.com")
	if err != nil {
		log.Println("Error while initialize bucket")
		log.Fatal(err)
	}

	keypem, err := os.Open("/home/modeck/go/src/education/key.pem")
	if err != nil {
		log.Println("Error while opening key.pem")
		log.Fatal(err)
	}
	keypembytes, err := ioutil.ReadAll(keypem)
	if err != nil {
		log.Fatal(err)
	}
	signedURLOptions := cstorage.SignedURLOptions{
		GoogleAccessID: "firebase-adminsdk-jxiir@education-modeck.iam.gserviceaccount.com",
		PrivateKey:     keypembytes,
		// SignBytes: func(b []byte) ([]byte, error) {
		// 	_, signedBytes, err := appengine.SignBytes(ctx, b)
		// 	return signedBytes, err
		// },
		Expires: time.Now().AddDate(0, 0, 1),
		Method:  "GET",
	}
	// client
	return &DB{
		Client:           client,
		Firebase:         app,
		Storage:          storage,
		Bucket:           bucket,
		signedURLOptions: signedURLOptions,
	}
}
