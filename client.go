package main

import (
	"fmt"
	"grpc-crud/server/pb/messages"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var client messages.UserServiceClient

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = messages.NewUserServiceClient(conn)

	g := gin.Default()
	g.POST("/add/user", createUser)
	g.GET("/find/user/:uid", getUser)
	g.PUT("/update/user/:uid", updateUser)
	g.DELETE("/delete/user/:uid", deleteUser)

	log.Fatal(g.Run(":8080"))
}

func createUser(ctx *gin.Context) {
	u := messages.User{}
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := client.Add(ctx, &u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Location", fmt.Sprintf("%s/find/user/%d",
		ctx.Request.Host, res.GetObjectId().GetUid()))
	ctx.JSON(http.StatusCreated, res)
}

func getUser(ctx *gin.Context) {
	uid, err := strconv.ParseUint(ctx.Param("uid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	obj := messages.ObjectId{Uid: uid}

	res, err := client.Find(ctx, &obj)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func updateUser(ctx *gin.Context) {

	u := messages.User{}
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := strconv.ParseUint(ctx.Param("uid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	obj := messages.ObjectId{Uid: uid}
	u.ObjectId = &obj

	res, err := client.Update(ctx, &u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func deleteUser(ctx *gin.Context) {

	uid, err := strconv.ParseUint(ctx.Param("uid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	obj := messages.ObjectId{Uid: uid}

	res, err := client.Delete(ctx, &obj)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Entity", fmt.Sprintf("%v", res.GetUid()))
	ctx.JSON(http.StatusNoContent, nil)
}
