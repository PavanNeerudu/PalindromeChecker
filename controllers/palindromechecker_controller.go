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
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "github.com/pavanneerudu/PalindromeChecker/api/v1"
)

// PalindromeCheckerReconciler reconciles a PalindromeChecker object
type PalindromeCheckerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.pavan.com,resources=palindromecheckers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.pavan.com,resources=palindromecheckers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.pavan.com,resources=palindromecheckers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PalindromeChecker object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *PalindromeCheckerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Entered PalindromeChecker Controller")

	PalindromeCheckerObject := &demov1.PalindromeChecker{}
	err := r.Get(ctx, req.NamespacedName, PalindromeCheckerObject)
	if err != nil {
		//Resource was not found.
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		//Other errors
		logger.Error(err, "Error looking up for PalindromeChecker Object")
		return ctrl.Result{}, err
	}

	logger.Info("Obtained Palindrome Checker Object")
	status := PalindromeCheckerObject.Status.Status

	if status == "" {
		PalindromeCheckerObject.Status.Status = PalindromeCheckerObject.Spec.Input
		logger.Info(fmt.Sprintf("Current Status: %s", PalindromeCheckerObject.Status.Status))
	} else {
		if status == "Palindrome" || status == "Not Palindrome" {
			logger.Info("Palindrome Checker process done!!")
			return ctrl.Result{}, nil //No more updation to etcd
		} else {
			if len(status) == 0 || len(status) == 1 {
				PalindromeCheckerObject.Status.Status = "Palindrome"
				logger.Info(fmt.Sprintf("%s is a Palindrome", PalindromeCheckerObject.Spec.Input))
			} else {
				time.Sleep(6 * time.Second)
				len := len(status)
				if status[0] == status[len-1] {
					PalindromeCheckerObject.Status.Status = status[1 : len-1]
					logger.Info(fmt.Sprintf("Current Status: %s", PalindromeCheckerObject.Status.Status))
				} else {
					PalindromeCheckerObject.Status.Status = "Not Palindrome"
					logger.Info(fmt.Sprintf("%s is not Palindrome", PalindromeCheckerObject.Spec.Input))
				}
			}
		}
	}

	//The change in status is not yet known to the cluster. So we update so, that the change is reflected in etcd
	err = r.Status().Update(ctx, PalindromeCheckerObject)
	if err != nil {
		logger.Error(err, "Error updating the status of PalindromeCheckerObject")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PalindromeCheckerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.PalindromeChecker{}).
		Complete(r)
}
