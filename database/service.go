package database

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetCountries     = "database.getCountries"
	MethodGetRegions       = "database.getRegions"
	MethodGetStreetsById   = "database.getStreetsById"
	MethodGetCountriesById = "database.getCountriesById"
	MethodGetCities        = "database.getCities"
	MethodGetCitiesById    = "database.getCitiesById"
	MethodGetUniversities  = "database.getUniversities"
	MethodGetSchools       = "database.getSchools"
	MethodGetSchoolClasses = "database.getSchoolClasses"
	MethodGetFaculties     = "database.getFaculties"
	MethodGetChairs        = "database.getChairs"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetCountries(r GetCountriesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCountries, r)
}

func (s *Service) GetRegions(r GetRegionsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetRegions, r)
}

func (s *Service) GetStreetsById(r GetStreetsByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetStreetsById, r)
}

func (s *Service) GetCountriesById(r GetCountriesByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCountriesById, r)
}

func (s *Service) GetCities(r GetCitiesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCities, r)
}

func (s *Service) GetCitiesById(r GetCitiesByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCitiesById, r)
}

func (s *Service) GetUniversities(r GetUniversitiesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetUniversities, r)
}

func (s *Service) GetSchools(r GetSchoolsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetSchools, r)
}

func (s *Service) GetSchoolClasses(r GetSchoolClassesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetSchoolClasses, r)
}

func (s *Service) GetFaculties(r GetFacultiesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetFaculties, r)
}

func (s *Service) GetChairs(r GetChairsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetChairs, r)
}
