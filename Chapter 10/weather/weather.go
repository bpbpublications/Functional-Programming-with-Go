package weather

type WeatherAPI interface {
    FetchTemperature(city string) (float64, error)
}

type WeatherService struct {
    API WeatherAPI
}

func NewWeatherService(api WeatherAPI) *WeatherService {
    return &WeatherService{API: api}
}

func (s *WeatherService) GetTemperature(city string) (float64, error) {
    return s.API.FetchTemperature(city)
}
