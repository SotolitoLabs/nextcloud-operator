/*
Copyright 2022.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	nextcloudoperatorv1alpha1 "github.com/SotolitoLabs/nextcloud-operator/api/v1alpha1"
)

// NextcloudReconciler reconciles a Nextcloud object
type NextcloudReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=nextcloud-operator.sotolitolabs.com,resources=nextclouds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=nextcloud-operator.sotolitolabs.com,resources=nextclouds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=nextcloud-operator.sotolitolabs.com,resources=nextclouds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Nextcloud object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *NextcloudReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Nextcloud Reconcile method: ", "context", ctx)
	// get the nextcloud resources
	nextcloud := &nextcloudoperatorv1alpha1.Nextcloud{}
	err := r.Get(ctx, req.NamespacedName, nextcloud)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Nextcloud object not found, might be under deletion process")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get Nextcloud instance")
		return ctrl.Result{}, err
	}
	logger.Info("Nextcloud instance: ", "values", nextcloud)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NextcloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&nextcloudoperatorv1alpha1.Nextcloud{}).
		Complete(r)
}
