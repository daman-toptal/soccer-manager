package router

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
	grpcRoot "protobuf-v1/golang"
)


type Response struct {
	Writer   http.ResponseWriter
	Data     interface{}
	Status   int
	GRPCData proto.Message
}

func RenderJSON(r Response) {
	var j []byte
	var err error

	if r.GRPCData != nil {
		mo := protojson.MarshalOptions{
			EmitUnpopulated: true,
		}

		j, err = mo.Marshal(r.GRPCData)
	} else {
		j, err = json.Marshal(r.Data)
	}

	r.Writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		r.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.Status > 0 {
		r.Writer.WriteHeader(r.Status)
	}

	r.Writer.Write(j)
}


func RenderHttpError(
	w http.ResponseWriter,
	r *http.Request,
	err *grpcRoot.HttpError) {

	RenderJSON(
		Response{
			Writer:   w,
			GRPCData: err,
		},
	)
}
