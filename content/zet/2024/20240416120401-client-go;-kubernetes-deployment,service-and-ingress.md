+++
title = "Client-go; kubernetes deployment,service and ingress"
categories = ["zet"]
tags = ["zet"]
slug = "client-go;-kubernetes-deployment,service-and-ingress"
date = "2024-04-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Client-go; kubernetes deployment,service and ingress

How to create a simple deployment exposed with an ingress in Kubernetes using
the `client-go` SDK.

## Preconditions:

- k3s
- traefik
- `echo.k3s.lcl` mapped to local IP in `/etc/hosts` (if demo purposes only)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"k8s.io/apimachinery/pkg/util/intstr"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	client := k8sClient()

	err := createDeployment(
		client,
		"default",
		"echo",
		2,
		"ealen/echo-server",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = createService(
		client,
		"default",
		"echo-svc",
		3000,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = createIngress(
		client,
		"default",
		"echo-ingress",
		3000,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func createDeployment(
	client *kubernetes.Clientset,
	ns string,
	deploymentName string,
	replicas int32,
	image string,
) error {
	dc := client.AppsV1().Deployments(ns)

	dp := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									Name:          deploymentName,
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 8888,
								},
							},
						},
					},
				},
			},
		},
	}

	slog.Info("deployment", "deployment", deploymentName, "status", "creating")
	result, err := dc.Create(context.Background(), dp, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	slog.Info("deployment", "deployment", result.GetObjectMeta().GetName(), "status", "created")
	return nil
}

func createService(client *kubernetes.Clientset, ns string, serviceName string, port int32) error {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
			Labels: map[string]string{
				"app": "echo",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port:       port,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: 80},
				},
			},
			Selector: map[string]string{
				"app": "echo",
			},
		},
	}

	slog.Info("service", "service", serviceName, "status", "creating")
	service, err := client.CoreV1().Services(ns).Create(context.TODO(), svc, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	slog.Info("service", "service", service.GetObjectMeta().GetName(), "status", "created")
	return nil
}
func Ptr[T any](v T) *T {
	return &v
}

