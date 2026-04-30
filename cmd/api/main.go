package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomes800/password-manager/database"
	"github.com/gomes800/password-manager/internal/handler"
	"github.com/gomes800/password-manager/internal/repository"
	"github.com/gomes800/password-manager/internal/service"
	"github.com/gomes800/password-manager/security"
)

func main() {
	db, err := database.InitDb("test.db")
	if err != nil {
		log.Fatal("Error starting database: %v", err)
	}
	defer db.Close()

	credentialRepo := repository.NewCredentialRepository(db)

	credentialService := service.NewCredentialService(credentialRepo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := credentialRepo.CreateTable(ctx); err != nil {
		log.Fatal("Error creating table: %v", err)
	}
	log.Println("Database started")

	credentialHandler := handler.NewCredentialHandler(credentialService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /credentials", credentialHandler.Save)
	mux.HandleFunc("GET /credentials/{id}", credentialHandler.GetByID)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		log.Println("Starting server on port 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server gracefully...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")

	password1 := "aleatoriosenha123456"
	bankPassword := []byte("senhasegura12345@")

	password2 := "churinchurin"
	lifePassword := []byte("senhahipersegura123456#")

	key1, salt1, err := security.HashPassword(password1)
	if err != nil {
		panic(err)
	}

	key2, salt2, err := security.HashPassword(password2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("key1: %x\n", key1)
	fmt.Printf("salt1: %x\n", salt1)

	fmt.Printf("key2: %x\n", key2)
	fmt.Printf("salt2: %x\n", salt2)

	ciphertext, nonce, err := security.Encrypt(key1, bankPassword)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %x\n", ciphertext)

	plaintext, err := security.Decrypt(key1, ciphertext, nonce)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext)

	ciphertext2, nonce2, err := security.Encrypt(key2, lifePassword)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %x\n", ciphertext2)

	plaintext2, err := security.Decrypt(key1, ciphertext2, nonce2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext2)
}
