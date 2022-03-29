package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"forum-test/database"
	"forum-test/request"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)


func main() {

	env := loadEnv()

	// Redirect all http requests to https
	srvHTTP := &http.Server{
		Addr: env.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Connection", "close")
			// url := "https://" + req.Host + req.URL.Path
			url := "https://localhost:" + env.HTTPSport // this should be used only on localhost
			if len(req.URL.RawQuery) > 0 {
				url += "?" + req.URL.RawQuery
			}
			log.Printf("redirect to: %s", url)
			http.Redirect(w, req, url, http.StatusMovedPermanently)
		}),
	}
	go func() { log.Fatal(srvHTTP.ListenAndServe()) }()

	database.Initialize(env)

	go request.FillBucket()
	go request.EmptyBucket()

	//if there are no certs, create them
	if _, err := os.Stat("assets/cert.pem"); os.IsNotExist(err) {
		generateCertAndKeys()
	}

	if _, err := os.Stat("assets/key.pem"); os.IsNotExist(err) {
		generateCertAndKeys()
	}

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{ tls.CurveP521, tls.CurveP384, tls.CurveP256 },
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	srvHTTPS := &http.Server{
		Addr: env.HTTPSport,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	public := http.FileServer(http.Dir("./public"))

	//checking if DB is empty,if this is the case, populate database with dummy data
	categories, _ := database.GetCategories()
	posts, _ := database.GetAllPosts()
	if len(categories) == 0 && len(posts) == 0 {
		log.Println("Populating database with dummy data")
		_ = request.GenerateDummyData()
	}

	http.Handle("/public/", http.StripPrefix("/public/", public))

	http.HandleFunc("/login", request.Auth(request.Login, "guest"))

	http.HandleFunc("/register", request.Auth(request.Register, "guest"))

	http.HandleFunc("/category/", request.Auth(request.Category, "everyone"))

	http.HandleFunc("/post", request.Auth(request.Post, "everyone"))

	http.HandleFunc("/logout", request.Auth(request.Logout, "everyone"))

	http.HandleFunc("/config", request.Auth(request.Config, "everyone"))

	http.HandleFunc("/new-post", request.Auth(request.NewPost, "everyone"))

	http.HandleFunc("/myposts", request.Auth(request.MyPosts, "user"))

	http.HandleFunc("/delete", request.Auth(request.Delete, "user"))

	http.HandleFunc("/activity", request.Auth(request.LikedPosts, "user"))

	http.HandleFunc("/profile", request.Auth(request.Profile, "user"))

	http.HandleFunc("/connect", request.Auth(request.ConnectOauth, "guest"))

	http.HandleFunc("/", request.Auth(request.IndexHandler, "everyone"))

	log.Printf("Please go to: https://localhost%s/", env.HTTPSport)

	log.Fatal(srvHTTPS.ListenAndServeTLS("assets/cert.pem", "assets/key.pem"))
}

func generateCertAndKeys() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"kood/Jõhvi"},
		},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(3 * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		log.Fatal("Failed to encode certificate to PEM")
	}
	if err := os.WriteFile("assets/cert.pem", pemCert, 0644); err != nil {
		log.Fatal(err)
	}
	log.Print("wrote cert.pem\n")

	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if pemKey == nil {
		log.Fatal("Failed to encode key to PEM")
	}
	if err := os.WriteFile("assets/key.pem", pemKey, 0600); err != nil {
		log.Fatal(err)
	}
	log.Print("wrote key.pem\n")
}

func loadEnv() *database.Env {
	var env = database.Env{
		Port:      os.Getenv("SERVER_PORT"),
		HTTPSport: os.Getenv("SERVER_HTTPS_PORT"),
		SQLuser:   os.Getenv("SQLITE_USERNAME"),
		SQLpass:   os.Getenv("SQLITE_PASSWORD"),
		DBpath:    os.Getenv("SQLITE_DB_PATH"),
	}
	return &env
}