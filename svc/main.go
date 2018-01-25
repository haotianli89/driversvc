package main

import (
	"github.com/haotianli89/driversvc/pb"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"fmt"
	"github.com/gocql/gocql"
	"log"
)


type Driversvc struct {

}


func getDrivers() ([]*driversvc.Driver, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "driverdb"
	cluster.Consistency = gocql.One
	session, _ := cluster.CreateSession()
	defer session.Close()

	iter := session.Query("SELECT id, name FROM drivers").Iter()

	var resultId, resultName string
	var drivers []*driversvc.Driver

	for iter.Scan(&resultId, &resultName) {
		drivers = append(drivers, &driversvc.Driver{Id: resultId, Name: resultName})
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return drivers, nil
}


func getDriver(id string) (*driversvc.Driver, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "driverdb"
	session, _ := cluster.CreateSession()
	defer session.Close()

	var resultId, resultName string
	if err := session.Query("SELECT id, name FROM drivers where id = ?", id).Consistency(gocql.One).Scan(&resultId, &resultName); err != nil {
		fmt.Println(err)

		return nil, err
	}

	return &driversvc.Driver{Id: resultId, Name: resultName}, nil
}


func (s *Driversvc) GetDrivers(ctx context.Context, req *driversvc.GetDriversRequest, rsp *driversvc.GetDriversResponse) error {
	fmt.Println("requestDump: ", *req)

	if len(req.Id) == 0 {
		rsp.Drivers, _ = getDrivers()
		fmt.Println(getDrivers())
	} else {
		driver, _ := getDriver(req.Id)
		rsp.Drivers = append(rsp.Drivers, driver)
	}

	return nil
}


func main() {
	service := micro.NewService(
		micro.Name("driversvc"),
	)

	service.Init()

	driversvc.RegisterDriversvcHandler(service.Server(), new(Driversvc))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
