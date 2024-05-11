package main

import (
	"fmt"
	"net/http"
	"orange-provider-wrapper/cmd"
	"orange-provider-wrapper/service"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"orange-provider-wrapper/config"
	"orange-provider-wrapper/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/urfave/cli"
)

const defaultPort = "8080"

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "orange provider wrapper service"
	app.Action = startAgent
	app.Flags = []cli.Flag{
		cmd.LogLevelFlag,
		cmd.LogDirFlag,
		cmd.PortFlag,
		cmd.ConfigFileFlag,
		cmd.OperationFlag,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func startAgent(ctx *cli.Context) {
	port := ctx.GlobalString(cmd.GetFlagName(cmd.PortFlag))
	if port == "" {
		port = defaultPort
	}

	initLog(ctx)
	cfgfile := ctx.GlobalString(cmd.GetFlagName(cmd.ConfigFileFlag))
	err := config.LoadConfig(cfgfile)
	if err != nil {
		fmt.Println("error on load config")
		panic(err)
	}

	op := ctx.GlobalString(cmd.GetFlagName(cmd.OperationFlag))
	if len(op) > 0 {
		switch strings.ToLower(op) {
		case "new-wallet":
			err = service.NewWallet()
			if err != nil {
				fmt.Printf("create new wallet failed: %v", err)
			}
		case "register-did":
			err = service.RegisterDID()
			if err != nil {
				fmt.Printf("register did failed: %v", err)
			}
		default:
			fmt.Println("invalid operation")
		}
		os.Exit(1)

	}
	err = service.InitAllServices()
	if err != nil {
		panic(err)
	}
	go CheckLogSize(ctx)
	router := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedHeaders:   []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
		ExposedHeaders:   []string{"Content-Length", "token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar"},
		AllowCredentials: false,
		MaxAge:           172800, // Maximum value not ignored by any of major browsers
		// Debug:true,
	}))

	router.Route("/", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowedHeaders:   []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
			ExposedHeaders:   []string{"Content-Length", "token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar"},
			AllowCredentials: false,
			MaxAge:           172800, // Maximum value not ignored by any of major browsers
			// Debug:true,
		}))
		for _, cfg := range config.GlobalConfig.APIConfigs {

			if strings.EqualFold(cfg.ProviderType, "DP") {
				fmt.Printf("register DP %s path: %v\n", "POST", cfg.ServerPath)
				r.Post(cfg.ServerPath, service.GlobalProxyService.GenerateDPHandleFunc(cfg))
			} else {
				fmt.Printf("register AP %s path: %v\n", "POST", cfg.ServerPath)
				r.Post(cfg.ServerPath, service.GlobalProxyService.GenerateAPHandleFunc(cfg))
			}
		}

	})

	go signalHandle()
	log.Infof("staring restful at port:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func initLog(ctx *cli.Context) {
	logLevel := ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag))
	disableLogFile := ctx.GlobalBool(cmd.GetFlagName(cmd.DisableLogFileFlag))
	if disableLogFile {
		log.InitLog(logLevel, log.Stdout)
	} else {
		logFileDir := ctx.String(cmd.GetFlagName(cmd.LogDirFlag))
		logFileDir = filepath.Join(logFileDir, "") + string(os.PathSeparator)
		log.InitLog(logLevel, logFileDir, log.Stdout)
	}
}
func CheckLogSize(ctx *cli.Context) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			isNeedNewFile := log.CheckIfNeedNewFile()
			if isNeedNewFile {
				log.ClosePrintLog()
				log.InitLog(ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag)), log.PATH, log.Stdout)
			}
		}
	}
}
func signalHandle() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("get a signal: stop the rest gateway process", si.String())
			os.Exit(1)
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
