package routes

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/thuongtruong109/nolink/helpers"
)

type request struct {
	URL         string        `json:"url" binding:"required"`
	IP          string        `json:"ip" binding:"required"`
	Short string        `json:"short" binding:"omitempty,min=3,max=10"`
	Expiry      time.Duration `json:"expiry" binding:"omitempty,min=1,max=365"`
}

type Response struct {
	URL             string        `json:"url"`
	Short     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helpers.CANNOT_PARSE_ENV})
	}

	if body.IP == "" {
		body.IP = c.IP()
	}

	r2 := helpers.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(helpers.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(helpers.Ctx, c.IP(), os.Getenv("API_QUOTA"), 1*time.Second).Err()
	} else {
		val, _ = r2.Get(helpers.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt > 0 {
			limit, _ := r2.TTL(helpers.Ctx, c.IP()).Result()
			return helpers.ResponseErr(c, 503, fmt.Sprintf("rate limit exceeded, retry after: %d", limit/time.Nanosecond/time.Second))
		}
	}

	if !govalidator.IsURL(body.URL) {
		return helpers.ResponseErr(c, 404, helpers.INVALID_URL)
	}

	if !helpers.RemoveDomainError(body.URL) {
		return helpers.ResponseErr(c, 503, helpers.CANNOT_SHORTEN_URL)
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	var id string

	if body.Short == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.Short
	}

	r := helpers.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(helpers.Ctx, id).Result()
	if val != "" {
		return helpers.ResponseErr(c, 404, helpers.SHORT_LINK_EXISTS)
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = r.Set(helpers.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		return helpers.ResponseErr(c, 503, helpers.CANNOT_CONNECT_TO_DB)
	}

	// add shorted url to index
	err = r.SAdd(helpers.Ctx, helpers.INDEX+body.IP, id).Err()
	if err != nil {
		return helpers.ResponseErr(c, 500, helpers.CANNOT_CONNECT_TO_DB)
	}

	resp := Response{
		URL:             body.URL,
		Short:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}

	r2.Decr(helpers.Ctx, c.IP())

	val, _ = r2.Get(helpers.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(helpers.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.Short = os.Getenv("DOMAIN_RETURN") + "/" + id

	return c.Status(fiber.StatusCreated).JSON(resp)
}
