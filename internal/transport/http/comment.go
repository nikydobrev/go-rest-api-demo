package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nikydobrev/go-rest-api-demo/internal/comment"
	log "github.com/sirupsen/logrus"
)

// GetComment - retrive comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}
	comment, err := h.Service.GetComment(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by ID", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Error(err)
		return
	}
}

// GetAllComments - retrive all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")

	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error retrieving all comments", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		log.Error(err)
		return
	}
}

// PostComment - post a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decodde JSON body", err)
		return
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Error posting a new comment", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Error(err)
		return
	}
}

// UpdateComment - update a new comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var updatedComment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&updatedComment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	updatedComment, err = h.Service.UpdateComment(uint(commentID), updatedComment)
	if err != nil {
		sendErrorResponse(w, "Error updating a comment", err)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedComment); err != nil {
		log.Error(err)
		return
	}
}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Error deleting a comment", err)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		log.Error(err)
		return
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		log.Error(err)
	}
}
