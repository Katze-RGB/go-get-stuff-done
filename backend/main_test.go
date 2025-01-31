package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	tests := []struct {
		desc    string
		route   string
		expCode int
	}{
		{
			desc:    "200 from home route",
			route:   "/",
			expCode: 200,
		},
		{
			desc:    "gimme a 404",
			route:   "/doesnt-exist",
			expCode: 404,
		},
	}
	app := fiber.New()
	setupRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expCode, resp.StatusCode, test.desc)
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		desc    string
		route   string
		body    string
		expCode int
	}{
		{ //this is gonna error out bc I couldn't figure out how to mock a db in time, I thought it was pretty good for like 6 hours of go experience though.
			desc:    "creation of task from a req with proper params",
			route:   "/todo_task",
			body:    `"description":"a task that should work", "estimated_length":5, "priority":3`,
			expCode: 200,
		},
		{
			desc:    "bounce said req if it comes in goofy (negative time, priority out of expected range)",
			route:   "/todo_task",
			body:    `{"description":"a task that should fail", "estimated_length":-3, "priority":3}`,
			expCode: 400,
		},
	}
	app := fiber.New()
	setupRoutes(app)

	for _, test := range tests {
		resp, _ := app.Test(postRequest(test.route, test.body))
		assert.Equalf(t, test.expCode, resp.StatusCode, test.desc)
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		desc    string
		route   string
		expCode int
	}{
		{
			desc:    "200 to delete an task that exists",
			route:   "/todo_task/1",
			expCode: 200,
		},
		{
			desc:    "gimme a 404",
			route:   "/todo_task/999999",
			expCode: 404,
		},
	}
	app := fiber.New()
	setupRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest("DELETE", test.route, nil)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expCode, resp.StatusCode, test.desc)
	}
}

func TestComplete(t *testing.T) {
	tests := []struct {
		desc    string
		route   string
		expCode int
	}{
		{
			desc:    "200 for completing a task that exists",
			route:   "/complete_task/1",
			expCode: 200,
		},
		{
			desc:    "gimme a 404 for a task that doesn't exist",
			route:   "/complete_task/9999",
			expCode: 404,
		},
	}
	app := fiber.New()
	setupRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, nil)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expCode, resp.StatusCode, test.desc)
	}
}

func TestProductivityReport(t *testing.T) {
	tests := []struct {
		desc    string
		route   string
		expCode int
	}{
		{
			desc:    "200 for a day that has tasks and is in the past",
			route:   "/productivity-report/2025-02-3",
			expCode: 200,
		},
	}
	app := fiber.New()
	setupRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, -1)
		assert.Equalf(t, test.expCode, resp.StatusCode, test.desc)
	}
}

func postRequest(url string, body string) *http.Request {
	req := httptest.NewRequest("POST", url, bytes.NewBufferString(body))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return req
}

func bodyFromResponse[T any](t *testing.T, resp *http.Response) T {
	defer resp.Body.Close()
	var body T
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.Nil(t, err)
	return body
}
