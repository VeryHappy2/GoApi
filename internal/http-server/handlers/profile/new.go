package profile

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/VeryHappy2/GoApi/internal/http-server/response"
	"github.com/VeryHappy2/GoApi/internal/repositories"
	"github.com/VeryHappy2/GoApi/internal/storage/sqlite"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// @Summary	New
// @Tags		profile
// @Accept		json
// @Produce	json
//
// @Param		input	body		repositories.Profile	true	"profile"
//
// @Success	200		{object}	response.DefaultResponse
// @Failure	400		{object}	response.DefaultResponse
// @Router		/profile/add [post]
func (Profile) New(log *slog.Logger, repository repositories.ProfileRep, db *sqlite.Storage) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, request *http.Request) {
		const op = "handlers.profile.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(request.Context())),
		)
		err := render.DecodeJSON(request.Body, &repository.Object)

		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			respWriter.WriteHeader(http.StatusBadRequest)
			render.JSON(respWriter, request, response.Error[int8]("empty request", nil))

			return
		}
		if err != nil {
			log.Error("failed to decode request body")
			respWriter.WriteHeader(http.StatusBadRequest)
			render.JSON(respWriter, request, response.Error[int8]("failed to decode request", nil))

			return
		}

		err = repository.Add(db)
		if err != nil {
			log.Error("failed to add profile", "err", err)
			respWriter.WriteHeader(http.StatusBadRequest)
			render.JSON(respWriter, request, response.Error[int8]("failed to add profile", nil))

			return
		}

		log.Info("Profile added")

		successFul := "success"
		render.JSON(respWriter, request, response.OK[string](&successFul))
	}
}
