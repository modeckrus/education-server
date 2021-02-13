package database

import (
	"context"
	"education/model"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const jwtsecret = "secret"

//FirebaseAuth ...
func (db *DB) FirebaseAuth(ctx context.Context, firetoken string) (*model.UserToken, error) {
	client, err := db.Firebase.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, firetoken)
	if err != nil {
		return nil, err
	}
	fireUser, err := client.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}
	coll := db.Client.Database("users").Collection("users")
	res := coll.FindOne(ctx, bson.M{"uid": fireUser.UID})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			var photo *model.Image
			if fireUser.DisplayName == "" {
				fireUser.DisplayName = strings.Split(fireUser.Email, "@")[0]
			}
			if fireUser.PhotoURL == "" {
				photo = defaultPhoto()
			}
			user := &model.User{
				DisplayName: fireUser.DisplayName,
				Name:        fireUser.DisplayName,
				Email:       fireUser.Email,
				UID:         fireUser.UID,
				Photo:       photo,
				ProviderID:  *(&fireUser.ProviderID),
			}
			log.Println("New User: ", user)
			isres, err := coll.InsertOne(ctx, user)
			if err != nil {
				return nil, err
			}
			user.ID = isres.InsertedID.(primitive.ObjectID)
			tokens, err := GenerateTokenPair(user.ID.Hex())
			if err != nil {
				return nil, err
			}
			return &model.UserToken{
				User:   *user,
				Tokens: *tokens,
			}, nil
		}
	}
	user := &model.User{
		EmailVerified: false,
		ProviderID:    "dbpassword",
	}
	err = res.Decode(user)
	if err != nil {
		return nil, err
	}
	tokens, err := GenerateTokenPair(user.ID.Hex())
	if err != nil {
		return nil, err
	}

	return &model.UserToken{
		User:   *user,
		Tokens: *tokens,
	}, nil
}

type UpdateUserInput struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	PhotoID     *string `json:"photoId"`
	Email       string  `json:"email"`
}

//UpdateUser ...
func (db *DB) UpdateUser(ctx context.Context, currUser model.CheckedUser, name string, displayName string, iphotoID *string) (*model.User, error) {
	coll := db.Client.Database("users").Collection("users")
	id, err := primitive.ObjectIDFromHex(currUser.ID)
	if err != nil {
		return nil, err
	}
	res := coll.FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return nil, err
	}
	user := &model.User{}
	err = res.Decode(user)
	if err != nil {
		return nil, err
	}

	var avatar *model.Image
	if iphotoID == nil {
		avatar = defaultPhoto()
	} else {
		pcoll := db.Client.Database("files").Collection("images")
		photoID, err := primitive.ObjectIDFromHex(*iphotoID)
		if err != nil {
			return nil, err
		}
		res := pcoll.FindOne(ctx, bson.M{"_id": photoID})
		if res.Err() != nil {
			return nil, res.Err()
		}
		err = res.Decode(avatar)
		if err != nil {
			return nil, err
		}
	}
	user.DisplayName = displayName
	user.Name = name
	user.Photo = avatar
	repres := coll.FindOneAndReplace(ctx, bson.M{"_id": id}, user)
	if repres.Err() != nil {
		return nil, repres.Err()
	}
	return user, nil
}

