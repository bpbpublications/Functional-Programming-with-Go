package weather

import "testing"

type MockWeatherAPI struct {
    FakeTemp float64
}

func (m *MockWeatherAPI) FetchTemperature(city string) (float64, error) {
    return m.FakeTemp, nil // Return a fake temperature
}

func TestGetTemperature(t *testing.T) {
    // Set up our mock with the desired fake temperature
    mockAPI := &MockWeatherAPI{FakeTemp: 25.0}
    service := NewWeatherService(mockAPI)

    temp, err := service.GetTemperature("Tokyo")
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Assert that the temperature returned by our service matches the mock's fake temperature
    if temp != mockAPI.FakeTemp {
        t.Errorf("Expected temperature to be %v, got %v", mockAPI.FakeTemp, temp)
    }
}