func createIngress(client *kubernetes.Clientset, ns string, ingressName string, port int32) error {
	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressName,
			Namespace: ns,
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: Ptr("traefik"),
			Rules: []networkingv1.IngressRule{
				{
					Host: "echo.k3s.lcl",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: Ptr(networkingv1.PathTypePrefix),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "echo-svc",
											Port: networkingv1.ServiceBackendPort{
												Number: port,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	slog.Info("ingress", "ingress", ingressName, "status", "creating")
	result, err := client.NetworkingV1().
		Ingresses(ingress.Namespace).
		Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	slog.Info("ingress", "ingress", result.GetObjectMeta().GetName(), "status", "created")
	return nil
}

func k8sClient() *kubernetes.Clientset {
  // If you are using $HOME/.kube/config uncomment this and remove the
  // os.Getenv("KUBECONFIG") line
	//userHomeDir, err := os.UserHomeDir()
	//if err != nil {
	//	fmt.Printf("error getting user home dir: %v\n", err)
	//	os.Exit(1)
	//}
	//kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	kubeConfigPath := os.Getenv("KUBECONFIG")
	fmt.Printf("Using kubeconfig: %s\n", kubeConfigPath)

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("Error getting kubernetes config: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("error getting kubernetes config: %v\n", err)
		os.Exit(1)
	}
	return clientset
}
```

Run with `go run .` and if you have a connection to the cluster, it should output:

```shell
Using kubeconfig: /etc/rancher/k3s/k3s.yaml
2024/04/16 22:07:07 INFO deployment deployment=echo status=creating
2024/04/16 22:07:07 INFO deployment deployment=echo status=created
2024/04/16 22:07:07 INFO service service=echo-svc status=creating
2024/04/16 22:07:07 INFO service service=echo-svc status=created
2024/04/16 22:07:07 INFO ingress ingress=echo-ingress status=creating
2024/04/16 22:07:07 INFO ingress ingress=echo-ingress status=created
```

If your `/etc/hosts` file has an entry that has your local IP pointing 
at `echo.k3s.lcl` then you should be able to curl the pod.

```shell
# example of what /etc/hosts should look like on your host machine
192.168.0.1 echo.k3s.lcl 
```

Output from curl'ing `echo.k3s.lcl`:

```
# curl echo.k3s.lcl | jq .
{
  "host": {
    "hostname": "echo.k3s.lcl",
    "ip": "::ffff:10.42.0.1",
    "ips": []
  },
  "http": {
    "method": "GET",
    "baseUrl": "",
    "originalUrl": "/",
    "protocol": "http"
  },
  "request": {
    "params": {
      "0": "/"
    },
    "query": {},
    "cookies": {},
    "body": {},
    "headers": {
      "host": "echo.k3s.lcl",
      "user-agent": "curl/7.81.0",
      "accept": "*/*",
      "x-forwarded-for": "10.42.0.1",
      "x-forwarded-host": "echo.k3s.lcl",
      "x-forwarded-port": "80",
      "x-forwarded-proto": "http",
      "x-forwarded-server": "traefik-5df5cdc88d-b9tsn",
      "x-real-ip": "10.42.0.1",
      "accept-encoding": "gzip"
    }
  },
  "environment": {
    "PATH": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
    "HOSTNAME": "echo-68cd9fb7bd-zxjmd",
    "NODE_VERSION": "20.11.0",
    "YARN_VERSION": "1.22.19",
    "ECHO_SVC_PORT_3000_TCP": "tcp://10.43.41.1:3000",
    "ECHO_SVC_PORT_3000_TCP_PORT": "3000",
    "KUBERNETES_PORT": "tcp://10.43.0.1:443",
    "KUBERNETES_SERVICE_PORT_HTTPS": "443",
    "KUBERNETES_PORT_443_TCP_PROTO": "tcp",
    "KUBERNETES_PORT_443_TCP_PORT": "443",
    "KUBERNETES_PORT_443_TCP_ADDR": "10.43.0.1",
    "ECHO_SVC_PORT_3000_TCP_PROTO": "tcp",
    "KUBERNETES_SERVICE_HOST": "10.43.0.1",
    "ECHO_SVC_SERVICE_HOST": "10.43.41.1",
    "ECHO_SVC_PORT_3000_TCP_ADDR": "10.43.41.1",
    "KUBERNETES_PORT_443_TCP": "tcp://10.43.0.1:443",
    "ECHO_SVC_SERVICE_PORT": "3000",
    "ECHO_SVC_PORT": "tcp://10.43.41.1:3000",
    "KUBERNETES_SERVICE_PORT": "443",
    "HOME": "/root"
  }
}
```

**Updated**

For resources which are not kubernetes primitives such as Traefik, you
can use a Dynamic Client to create `unstructured.Unstructired{}`.

In this example we're replacing the `networking.k8s.io` `Ingress` with a
`traefik.io/v1alpha1` `IngressRoute`. 

```go
func k8sDynClient() *dynamic.DynamicClient {
	kubeConfigPath := os.Getenv("KUBECONFIG")
	fmt.Printf("Using kubeconfig: %s\n", kubeConfigPath)

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("Error getting kubernetes config: %v\n", err)
		os.Exit(1)
	}

  // Everything is the same as k8sClient except this line.
	clientset, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("error getting kubernetes config: %v\n", err)
		os.Exit(1)
	}
	return clientset
}


func createDynamicIngressRoute(client *dynamic.DynamicClient, ns string, ingressName string, port int32) error {
	ingress := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "traefik.io/v1alpha1",
			"kind":       "IngressRoute",
			"metadata": map[string]interface{}{
				"name":      ingressName,
				"namespace": ns,
			},
			"spec": map[string]interface{}{
				"entryPoints": []interface{}{
					"web",
				},
				"routes": []interface{}{
					map[string]interface{}{
						"match": "Host(`echo.k3s.lcl`) && PathPrefix(`/`)",
						"kind":  "Rule",
						"services": []interface{}{
							map[string]interface{}{
								"name":      "echo-svc",
								"port":      port,
								"namespace": "default",
								"kind":      "Service",
							},
						},
					},
				},
			},
		},
	}
	slog.Info("ingress", "ingress", ingressName, "status", "creating")
	result, err := client.Resource(schema.GroupVersionResource{
		Group:    "traefik.io",
		Version:  "v1alpha1",
		Resource: "ingressroutes",
	}).Namespace(ns).Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	slog.Info("ingress", "ingress", result.GetName(), "status", "created")
	return nil
}
```

Tags:

  #kubernetes #go
