package main
//
//import (
//	"github.com/sirupsen/logrus"
//	"net/url"
//	"strings"
//	"testing"
//)
//
//func TestURLParse(t *testing.T)  {
//	chart := "https://charts.bitnami.com/bitnami/bbb/airflow-6.7.1.tgz"
//	index := "https://charts.bitnami.com/bitnami"
//
//	indexURL, err := url.Parse(strings.TrimSpace(index))
//	if err != nil {
//		return
//	}
//	logrus.Info("index url:", indexURL)
//	chartURL, err := indexURL.Parse(strings.TrimSpace(chart))
//	if err != nil {
//		return
//	}
//
//	urlx, _ := url.ParseRequestURI(chartURL.String())
//	logrus.Info("chart url:", chartURL, urlx.String())
//}