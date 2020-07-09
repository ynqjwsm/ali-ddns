package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
)

var(
	address string
	port string
	zone string
	accessKeyId string
	accessKeySecret string
	
)

func updateRecordAddress(domain string, record string, ip string, id string, sec string) string {
	client, err := alidns.NewClientWithAccessKey(zone, id, sec)
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = domain
	if err != nil {
		return "list domain records failed."
	}
	response, _ := client.DescribeDomainRecords(request)
	for _, rr := range response.DomainRecords.Record{
		if rr.RR == record {
			updateRequest := alidns.CreateUpdateDomainRecordRequest()
			updateRequest.RR = rr.RR
			updateRequest.RecordId = rr.RecordId
			updateRequest.Type = rr.Type
			updateRequest.Value = ip
			updateResponse , _ := client.UpdateDomainRecord(updateRequest)
			return updateResponse.GetHttpContentString()
		}
	}

	return "record not found."
}

func update(w http.ResponseWriter, r *http.Request)  {

	if "" == r.FormValue("domain") || "" == r.FormValue("record") || "" == r.FormValue("address") {
		fmt.Fprintln(w, "parameter [domain], [record], [address] can't be empty.")
		fmt.Fprintln(w, "support parameters :[domain], [record], [address] [id], [sec]")
		fmt.Fprintln(w, "eg :http://ip:port/update?domain=your.domian&record=example&address=8.8.8.8")
		fmt.Fprintln(w, "eg :http://ip:port/update?domain=your.domian&record=example&address=8.8.8.8&id=yoru_access_key_id&sec=your_access_key_secret")

		setup := `
Install as a service:
edit /etc/systemd/system/ali-ddns.service

[Unit]
Description=Ali-ddns
[Service]
Type=simple
PIDFile=/var/run/ali-ddns.pid
ExecStart=/path/to/your/file/ali-ddns server -a {bind_address} -p {bind_port} -i {AccessKeyId} -s {AccessKeySecret} -z {zone}
User=root
Group=root
[Install]
WantedBy=multi-user.target`

		fmt.Fprintln(w, setup)
		return
	}
	domain := r.FormValue("domain")
	record := r.FormValue("record")
	address := r.FormValue("address")

	finalAccessKeyId := r.FormValue("id")
	if "" != accessKeyId {
		finalAccessKeyId = accessKeyId
	}
	if "" == finalAccessKeyId{
		fmt.Fprint(w, "AccessKeyId error.")
		return
	}

	finalAccessKeySecret := r.FormValue("sec")
	if "" != accessKeyId {
		finalAccessKeySecret = accessKeySecret
	}
	if "" == finalAccessKeySecret{
		fmt.Fprint(w, "AccessKeySecret error.")
		return
	}
	fmt.Fprint(w, updateRecordAddress(domain, record, address, finalAccessKeyId, finalAccessKeySecret))
}

func main() {

	app := &cli.App{
		Name: "ali-ddns",
		Usage: "update you dns record of ali cloud.",
		Version: "0.0.1",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "server",
				Usage: "start a http server, url is http://host:port/update",
				Action: func(context *cli.Context) error {
					fullAddress := address + ":" + port
					fmt.Print("Server is running on ", fullAddress)
					http.HandleFunc("/update", update)
					http.ListenAndServe(fullAddress, nil)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Aliases: []string{"a"},
						Usage: "address to bind, default is 0.0.0.0, bind all interfaces.",
						Value: "0.0.0.0",
						DefaultText: "0.0.0.0",
						Destination: &address,
					},
					&cli.StringFlag{
						Name:  "port",
						Aliases: []string{"p"},
						Usage: "port to bind, default is 8724.",
						Value: "8724",
						DefaultText: "8724",
						Destination: &port,
					},
					&cli.StringFlag{
						Name:  "zone",
						Aliases: []string{"z"},
						Usage: "zone to connect , default is cn-hangzhou.",
						Value: "cn-hangzhou",
						DefaultText: "cn-hangzhou",
						Destination: &zone,
					},&cli.StringFlag{
						Name:  "accessKeyId",
						Aliases: []string{"i"},
						Usage: "ali cloud accessKeyId.",
						Value: "",
						Destination: &accessKeyId,
					},&cli.StringFlag{
						Name:  "accessKeySecret",
						Aliases: []string{"s"},
						Usage: "ali cloud accessKeySecret.",
						Value: "",
						Destination: &accessKeySecret,
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
