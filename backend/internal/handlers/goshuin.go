package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"stamp-backend/internal/ent"
)

// GetGoshuinCollections 御朱印コレクション一覧を取得します
func GetGoshuinCollections(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collections, err := client.GoshuinCollection.Query().All(r.Context())
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to fetch goshuin collections")
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"collections": collections,
		})
	}
}

// CreateGoshuinCollection 新しい御朱印コレクションを作成します
func CreateGoshuinCollection(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			TempleID int    `json:"temple_id"`
			ImageURL string `json:"image_url"`
			Notes    string `json:"notes"`
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Failed to read request body")
			return
		}

		if err := json.Unmarshal(body, &req); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON format")
			return
		}

		if req.TempleID == 0 {
			writeError(w, http.StatusBadRequest, "Temple ID is required")
			return
		}

		collection, err := client.GoshuinCollection.Create().
			SetTempleID(req.TempleID).
			SetImageURL(req.ImageURL).
			SetNotes(req.Notes).
			Save(r.Context())

		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to create goshuin collection")
			return
		}

		writeJSON(w, http.StatusCreated, map[string]interface{}{
			"collection": collection,
		})
	}
}

// GetGoshuinCollection 特定の御朱印コレクションを取得します
func GetGoshuinCollection(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 5 {
			writeError(w, http.StatusBadRequest, "Invalid collection ID")
			return
		}

		id, err := strconv.Atoi(pathParts[4])
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid collection ID")
			return
		}

		collection, err := client.GoshuinCollection.Get(r.Context(), id)
		if err != nil {
			writeError(w, http.StatusNotFound, "Goshuin collection not found")
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"collection": collection,
		})
	}
}

// UpdateGoshuinCollection 御朱印コレクションを更新します
func UpdateGoshuinCollection(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 5 {
			writeError(w, http.StatusBadRequest, "Invalid collection ID")
			return
		}

		id, err := strconv.Atoi(pathParts[4])
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid collection ID")
			return
		}

		var req struct {
			ImageURL string `json:"image_url"`
			Notes    string `json:"notes"`
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Failed to read request body")
			return
		}

		if err := json.Unmarshal(body, &req); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON format")
			return
		}

		collection, err := client.GoshuinCollection.UpdateOneID(id).
			SetImageURL(req.ImageURL).
			SetNotes(req.Notes).
			Save(r.Context())

		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to update goshuin collection")
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"collection": collection,
		})
	}
}

// DeleteGoshuinCollection 御朱印コレクションを削除します
func DeleteGoshuinCollection(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 5 {
			writeError(w, http.StatusBadRequest, "Invalid collection ID")
			return
		}

		id, err := strconv.Atoi(pathParts[4])
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid collection ID")
			return
		}

		if err := client.GoshuinCollection.DeleteOneID(id).Exec(r.Context()); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to delete goshuin collection")
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"message": "Goshuin collection deleted successfully",
		})
	}
}
