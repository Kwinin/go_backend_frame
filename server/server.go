package server

import (
	"backend_ft_tcs/constant"
	"backend_ft_tcs/docs"
	"backend_ft_tcs/fisco"
	"backend_ft_tcs/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"strings"
)

type Server interface {
	Run(port string, configPath string) error
	Close() error
}

type defaultServer struct {
	name string

	conf        *serverConfig
	db          *gorm.DB
	engine      *gin.Engine
	fiscoClient fisco.FiscoService
}

func NewServer(name string) Server {
	s := new(defaultServer)
	s.name = name
	return s
}

func (rs *defaultServer) Run(port string, configPath string) error {
	// config
	if err := rs.config(configPath); err != nil {
		return fmt.Errorf("rs.config(): %s", err.Error())
	}

	// mysql
	if err := rs.dbClient(); err != nil {
		return fmt.Errorf("rs.dbClient(): %s", err.Error())
	}

	// router
	if err := rs.initRouter(); err != nil {
		return fmt.Errorf("rs.router(): %s", err.Error())
	}

	// init
	if err := rs.init(); err != nil {
		return fmt.Errorf("rs.init(): %s", err.Error())
	}

	// port
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return rs.engine.Run(port)
}

func (rs *defaultServer) Close() error {

	if rs.db != nil {
		if err := rs.db.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (rs *defaultServer) config(configPath string) error {
	rs.conf = new(serverConfig)
	err := configor.Load(rs.conf, configPath)
	if err != nil {
		return err
	}
	return nil
}

func (rs *defaultServer) dbClient() error {
	db, err := GetDBConnection(rs.conf.DbConfig)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}
	rs.db = db
	return nil
}

func (rs *defaultServer) initRouter() error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.CORSMiddleware())
	rs.engine = r
	if rs.conf.Debug {
		// swagger文档
		rs.SwaggerDoc(r)
	}
	var hander = func(c *gin.Context) {
		c.Set(constant.ContextDb, rs.db)
		c.Next()
	}
	rs.engine.Use(hander)
	rs.routers()

	return nil
}

func (rs *defaultServer) SwaggerDoc(ctx *gin.Engine) {
	docs.SwaggerInfo.Title = AppTitle
	docs.SwaggerInfo.Description = AppName
	docs.SwaggerInfo.Version = AppVersion
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = rs.conf.SwagConfig.BasePath
	docs.SwaggerInfo.Host = rs.conf.SwagConfig.Host

	ctx.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (rs *defaultServer) init() error {
	log.SetLogLevel(rs.conf.LogLevel)

	return nil
}
