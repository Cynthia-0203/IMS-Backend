package http

import (
	"fishnet/domain"
	confUsecase "fishnet/service/conf/usecase"
	usecase "fishnet/service/device/usecase"
)

var _deviceUsecase domain.DeviceUsecase
var _sensorUsecase domain.SensorUsecase
var _sensorTypeUsecase domain.SensorTypeUsecase
var _sensorDataUsecase domain.SensorDataUsecase

var _wordcaseUsecase domain.WordcaseUsecase

func init() {
	_deviceUsecase = usecase.NewDeviceUsecase()
	_sensorUsecase = usecase.NewSensorUsecase()
	_sensorTypeUsecase = usecase.NewSensorTypeUsecase()
	_sensorDataUsecase = usecase.NewSensorDataUsecase()

	_wordcaseUsecase = confUsecase.NewWordcaseUsecase()
}
