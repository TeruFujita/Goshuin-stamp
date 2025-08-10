package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"stamp-backend/internal/ent"
	"stamp-backend/internal/handlers"
)

// Server HTTPサーバー構造体
type Server struct {
	client *ent.Client
	mux    *http.ServeMux
}

// New 新しいサーバーインスタンスを作成します
func New(client *ent.Client) *Server {
	s := &Server{
		client: client,
		mux:    http.NewServeMux(),
	}
	s.setupRoutes()
	return s
}

// setupRoutes ルートを設定します
func (s *Server) setupRoutes() {
	// ヘルスチェックエンドポイント
	s.mux.HandleFunc("GET /health", s.handleHealth)

	// APIルート
	s.mux.HandleFunc("GET /api/v1/temples", s.handleGetTemples)
	s.mux.HandleFunc("GET /api/v1/temples/{id}", s.handleGetTemple)
	s.mux.HandleFunc("GET /api/v1/temples/nearby", s.handleGetNearbyTemples)
	
	s.mux.HandleFunc("GET /api/v1/goshuin", s.handleGetGoshuinCollections)
	s.mux.HandleFunc("POST /api/v1/goshuin", s.handleCreateGoshuinCollection)
	s.mux.HandleFunc("GET /api/v1/goshuin/{id}", s.handleGetGoshuinCollection)
	s.mux.HandleFunc("PUT /api/v1/goshuin/{id}", s.handleUpdateGoshuinCollection)
	s.mux.HandleFunc("DELETE /api/v1/goshuin/{id}", s.handleDeleteGoshuinCollection)
	
	s.mux.HandleFunc("GET /api/v1/guide", s.handleGetGuide)

	// CORS対応のミドルウェアを追加
	s.mux.HandleFunc("OPTIONS /", s.handleCORS)
}

// Run サーバーを起動します
func (s *Server) Run(addr string) error {
	log.Printf("Server starting on %s", addr)
	return http.ListenAndServe(addr, s.corsMiddleware(s.mux))
}

// corsMiddleware CORSミドルウェア
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// handleHealth ヘルスチェックハンドラー
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	s.writeJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"message": "Goshuin App API is running",
	})
}

// handleCORS CORSハンドラー
func (s *Server) handleCORS(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// writeJSON JSONレスポンスを書き込みます
func (s *Server) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError エラーレスポンスを書き込みます
func (s *Server) writeError(w http.ResponseWriter, status int, message string) {
	s.writeJSON(w, status, map[string]interface{}{
		"error": message,
	})
}

// 寺社関連のハンドラー
func (s *Server) handleGetTemples(w http.ResponseWriter, r *http.Request) {
	handlers.GetTemples(s.client)(w, r)
}

func (s *Server) handleGetTemple(w http.ResponseWriter, r *http.Request) {
	handlers.GetTemple(s.client)(w, r)
}

func (s *Server) handleGetNearbyTemples(w http.ResponseWriter, r *http.Request) {
	handlers.GetNearbyTemples(s.client)(w, r)
}

// 御朱印関連のハンドラー
func (s *Server) handleGetGoshuinCollections(w http.ResponseWriter, r *http.Request) {
	handlers.GetGoshuinCollections(s.client)(w, r)
}

func (s *Server) handleCreateGoshuinCollection(w http.ResponseWriter, r *http.Request) {
	handlers.CreateGoshuinCollection(s.client)(w, r)
}

func (s *Server) handleGetGoshuinCollection(w http.ResponseWriter, r *http.Request) {
	handlers.GetGoshuinCollection(s.client)(w, r)
}

func (s *Server) handleUpdateGoshuinCollection(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateGoshuinCollection(s.client)(w, r)
}

func (s *Server) handleDeleteGoshuinCollection(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteGoshuinCollection(s.client)(w, r)
}

// ガイド関連のハンドラー
func (s *Server) handleGetGuide(w http.ResponseWriter, r *http.Request) {
	handlers.GetGuide(s.client)(w, r)
}
