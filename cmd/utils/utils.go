package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"
	"terrakube/client/client"
	"terrakube/config"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/spf13/viper"
)

func NewClient() *client.Client {
	baseURL, err := url.Parse(viper.GetString("api-url"))
	if err != nil {
		fmt.Printf("Error parsing API URL: %v\n", err)
		os.Exit(1)
	}

	return client.NewClient(nil, viper.GetString("pat"), baseURL)
}

func RenderOutput(result interface{}) {
	format := config.CliConfig.OutputFormat
	switch format {
	case "json":
		printJSON, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatal("Failed to generate json", err)
		}
		fmt.Printf("%s\n", string(printJSON))
	case "tsv":
		data, _ := splitInterface(result)
		for _, v := range data {
			fmt.Println(strings.Join(v[:], "\t"))
		}
	case "table":
		data, header := splitInterface(result)
		if len(data) > 0 {
			table := tablewriter.NewWriter(os.Stdout)
			table.Header(header)
			err := table.Bulk(data)
			if err != nil {
				return
			}
			err = table.Render()
			if err != nil {
				return
			}
		}
	case "markdown":
		data, header := splitInterface(result)
		if len(data) > 0 {
			table := tablewriter.NewTable(os.Stdout,
				tablewriter.WithRenderer(renderer.NewMarkdown()),
			)
			table.Header(header)
			err := table.Bulk(data)
			if err != nil {
				return
			}
			err = table.Render()
			if err != nil {
				return
			}
		}
	case "none":

	}
}

func splitInterface(input interface{}) ([][]string, []string) {
	reflectData := reflect.ValueOf(input)
	headers := make([]string, 0)
	headers = append(headers, "ID")
	result := make([][]string, 0)
	if reflectData.Kind() == reflect.Slice {
		for i := 0; i < reflectData.Len(); i++ {
			data := reflectData.Index(i).Interface()
			d := reflect.Indirect(reflect.ValueOf(data))

			row := make([]string, 0)
			id := d.FieldByName("ID").String()
			row = append(row, id)

			attr := reflect.Indirect(reflect.ValueOf(d.FieldByName("Attributes").Interface()))
			for j := 0; j < attr.NumField(); j++ {
				if i == 0 {
					headers = append(headers, attr.Type().Field(j).Name)
				}
				row = append(row, attr.Field(j).String())
			}
			result = append(result, row)
		}
	} else {
		d := reflect.Indirect(reflectData)
		row := make([]string, 0)
		id := d.FieldByName("ID").String()
		row = append(row, id)

		attr := reflect.Indirect(reflect.ValueOf(d.FieldByName("Attributes").Interface()))
		for j := 0; j < attr.NumField(); j++ {
			headers = append(headers, attr.Type().Field(j).Name)
			row = append(row, attr.Field(j).String())
		}
		result = append(result, row)

	}
	return result, headers
}
