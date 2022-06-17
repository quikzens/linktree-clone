package link

import (
	"context"
	"errors"
	"linktree-clone/db"
	"linktree-clone/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"gopkg.in/mgo.v2/bson"
)

func CreateLink(c *gin.Context) {
	payload, _ := c.Get("user")
	userPayload, _ := payload.(*util.UserPayload)
	userEmail := userPayload.Email

	newLinkID := uuid.NewString()
	newLink := link{
		ID:         newLinkID,
		EmailOwner: userEmail,
		Title:      "Title",
		Url:        "Url",
		IsActive:   true,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  0,
	}
	_, err := db.LinkColl.InsertOne(context.TODO(), newLink)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	var user userLinks
	err = db.UserColl.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	updatedLinks := []string{newLinkID}
	updatedLinks = append(updatedLinks, user.Links...)
	_, err = db.UserColl.UpdateOne(context.TODO(), bson.M{"email": userEmail}, bson.M{
		"$set": updateUserLinks{
			Links:     updatedLinks,
			UpdatedAt: time.Now().Unix(),
		},
	})
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	util.SendSuccess(c, newLink)
}

func UpdateLink(c *gin.Context) {
	payload, _ := c.Get("user")
	userPayload, _ := payload.(*util.UserPayload)
	userEmail := userPayload.Email
	linkID := c.Param("id")

	var req updateLinkRequest
	err := c.Bind(&req)
	if err != nil {
		util.SendBadRequest(c, err)
		return
	}

	var link link
	err = db.LinkColl.FindOne(context.TODO(), bson.M{"id": linkID}).Decode(&link)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	if link.EmailOwner != userEmail {
		util.SendBadRequest(c, errors.New("the link is not belong to this user"))
		return
	}

	req.UpdatedAt = time.Now().Unix()
	_, err = db.LinkColl.UpdateOne(context.TODO(), bson.M{"id": linkID}, bson.M{"$set": req})
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	util.SendSuccess(c, nil)
}

func DeleteLink(c *gin.Context) {
	payload, _ := c.Get("user")
	userPayload, _ := payload.(*util.UserPayload)
	userEmail := userPayload.Email
	linkID := c.Param("id")

	var link link
	err := db.LinkColl.FindOne(context.TODO(), bson.M{"id": linkID}).Decode(&link)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	if link.EmailOwner != userEmail {
		util.SendBadRequest(c, errors.New("the link is not belong to this user"))
		return
	}

	_, err = db.LinkColl.DeleteOne(context.TODO(), bson.M{"id": linkID})
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	var user userLinks
	err = db.UserColl.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	updatedLinks := user.Links
	deletedIndex := slices.Index(updatedLinks, linkID)
	updatedLinks = slices.Delete(updatedLinks, deletedIndex, deletedIndex+1)
	_, err = db.UserColl.UpdateOne(context.TODO(), bson.M{"email": userEmail}, bson.M{
		"$set": updateUserLinks{
			Links:     updatedLinks,
			UpdatedAt: time.Now().Unix(),
		},
	})
	if err != nil {
		util.SendServerError(c, err)
		return
	}

	util.SendSuccess(c, nil)
}
