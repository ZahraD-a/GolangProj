package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/tasks/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "Tasks"
const colName = "task1"

//MOST IMPORTANT , because there is a collection in the db
var collection *mongo.Collection

// connect with monogoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
}

//create new task
func CreateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var task model.Tasks // call the struct which is in the file model.go
	_ = json.NewDecoder(r.Body).Decode(&task)

	insertOneTask(task)
	json.NewEncoder(w).Encode(task)

}
func insertOneTask(task model.Tasks) {
	inserted, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 tasks in db with id: ", inserted.InsertedID)
}

//------------------------affiche toutes les tasks----------------------------------
func GetMyAlltasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var tasks []primitive.M
	for cur.Next(context.Background()) {
		var task bson.M
		err := cur.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	defer cur.Close(context.Background())

	json.NewEncoder(w).Encode(tasks)
}

//<-------------delete all tasks from mongodb------------------------>
func DeleteAlltasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAlltasks() //call the function delteAlltasks
	json.NewEncoder(w).Encode(count)
}

//definition de la fonction delete alltask
func deleteAlltasks() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NUmber of tasks delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//------------------delete One task by giving Id-----------------------
func DeleteAtask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func deleteTask(taskId string) {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("task got delete with delete count: ", deleteCount)
}

// ------------update 1 task------------------
func UpdateTask(taskId string) {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	var task model.Tasks
	update := bson.M{
		"$set": bson.M{
			"Titre":      task.Titre,
			"DateDebut":  task.DateDebut,
			"estimation": task.Estimation,
			"status":     task.Status,
		}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified  ", result)
}
func UpdateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	var task model.Tasks
	json.NewDecoder(r.Body).Decode(&task)

	UpdateTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
