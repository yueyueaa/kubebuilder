/*
Copyright 2023.

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

package controller

import (
	"context"

	"github.com/go-logr/logr"
	kapps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	batchv1 "tutorial.kubebuilder.io/project/api/v1"
)

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch.tutorial.kubebuilder.io,resources=apps/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("app", req.NamespacedName)

	var app batchv1.App
	if err := r.Get(ctx, req.NamespacedName, &app); err != nil {
		log.Error(err, "unable to fetch app")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	deploy := &kapps.Deployment{}
	deploy.Name = app.Name
	deploy.Namespace = app.Namespace
	deploy.Spec.Replicas = app.Spec.Replicas
	deploy.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app.kubernetes.io/name":     "app",
			"app.kubernetes.io/instance": "app-sample",
			"app.kubernetes.io/part-of":  "yueyuea",
		},
	}
	deploy.Spec.Template.ObjectMeta = metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/name":       "app",
			"app.kubernetes.io/instance":   "app-sample",
			"app.kubernetes.io/part-of":    "yueyuea",
			"app.kubernetes.io/managed-by": "kustomize",
			"app.kubernetes.io/created-by": "yueyuea",
		},
	}
	deploy.Spec.Template.Spec.Containers = []corev1.Container{
		{
			Name:  "nginx-sample",
			Image: "nginx:1.14.2",
			Ports: []corev1.ContainerPort{
				{
					ContainerPort: 80,
				},
			},
		},
	}

	if err := controllerutil.SetControllerReference(&app, deploy, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	foundDeployment := &kapps.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		log.V(1).Info("Creating Deployment", "deployment", deploy.Name)
		err = r.Create(ctx, deploy)
	} else if err == nil {
		if foundDeployment.Spec.Replicas != deploy.Spec.Replicas {
			foundDeployment.Spec.Replicas = deploy.Spec.Replicas
			log.V(1).Info("Updating Deployment", "deployment", deploy.Name)
			err = r.Update(ctx, foundDeployment)
		}
	}

	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	log.Log.Info("setup mgr")

	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.App{}).
		Owns(&kapps.Deployment{}).
		Complete(r)
}
