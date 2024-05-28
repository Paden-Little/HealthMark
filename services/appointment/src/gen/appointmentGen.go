// Package appointment provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version 2.1.0 DO NOT EDIT.
package appointment

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// Appointment defines model for Appointment.
type Appointment struct {
	Date        *string `json:"date,omitempty"`
	Description *string `json:"description,omitempty"`
	EndTime     *string `json:"end_time,omitempty"`
	Id          *string `json:"id,omitempty"`
	Patient     *string `json:"patient,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	Service     *int    `json:"service,omitempty"`
	StartTime   *string `json:"start_time,omitempty"`
}

// AppointmentResponse defines model for AppointmentResponse.
type AppointmentResponse struct {
	Id      *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
}

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse struct {
	Message *string `json:"message,omitempty"`
}

// NotFoundResponse defines model for NotFoundResponse.
type NotFoundResponse struct {
	Id      *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
}

// PostAppointmentJSONRequestBody defines body for PostAppointment for application/json ContentType.
type PostAppointmentJSONRequestBody = Appointment

// PatchAppointmentAppointmentIdJSONRequestBody defines body for PatchAppointmentAppointmentId for application/json ContentType.
type PatchAppointmentAppointmentIdJSONRequestBody = Appointment

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new appointment
	// (POST /appointment)
	PostAppointment(c *gin.Context)
	// Delete a specific appointment
	// (DELETE /appointment/{appointmentId})
	DeleteAppointmentAppointmentId(c *gin.Context, appointmentId string)
	// Get a specific appointment
	// (GET /appointment/{appointmentId})
	GetAppointmentAppointmentId(c *gin.Context, appointmentId string)
	// Update an appointment by ID
	// (PATCH /appointment/{appointmentId})
	PatchAppointmentAppointmentId(c *gin.Context, appointmentId string)
	// Health check
	// (GET /health)
	CheckHealth(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostAppointment operation middleware
func (siw *ServerInterfaceWrapper) PostAppointment(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostAppointment(c)
}

// DeleteAppointmentAppointmentId operation middleware
func (siw *ServerInterfaceWrapper) DeleteAppointmentAppointmentId(c *gin.Context) {

	var err error

	// ------------- Path parameter "appointmentId" -------------
	var appointmentId string

	err = runtime.BindStyledParameterWithOptions("simple", "appointmentId", c.Param("appointmentId"), &appointmentId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter appointmentId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteAppointmentAppointmentId(c, appointmentId)
}

// GetAppointmentAppointmentId operation middleware
func (siw *ServerInterfaceWrapper) GetAppointmentAppointmentId(c *gin.Context) {

	var err error

	// ------------- Path parameter "appointmentId" -------------
	var appointmentId string

	err = runtime.BindStyledParameterWithOptions("simple", "appointmentId", c.Param("appointmentId"), &appointmentId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter appointmentId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAppointmentAppointmentId(c, appointmentId)
}

// PatchAppointmentAppointmentId operation middleware
func (siw *ServerInterfaceWrapper) PatchAppointmentAppointmentId(c *gin.Context) {

	var err error

	// ------------- Path parameter "appointmentId" -------------
	var appointmentId string

	err = runtime.BindStyledParameterWithOptions("simple", "appointmentId", c.Param("appointmentId"), &appointmentId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter appointmentId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchAppointmentAppointmentId(c, appointmentId)
}

// CheckHealth operation middleware
func (siw *ServerInterfaceWrapper) CheckHealth(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CheckHealth(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/appointment", wrapper.PostAppointment)
	router.DELETE(options.BaseURL+"/appointment/:appointmentId", wrapper.DeleteAppointmentAppointmentId)
	router.GET(options.BaseURL+"/appointment/:appointmentId", wrapper.GetAppointmentAppointmentId)
	router.PATCH(options.BaseURL+"/appointment/:appointmentId", wrapper.PatchAppointmentAppointmentId)
	router.GET(options.BaseURL+"/health", wrapper.CheckHealth)
}
