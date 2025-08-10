package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"stamp-backend/internal/ent"
)

// GetTemples 寺社一覧を取得します
func GetTemples(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temples, err := client.Temple.Query().All(r.Context())
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to fetch temples")
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"temples": temples,
		})
	}
}

// GetTemple 特定の寺社を取得します
func GetTemple(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// URLパスからIDを抽出
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 5 {
			writeError(w, http.StatusBadRequest, "Invalid temple ID")
			return
		}
		
		id, err := strconv.Atoi(pathParts[4])
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid temple ID")
			return
		}

		temple, err := client.Temple.Get(r.Context(), id)
		if err != nil {
			writeError(w, http.StatusNotFound, "Temple not found")
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"temple": temple,
		})
	}
}

// GetNearbyTemples 現在地から近い寺社を取得します
func GetNearbyTemples(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lat := r.URL.Query().Get("lat")
		lng := r.URL.Query().Get("lng")
		radius := r.URL.Query().Get("radius")
		if radius == "" {
			radius = "10" // デフォルト10km
		}

		if lat == "" || lng == "" {
			writeError(w, http.StatusBadRequest, "Latitude and longitude are required")
			return
		}

		// TODO: 実際の位置情報検索ロジックを実装
		// 現在はダミーデータを返す
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"temples": []map[string]interface{}{
				{
					"id":         1,
					"name":       "Senso-ji Temple",
					"name_en":    "Senso-ji Temple",
					"latitude":   35.7148,
					"longitude":  139.7967,
					"distance":   0.5,
				},
				{
					"id":         2,
					"name":       "Meiji Shrine",
					"name_en":    "Meiji Shrine",
					"latitude":   35.6764,
					"longitude":  139.6993,
					"distance":   2.1,
				},
			},
		})
	}
}

// writeJSON JSONレスポンスを書き込みます
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError エラーレスポンスを書き込みます
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]interface{}{
		"error": message,
	})
}
