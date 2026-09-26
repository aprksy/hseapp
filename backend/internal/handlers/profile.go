package handlers

import (
"encoding/json"
"net/http"

"backend/internal/services"
)

type ProfileHandler struct {
profileService *services.ProfileService
}

type ProfileResponse struct {
ID        string `json:"id"`
Email     string `json:"email"`
Name      string `json:"name"`
Role      string `json:"role"`
TenantID  string `json:"tenant_id"`
SiteIDs   []string `json:"site_ids"`
Language  string `json:"language"`
Theme     string `json:"theme"`
}

type UpdateProfileRequest struct {
Name     *string `json:"name,omitempty"`
Language *string `json:"language,omitempty"`
Theme    *string `json:"theme,omitempty"`
}

type PreferencesResponse struct {
Language string `json:"language"`
Theme    string `json:"theme"`
}

type UpdatePreferencesRequest struct {
Language *string `json:"language,omitempty"`
Theme    *string `json:"theme,omitempty"`
}

func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
return &ProfileHandler{
profileService: profileService,
}
}

// GetProfile returns the current user's profile
// @Summary Get user profile
// @Description Returns the authenticated user's profile information
// @Tags Profile
// @Produce json
// @Security BearerAuth
// @Success 200 {object} ProfileResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/profile [get]
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}

// Get user ID from context (set by auth middleware)
userID := r.Context().Value("user_id").(string)
if userID == "" {
http.Error(w, "Unauthorized", http.StatusUnauthorized)
return
}

profile, err := h.profileService.GetProfile(r.Context(), userID)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

response := ProfileResponse{
ID:       profile.ID,
Email:    profile.Email,
Name:     profile.Name,
Role:     profile.Role,
TenantID: profile.TenantID,
SiteIDs:  profile.SiteIDs,
Language: profile.Language,
Theme:    profile.Theme,
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}

// UpdateProfile updates the user's profile
// @Summary Update user profile
// @Description Updates the authenticated user's profile information
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body UpdateProfileRequest true "Profile update data"
// @Success 200 {object} ProfileResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/profile [put]
func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPut {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}

userID := r.Context().Value("user_id").(string)
if userID == "" {
http.Error(w, "Unauthorized", http.StatusUnauthorized)
return
}

var req UpdateProfileRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, "Invalid request body", http.StatusBadRequest)
return
}

profile, err := h.profileService.UpdateProfile(r.Context(), userID, req)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

response := ProfileResponse{
ID:       profile.ID,
Email:    profile.Email,
Name:     profile.Name,
Role:     profile.Role,
TenantID: profile.TenantID,
SiteIDs:  profile.SiteIDs,
Language: profile.Language,
Theme:    profile.Theme,
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}

// GetPreferences returns the user's preferences
// @Summary Get user preferences
// @Description Returns the authenticated user's preferences (language, theme)
// @Tags Preferences
// @Produce json
// @Security BearerAuth
// @Success 200 {object} PreferencesResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/preferences [get]
func (h *ProfileHandler) GetPreferences(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}

userID := r.Context().Value("user_id").(string)
if userID == "" {
http.Error(w, "Unauthorized", http.StatusUnauthorized)
return
}

prefs, err := h.profileService.GetPreferences(r.Context(), userID)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

response := PreferencesResponse{
Language: prefs.Language,
Theme:    prefs.Theme,
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}

// UpdatePreferences updates the user's preferences
// @Summary Update user preferences
// @Description Updates the authenticated user's preferences (language, theme)
// @Tags Preferences
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body UpdatePreferencesRequest true "Preference update data"
// @Success 200 {object} PreferencesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/preferences [put]
func (h *ProfileHandler) UpdatePreferences(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPut {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}

userID := r.Context().Value("user_id").(string)
if userID == "" {
http.Error(w, "Unauthorized", http.StatusUnauthorized)
return
}

var req UpdatePreferencesRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, "Invalid request body", http.StatusBadRequest)
return
}

prefs, err := h.profileService.UpdatePreferences(r.Context(), userID, req)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

response := PreferencesResponse{
Language: prefs.Language,
Theme:    prefs.Theme,
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}
