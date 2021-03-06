package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/revanthreddy7/mongo_golang/models"


)

type UserController struct{
	session *mgo.Session
}

type PostController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func NewPostController(s *mgo.Session) *PostController{
	return &PostController{s}
}

func (uc UserController) GetUser (w http.ResponseWriter,r *http.Request, p httprouter.Params){
	id:= p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid :=bson.ObjectIdHex(id)

	u := models.User{}
	
	if err := uc.session.DB("mongo_golang").C("users").FindId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return 
	}

	uj, err:=json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	u:=models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)




}


func (uc UserController) DeleteUser (w http.ResponseWriter,r *http.Request, p httprouter.Params){
	id :=p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo_golang").C("users").RemoveId(oid); err != nil{
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Deleted user",oid,"\n")
}


func (uc PostController) CreatePost(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
	u:=models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("posts").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}


	fmt.Println("Give the link id")


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)


}
func (uc PostController) GetPost (w http.ResponseWriter,r *http.Request, p httprouter.Params){
	id:= p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid :=bson.ObjectIdHex(id)

	u := models.User{}
	
	if err := uc.session.DB("mongo_golang").C("posts").FindId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return 
	}

	uj, err:=json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}


func (uc PostController) DeletePost (w http.ResponseWriter,r *http.Request, p httprouter.Params){
	id :=p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo_golang").C("posts").RemoveId(oid); err != nil{
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Deleted Post",oid,"\n")
}
