package v1

import (
	"context"
	"fmt"

	v1 "go-grpc-rest/pkg/api"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}

// implementation of v1.ToDoServiceServer proto interface
type toDoServiceServer struct {
	db *mongo.Database
}

// create ToDo service
func NewToDoServiceServer(db *mongo.Database) v1.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

// checks if the API versino requested by client is supported by server
func (s *toDoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *toDoServiceServer) GetTodoCollection(ctx context.Context) *mongo.Collection {
	return s.db.Collection("todo")
}

// Create new todo task
func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.GetTodoCollection(ctx)

	todo := req.GetToDo()
	data := Todo{
		Title:       todo.GetTitle(),
		Description: todo.GetDescription(),
	}

	res, err := c.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
	}

	oid := res.InsertedID.(primitive.ObjectID)

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  oid.Hex(),
	}, nil
}

// Read todo task
func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.GetTodoCollection(ctx)

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	result := c.FindOne(ctx, bson.M{"_id": oid})

	data := Todo{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object Id %s: %v", req.GetId(), err))
	}

	return &v1.ReadResponse{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Id:          oid.Hex(),
			Title:       data.Title,
			Description: data.Description,
		},
	}, nil
}

// Update todo task
func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.GetTodoCollection(ctx)

	todo := req.GetToDo()

	oid, err := primitive.ObjectIDFromHex(todo.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	update := bson.M{
		"title":       todo.GetTitle(),
		"description": todo.GetDescription(),
	}

	filter := bson.M{"_id": oid}

	// update ToDo
	res := c.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	decoded := Todo{}
	err = res.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find blog with supplied ID: %v", err),
		)
	}

	return &v1.UpdateResponse{
		Api: apiVersion,
		Updated: &v1.ToDo{
			Id:          decoded.ID.Hex(),
			Title:       decoded.Title,
			Description: decoded.Description,
		},
	}, nil
}

// Delete todo task
func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.GetTodoCollection(ctx)

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	_, err = c.DeleteOne(ctx, bson.M{"_id": oid})

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: "true",
	}, nil
}

// Read all todo tasks
func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.GetTodoCollection(ctx)

	data := &Todo{}

	cursor, err := c.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}

	defer cursor.Close(ctx)

	list := []*v1.ToDo{}
	for cursor.Next(ctx) {
		err := cursor.Decode(data)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		list = append(list, &v1.ToDo{
			Id:          data.ID.Hex(),
			Title:       data.Title,
			Description: data.Description,
		})
	}

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		ToDos: list,
	}, nil
}
