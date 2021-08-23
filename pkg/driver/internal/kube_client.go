/*
Copyright 2021, CTERA Networks.

Portions Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"context"

	kubefilerv1alpha1 "github.com/ctera/kubefiler-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	kubeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func GetKubeFilerExport(ctx context.Context, kubeClient kubeclient.Client, namespace, name string) (*kubefilerv1alpha1.KubeFilerExport, error) {
	kubeFilerExport := &kubefilerv1alpha1.KubeFilerExport{}
	err := kubeClient.Get(
		ctx,
		kubeclient.ObjectKey{
			Namespace: namespace,
			Name:      name,
		},
		kubeFilerExport,
	)
	if err != nil {
		return nil, err
	}
	return kubeFilerExport, nil
}

func GetKubeFiler(ctx context.Context, kubeClient kubeclient.Client, namespace, name string) (*kubefilerv1alpha1.KubeFiler, error) {
	kubeFiler := &kubefilerv1alpha1.KubeFiler{}
	err := kubeClient.Get(
		ctx,
		kubeclient.ObjectKey{
			Namespace: namespace,
			Name:      name,
		},
		kubeFiler,
	)
	if err != nil {
		return nil, err
	}
	return kubeFiler, nil
}

func GetSecret(ctx context.Context, kubeClient kubeclient.Client, namespace, name string) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	err := kubeClient.Get(
		ctx,
		kubeclient.ObjectKey{
			Namespace: namespace,
			Name:      name,
		},
		secret,
	)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func GetService(ctx context.Context, kubeClient kubeclient.Client, namespace, name string) (*corev1.Service, error) {
	service := &corev1.Service{}
	err := kubeClient.Get(
		ctx,
		kubeclient.ObjectKey{
			Namespace: namespace,
			Name:      name,
		},
		service,
	)
	if err != nil {
		return nil, err
	}
	return service, nil
}
