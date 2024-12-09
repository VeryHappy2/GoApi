package profile

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/VeryHappy2/GoApi/internal/http-server/response"
	"github.com/VeryHappy2/GoApi/internal/repositories"
	"github.com/VeryHappy2/GoApi/internal/storage/sqlite"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type RespProfile struct {
	Status string  `json:"status" example:"success"`
	Data   Profile `json:"data"`
	Error  string  `json:"error,omitempty" example:""`
}

// @Summary		GetById.
// @Description	Getting by an profile.
// @Tags			profile
// @Accept			json
// @Produce		json
// @Param        id     path      int     true  "ID of the profile"
// @Success		200		{object}	RespProfile	"api response"
// @Failure      404   {object}  response.DefaultResponse "profile not found"
// @Failure      400   {object}  response.DefaultResponse "invalid request"
// @Router			/profile/getById/{id} [post]
func (p Profile) GetById(log *slog.Logger, repository repositories.ProfileRep, db *sqlite.Storage) http.HandlerFunc {
	return func(respWriter http.ResponseWriter, request *http.Request) {
		const op = "handlers.profile.getById"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(request.Context())),
		)

		idStr := chi.URLParam(request, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			log.Error("failed to decode request body")
			respWriter.WriteHeader(http.StatusBadRequest)
			render.JSON(respWriter, request, response.Error[int8]("failed to decode request", nil))

			return
		}

		repository.Object.Id = id

		profile, err := repository.GetById(db)
		if err != nil {
			log.Error("failed to select profile")
			respWriter.WriteHeader(http.StatusNotFound)
			render.JSON(respWriter, request, response.Error[int8]("failed to select profile", nil))

			return
		}

		log.Info("Profile added")

		render.JSON(respWriter, request, response.OK(profile))
	}
}
