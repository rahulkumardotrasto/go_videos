package app

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"./providers"
	"./usecase"
	"google.golang.org/grpc"

	gzip "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//App has router and db instances
type App struct{}

var videoService = usecase.VideoService{}
var auth = providers.Auth{}

//Init initializes the app with predefined configuration
func (app *App) Init() {
	rand.Seed(time.Now().UnixNano())
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	routes.Use(gzip.Gzip(gzip.DefaultCompression))
	routes.Use(gin.ErrorLogger())
	routes.Use(gin.Recovery())

	routes.StaticFile("/favicon.ico", "")
	routes.Run(":8000")

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	videos := routes.Group("videos")
	{
		videos.POST("/videos", UploadVideo).Use(auth.Authenticate())
	}
}

//UploadVideo ...
func UploadVideo(c *gin.Context) {

	req := &pb.HelloRequest{Name: name}
	res, err := client.SayHello(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	statusCode, success, message := videoService.UploadVideo(*c)
	c.Header("Access-Control-Allow-Origin", "*")
	c.SecureJSON(statusCode, gin.H{
		"success": success,
		"message": message,
	})
}
