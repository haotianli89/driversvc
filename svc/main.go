package main

import (
	"github.com/haotianli89/driversvc/pb"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"github.com/uber-go/dosa"
	_ "github.com/uber-go/dosa/connectors/cassandra"
)


type Driver struct {
	dosa.Entity `dosa:"primaryKey=Id"`
	Id     dosa.UUID
	Name   string
	//Trips  []string
}


type Driversvc struct {

}


//func getDrivers() ([]*driversvc.Driver, error) {
//	registry, err := dosa.NewRegistrar("driverdb", "driversvc", &Driver{})
//	if err != nil {
//		panic(err.Error())
//	}
//
//	connector, err := dosa.GetConnector("cassandra", nil)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	client := dosa.NewClient(registry, connector)
//	if err := client.Initialize(context.TODO()); err != nil {
//		panic(err.Error())
//	}
//
//	var drivers []*driversvc.Driver
//	var object dosa.DomainObject
//	fields := []string{"id", "name"}
//
//	if _, err := client.MultiRead(context.TODO(), fields, object); err != nil {
//		panic(err.Error())
//	}
//
//
//	fmt.Println("object:", object)
//
//	return drivers, nil
//}


func getDrivers() ([]*driversvc.Driver, error) {
	cluster := gocql.NewCluster("172.17.0.2")
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
