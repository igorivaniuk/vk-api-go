package database

type GetCountriesRequest struct {
	NeedAll bool   `param:"need_all"`
	Code    string `param:"code"`
	Offset  int    `param:"offset"`
	Count   int    `param:"count"`
}

type GetRegionsRequest struct {
	CountryId int    `param:"country_id"`
	Q         string `param:"q"`
	Offset    int    `param:"offset"`
	Count     int    `param:"count"`
}

type GetStreetsByIdRequest struct {
	StreetIds string `param:"street_ids"`
}

type GetCountriesByIdRequest struct {
	CountryIds string `param:"country_ids"`
}

type GetCitiesRequest struct {
	CountryId int    `param:"country_id"`
	RegionId  int    `param:"region_id"`
	Q         string `param:"q"`
	NeedAll   bool   `param:"need_all"`
	Offset    int    `param:"offset"`
	Count     int    `param:"count"`
}

type GetCitiesByIdRequest struct {
	CityIds string `param:"city_ids"`
}

type GetUniversitiesRequest struct {
	Q         string `param:"q"`
	CountryId int    `param:"country_id"`
	CityId    int    `param:"city_id"`
	Offset    int    `param:"offset"`
	Count     int    `param:"count"`
}

type GetSchoolsRequest struct {
	Q      string `param:"q"`
	CityId int    `param:"city_id"`
	Offset int    `param:"offset"`
	Count  int    `param:"count"`
}

type GetSchoolClassesRequest struct {
	CountryId int `param:"country_id"`
}

type GetFacultiesRequest struct {
	UniversityId int `param:"university_id"`
	Offset       int `param:"offset"`
	Count        int `param:"count"`
}

type GetChairsRequest struct {
	FacultyId int `param:"faculty_id"`
	Offset    int `param:"offset"`
	Count     int `param:"count"`
}
