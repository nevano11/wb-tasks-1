package main

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

// Реализовать паттерн «адаптер» на любом примере

// Идея. Для запуска приложения требуется запустить несколько сервисов, каждый из которых
// запускается методом Start()
// Однако для старта Сервера необходимо его сконфигурировать и запустить Listening
// Для запуска с помощью Start() создается адаптер

// Сервер
type Server interface {
	Config(configFile string) error
	ListenAndServe() error
}

type WebServer struct{}

func (ws *WebServer) Config(configFile string) error {
	fmt.Println("Server configuring...")
	return nil
}
func (ws *WebServer) ListenAndServe() error {
	fmt.Println("Listening...")
	return nil
}

// Адаптер для сервера
type ServerAdapter struct {
	server Server
}

func (sa *ServerAdapter) Start() error {
	err := sa.server.Config("path/to/file")
	if err != nil {
		return err
	}
	err = sa.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

type NatsStreaming struct{}

func (ns *NatsStreaming) Start() error {
	fmt.Println("Nats streaming starting...")
	return nil
}

// Cервис
type Service interface {
	Start() error
}

type ApplicationStarter struct {
}

func (as *ApplicationStarter) StartServices(services []Service) error {
	var err error
	for _, service := range services {
		err = service.Start()
		if err != nil {
			return errors.New("Error on start services: " + err.Error())
		}
	}
	return nil
}

func Test21(t *testing.T) {
	services := []Service{
		&ServerAdapter{server: &WebServer{}},
		&NatsStreaming{},
	}

	appStarter := ApplicationStarter{}
	err := appStarter.StartServices(services)
	if err != nil {
		log.Fatalf("Failed to start services: %s", err.Error())
	}
}
