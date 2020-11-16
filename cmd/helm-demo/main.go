package main

import (
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var (
	settings = cli.New()
	kubeconfig = "/Users/mojo/.kube/config"
	namespace = "default"
	charturl = "https://charts.bitnami.com/bitnami/airflow-6.7.1.tgz"
)

func main() {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Error("kube client init failed", err)
		return
	}
	d := driver.NewSecrets(clientset.CoreV1().Secrets(namespace))
	d.Log = logrus.Infof
	storage := storage.Init(d)

	cfg, _ := NewActionConfig(storage, config, namespace)
	client := action.NewList(cfg)
	client.SetStateMask()
	releases, err := client.Run()
	if err != nil {
		log.Fatal("error=>",err)
		return
	}

	for _, release := range releases {
		log.Println("release name:",release.Name, release.Chart.Metadata.Annotations)
	}
}

func install(cfg *action.Configuration)  {
	chart, err := loadChart(charturl)
	if err != nil {
		logrus.Error("load chart failed", err)
		return
	}
	installer := action.NewInstall(cfg)
	installer.ReleaseName = "mojo1"
	installer.Namespace = namespace

	chart.Metadata.Annotations["tenant"] = "uuid-key"
	rel, err := installer.Run(chart, chart.Values)
	if err != nil {
		logrus.Error("run install failed:", err)
		return
	}
	logrus.Info("annotation of release:", rel.Chart.Metadata.Annotations)
}


// NewActionConfig creates an action.Configuration, which can then be used to create Helm 3 actions.
// Among other things, the action.Configuration controls which namespace the command is run against.
func NewActionConfig(store *storage.Storage, config *rest.Config, namespace string) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)
	restClientGetter := NewConfigFlagsFromCluster(namespace, config)
	actionConfig.RESTClientGetter = restClientGetter
	actionConfig.KubeClient = kube.New(restClientGetter)
	actionConfig.Releases = store
	actionConfig.Log = logrus.Infof
	return actionConfig, nil
}

// NewConfigFlagsFromCluster returns ConfigFlags with default values set from within cluster.
func NewConfigFlagsFromCluster(namespace string, clusterConfig *rest.Config) *genericclioptions.ConfigFlags {
	impersonateGroup := []string{}
	insecure := false

	// CertFile and KeyFile must be nil for the BearerToken to be used for authentication and authorization instead of the pod's service account.
	return &genericclioptions.ConfigFlags{
		Insecure:         &insecure,
		Timeout:          stringptr("0"),
		Namespace:        stringptr(namespace),
		APIServer:        stringptr(clusterConfig.Host),
		CAFile:           stringptr(clusterConfig.CAFile),
		BearerToken:      stringptr(clusterConfig.BearerToken),
		ImpersonateGroup: &impersonateGroup,
	}
}


func stringptr(val string) *string {
	return &val
}