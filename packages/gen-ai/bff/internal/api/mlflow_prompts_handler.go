package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	helper "github.com/opendatahub-io/gen-ai/internal/helpers"
	mlflowsdk "github.com/opendatahub-io/mlflow-go/mlflow"
)

// MLFlowListPromptsHandler handles GET /api/v1/mlflow/prompts
func (app *App) MLFlowListPromptsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	client, err := helper.GetContextMLFlowClient(ctx)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	prompts, err := app.repositories.MLFlowPrompts.ListPrompts(client, ctx)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.WriteJSON(w, http.StatusOK, prompts, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// MLFlowLoadPromptHandler handles GET /api/v1/mlflow/prompts/:name
func (app *App) MLFlowLoadPromptHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	name := ps.ByName("name")

	if name == "" {
		app.badRequestResponse(w, r, fmt.Errorf("missing required parameter: name"))
		return
	}

	client, err := helper.GetContextMLFlowClient(ctx)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	prompt, err := app.repositories.MLFlowPrompts.LoadPrompt(client, ctx, name)
	if err != nil {
		if mlflowsdk.IsNotFound(err) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.WriteJSON(w, http.StatusOK, prompt, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// MLFlowListPromptVersionsHandler handles GET /api/v1/mlflow/prompts/:name/versions
func (app *App) MLFlowListPromptVersionsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	name := ps.ByName("name")

	if name == "" {
		app.badRequestResponse(w, r, fmt.Errorf("missing required parameter: name"))
		return
	}

	client, err := helper.GetContextMLFlowClient(ctx)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	versions, err := app.repositories.MLFlowPrompts.ListPromptVersions(client, ctx, name)
	if err != nil {
		if mlflowsdk.IsNotFound(err) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.WriteJSON(w, http.StatusOK, versions, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// MLFlowRegisterPromptHandler handles POST /api/v1/mlflow/prompts
func (app *App) MLFlowRegisterPromptHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var req struct {
		Name     string `json:"name"`
		Template string `json:"template"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if req.Name == "" || req.Template == "" {
		app.badRequestResponse(w, r, fmt.Errorf("missing required parameters: name and template"))
		return
	}

	client, err := helper.GetContextMLFlowClient(ctx)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	prompt, err := app.repositories.MLFlowPrompts.RegisterPrompt(client, ctx, req.Name, req.Template)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.WriteJSON(w, http.StatusCreated, prompt, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
