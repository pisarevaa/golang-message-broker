package api

import (
	"net/http"
)

//	@Summary	Send message to queue subscribers
//	@Tags		Messages
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Failure	409	{object}	ErrorResponse
//	@Router		/v1/queues/{queue_name}/messages [post].
func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	// err := h.Service.UpdateDatabase(r.Context())
	// if err != nil {
	// 	log.Error().Err(err).Msg("database sync failed")
	// 	writeError(w, "Failed to start sync", http.StatusInternalServerError)
	// 	return
	// }
	writeJSON(w, map[string]string{"status": "database sync completed successfully"})
}

//	@Summary	Get message from queue
//	@Tags		Messages
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Failure	409	{object}	ErrorResponse
//	@Router		/v1/queues/{queue_name}/messages [get].
func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	// err := h.Service.UpdateDatabase(r.Context())
	// if err != nil {
	// 	log.Error().Err(err).Msg("database sync failed")
	// 	writeError(w, "Failed to start sync", http.StatusInternalServerError)
	// 	return
	// }
	writeJSON(w, map[string]string{"status": "database sync completed successfully"})
}

//	@Summary	Subscribe to a queue
//	@Tags		Subscriptions
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Failure	409	{object}	ErrorResponse
//	@Router		/v1/queues/{queue_name}/subscriptions [post].
func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	// err := h.Service.UpdateDatabase(r.Context())
	// if err != nil {
	// 	log.Error().Err(err).Msg("database sync failed")
	// 	writeError(w, "Failed to start sync", http.StatusInternalServerError)
	// 	return
	// }
	writeJSON(w, map[string]string{"status": "database sync completed successfully"})
}
