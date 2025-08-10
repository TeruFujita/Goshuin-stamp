package handlers

import (
	"net/http"

	"stamp-backend/internal/ent"
)

// GetGuide 御朱印のガイド情報を取得します
func GetGuide(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 御朱印のガイド情報を返す
		guide := map[string]interface{}{
			"title":       "Goshuin Guide",
			"description": "Learn about Japanese temple stamps and how to collect them",
			"sections": []map[string]interface{}{
				{
					"title":     "What is Goshuin?",
					"content":   "Goshuin (御朱印) are special stamps or calligraphy that you can receive at Japanese temples and shrines. They serve as proof of your visit and are considered sacred items.",
					"image_url": "/images/goshuin-example.jpg",
				},
				{
					"title":     "How to Receive Goshuin",
					"content":   "1. Visit the temple or shrine during opening hours\n2. Look for the goshuin office (御朱印所)\n3. Pay the fee (usually 300-500 yen)\n4. Present your goshuin book or paper\n5. Wait while the priest writes the goshuin",
					"image_url": "/images/goshuin-process.jpg",
				},
				{
					"title":     "Etiquette and Manners",
					"content":   "- Dress modestly and respectfully\n- Be quiet and respectful in sacred areas\n- Don't take photos of the goshuin writing process\n- Handle your goshuin book with care\n- Don't rush the priest while they're writing",
					"image_url": "/images/temple-manners.jpg",
				},
				{
					"title":     "Goshuin Book (御朱印帳)",
					"content":   "A goshuin book is a special notebook designed to collect goshuin. You can purchase one at most temples and shrines, or bring your own. Traditional books are made of washi paper and have beautiful covers.",
					"image_url": "/images/goshuin-book.jpg",
				},
				{
					"title":     "Best Practices",
					"content":   "- Start with famous temples in your area\n- Visit during weekdays to avoid crowds\n- Check temple websites for special goshuin\n- Keep your goshuin book in a protective case\n- Document your visits with photos",
					"image_url": "/images/temple-visit.jpg",
				},
			},
			"tips": []string{
				"Some temples offer special goshuin for different seasons",
				"Many temples have multiple goshuin designs",
				"Some temples require advance reservations for goshuin",
				"Goshuin are considered sacred items, so treat them with respect",
				"You can collect goshuin at both temples (寺) and shrines (神社)",
			},
		}

		writeJSON(w, http.StatusOK, guide)
	}
}
