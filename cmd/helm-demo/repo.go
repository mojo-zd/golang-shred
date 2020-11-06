package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/chart"
	helm3loader "helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/repo"
	"io/ioutil"
	"net/http"
	"net/url"
	"sigs.k8s.io/yaml"
	"strings"
	"time"
)

var (
	index           = "/Users/mojo/Downloads/index.yaml"
	ErrNoAPIVersion = errors.New("no api version")
)

func loadIndex() (*repo.IndexFile, error) {
	i := &repo.IndexFile{}
	data, err := ioutil.ReadFile(index)
	if err != nil {
		logrus.Error(err)
		return i, err
	}

	if err := yaml.UnmarshalStrict(data, i); err != nil {
		return i, err
	}
	i.SortEntries()
	if i.APIVersion == "" {
		return i, ErrNoAPIVersion
	}
	return i, nil
}

func resolveChartURL(index, chart string) (string, error) {
	indexURL, err := url.Parse(strings.TrimSpace(index))
	if err != nil {
		return "", err
	}
	chartURL, err := indexURL.Parse(strings.TrimSpace(chart))
	if err != nil {
		return "", err
	}
	return chartURL.String(), nil
}

func loadChart(chartURL string) (*chart.Chart, error) {
	req, err := http.NewRequest("GET", chartURL, nil)
	if err != nil {
		return nil, err
	}

	logrus.Info("start to download tar")
	cli := http.Client{Timeout: time.Minute * 5}
	res, err := cli.Do(req)
	logrus.Info("finish to download tar")
	if err != nil {
		return nil, err
	}
	data, err := readResponseBody(res)
	if err != nil {
		return nil, err
	}
	// We only return an error when loading using the helm2loader (ie. chart v1)
	// if we require v1 support, otherwise we continue to load using the
	// helm3 v2 loader.

	helm3Chart, err := helm3loader.LoadArchive(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return helm3Chart, nil
}

func readResponseBody(res *http.Response) ([]byte, error) {
	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("chart download request failed")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