//UpdateUser ...
// func (db *DB) UpdateUser(ctx context.Context, currUser model.UserToken, user model.UserInput) (*model.User, error) {
// 	coll := db.Client.Database("users").Collection("users")
// 	id, err := primitive.ObjectIDFromHex(currUser.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res := coll.FindOne(ctx, bson.M{"_id": id})
// 	if res.Err() != nil {
// 		log.Println("Error while finding user: ")
// 		log.Println(res.Err().Error())
// 		return nil, res.Err()
// 	}
// 	nuser := model.MUser{}
// 	err = res.Decode(&nuser)
// 	if err != nil {
// 		log.Println("Error while decoding: ")
// 		log.Println(err)
// 		return nil, err
// 	}
// 	log.Println("User before: ")
// 	nuserjs, err := json.Marshal(nuser)
// 	if err != nil {
// 		log.Println("Error while marshaling: ")
// 		log.Println(err)
// 		return nil, err
// 	}
// 	log.Println(string(nuserjs))
// 	var name, displayName, email string
// 	var photo model.MImage
// 	if user.Name == "" {
// 		name = nuser.Name
// 	} else {
// 		name = user.Name
// 	}
// 	if user.DisplayName == "" {
// 		displayName = nuser.DisplayName
// 	} else {
// 		displayName = user.DisplayName
// 	}
// 	if user.Email == "" {
// 		email = nuser.Email
// 		nuser.EmailVerified = false
// 	} else {
// 		email = user.Email
// 	}
// 	if user.PhotoID != nil {
// 		pcoll := db.Client.Database("files").Collection("images")
// 		photoID, err := primitive.ObjectIDFromHex(*user.PhotoID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		res := pcoll.FindOne(ctx, bson.M{"_id": photoID})
// 		if res.Err() != nil {
// 			return nil, res.Err()
// 		}
// 		var imgFile model.MImage
// 		err = res.Decode(&imgFile)
// 		if err != nil {
// 			return nil, err
// 		}
// 		photo = imgFile
// 	} else {
// 		log.Println("No photo")
// 	}
// 	// if user.PhotoURL == "" {
// 	// 	photoURL = nuser.PhotoURL
// 	// 	blur = nuser.PhotoBlur
// 	// } else {
// 	// 	photoURL = user.PhotoURL
// 	// 	blur, err = createBlur(photoURL)
// 	// 	if err != nil {
// 	// 		log.Println("Error while create blur")
// 	// 		log.Println(err)
// 	// 		return nil, err
// 	// 	}
// 	// }
// 	nuser = model.MUser{
// 		ID:          id,
// 		UID:         nuser.UID,
// 		ProviderID:  nuser.ProviderID,
// 		Password:    nuser.Password,
// 		Name:        name,
// 		DisplayName: displayName,
// 		Email:       email,
// 		Photo:       &photo,
// 	}
// 	log.Println("User after:")
// 	nuserjs, err = json.Marshal(nuser)
// 	if err != nil {
// 		log.Println("Error while marshaling: ")
// 		log.Println(err)
// 		return nil, err
// 	}
// 	log.Println(string(nuserjs))
// 	coll.DeleteOne(ctx, bson.M{"_id": id})
// 	_, err = coll.InsertOne(ctx, &nuser)
// 	if err != nil {
// 		log.Println("Error while update user")
// 		log.Println(err)
// 		return nil, err
// 	}
// 	res = coll.FindOne(ctx, bson.M{"_id": currUser.ID})
// 	if res.Err() != nil {
// 		log.Println("Error while finding in the end to return: ")
// 		log.Println(res.Err())
// 		return nil, res.Err()
// 	}
// 	err = res.Decode(&nuser)
// 	if err != nil {
// 		log.Println("Error while decoding in the end")
// 		log.Println(err)
// 		return nil, err
// 	}
// 	rphoto := &model.Image{
// 		ID:       photo.ID.Hex(),
// 		UID:      photo.UID,
// 		Name:     photo.Name,
// 		URL:      photo.URL,
// 		Crops:    photo.Crops,
// 		BlurHash: photo.BlurHash,
// 	}
// 	ruser := model.User{
// 		ID:          nuser.ID.Hex(),
// 		UID:         nuser.UID,
// 		ProviderID:  nuser.ProviderID,
// 		Password:    nuser.Password,
// 		Name:        name,
// 		DisplayName: displayName,
// 		Email:       email,
// 		Photo:       rphoto,
// 	}
// 	return &ruser, nil

// }

//GetUserByToken ...
func (db *DB) GetUserByToken(jwtstr string) (model.CheckedUser, error) {
	token, err := jwt.Parse(jwtstr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtsecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		user := model.CheckedUser{}

		user.ID = (claims["uid"].(string))
		if claims["roles"] != nil {
			rolestr := (claims["roles"]).(string)
			user.Roles = strings.Split(rolestr, "|")
		}
		return user, nil
	}

	return model.CheckedUser{}, err
}

//GenerateTokenPair jwt token
func GenerateTokenPair(id string) (*model.Tokens, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = id
	claims["roles"] = "user"

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(jwtsecret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte(jwtsecret))
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		Token:        t,
		RefreshToken: rt,
	}, nil
}
