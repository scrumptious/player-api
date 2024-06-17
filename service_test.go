package main_test

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/scrumptious/weather-service/internal/data"
	"github.com/scrumptious/weather-service/internal/service"
)

var _ = Describe("Service", func() {
	var res1, res2 *data.WeatherData
	var err1, err2, err3, err4 error
	var ws service.Service
	BeforeEach(func() {
		ws = service.NewWeatherService(fmt.Sprintf("%s&res=daily", "http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/all/json/locationID?key=3768a301-4afa-4038-8ce0-c1eacf4207a4"))
	})

	Describe("Getting weather", func() {
		Context("with locationID equal 3080", func() {
			It("should contain weather data", func(ctx SpecContext) {
				ctxV1 := context.WithValue(ctx, "locationID", "3080")
				res1, err1 = ws.GetWeather(ctxV1)

				Expect(err1).ToNot(HaveOccurred())
				Expect(err2).ToNot(HaveOccurred())

				Expect(res1.Day).ToNot(Equal(""))
				Expect(res1.Imperial).ToNot(Equal(""))
				Expect(res1.MaxUV).ToNot(Equal(""))
				Expect(res1.Humidity).ToNot(Equal(""))
				Expect(res1.WindDirection).ToNot(Equal(""))
				Expect(res1.WindSpeed).ToNot(Equal(""))
				Expect(res1.Temperature).ToNot(Equal(""))
				Expect(res1.LocationID).To(Equal(3080))
			})
		})

		Context("with locationID equal 99060", func() {
			It("should contain weather data", func(ctx SpecContext) {
				ctxV2 := context.WithValue(ctx, "locationID", "99060")
				res2, err3 = ws.GetWeather(ctxV2)

				Expect(err3).ToNot(HaveOccurred())
				Expect(err4).ToNot(HaveOccurred())

				Expect(res2.Day).ToNot(Equal(""))
				Expect(res2.Imperial).ToNot(Equal(""))
				Expect(res2.MaxUV).ToNot(Equal(""))
				Expect(res2.Humidity).ToNot(Equal(""))
				Expect(res2.WindDirection).ToNot(Equal(""))
				Expect(res2.WindSpeed).ToNot(Equal(""))
				Expect(res2.Temperature).ToNot(Equal(""))
				Expect(res2.LocationID).To(Equal(99060))
			})
		})
	})
})
