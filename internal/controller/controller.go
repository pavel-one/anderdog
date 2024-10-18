package controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/oschwald/geoip2-golang"
	"github.com/pavel-one/anderdog/internal/repository"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	rep *repository.VisitRepository
	geo *geoip2.Reader
}

func New(rep *repository.VisitRepository, geo *geoip2.Reader) *Controller {
	return &Controller{
		rep: rep,
		geo: geo,
	}
}

func (c *Controller) Index(ctx fiber.Ctx) error {
	rec, err := c.geo.City(net.ParseIP(ctx.IP()))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	cityName := ""
	if rec.City.Names != nil {
		cityName = rec.City.Names["ru"]
	}

	i, err := c.rep.Create(repository.Visit{
		Time: time.Now(),
		IP:   ctx.IP(),
		City: cityName,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	f, err := os.ReadFile("frontend/index.html")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	iStr := strconv.Itoa(i)
	out := strings.ReplaceAll(string(f), `%d`, iStr)

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return ctx.SendString(out)
}
