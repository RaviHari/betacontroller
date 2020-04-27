/*


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
	// "errors"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	betactrlv1beta1 "github.com/RaviHari/beta-controller/api/v1beta1"
)

// BetaResourceReconciler reconciles a BetaResource object
type BetaResourceReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=betactrl.betatest.io,resources=betaresources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=betactrl.betatest.io,resources=betaresources/status,verbs=get;update;patch

func (r *BetaResourceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	context := context.Background()
	log := r.Log.WithValues("betaresource", req.NamespacedName)

	// your logic here
	tinstance := &betactrlv1beta1.BetaResource{}
	if err := r.Get(context, req.NamespacedName, tinstance); err != nil {
		if errors.IsNotFound(err) {
			log.Info("tinstance document not found ", "tinstance:", req.Name)
			log.Info("reconcile completed")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get tinstance, reconcile failed")
		return ctrl.Result{}, err
	}

	log.Info("tinstance document version is", "tinstanceDocumentVersion:", tinstance.ObjectMeta.ResourceVersion)
	//Handles the panic inside the controller execution
	defer func() {
		if err := recover(); err != nil {
			log.Info("Panic occured during controller reconcilation", "tinstance", tinstance.Spec.Foo, "namespace", tinstance.ObjectMeta.Namespace, "Error:", err)
		}
	}()

	log.Info("tinstance testspec to be reconciled", "tinstance:", tinstance.Spec.Foo)
	tinstance.Status.Status = "Success"
	tinstance.Status.Message = "Created Test " + tinstance.Spec.Foo
	err := r.Update(context, tinstance)
	if err != nil {
		log.Error(err, "Update Status failed")
		return ctrl.Result{}, err
	}

	log.Info("resource status synced")

	// return ctrl.Result{}, nil

	return ctrl.Result{}, nil
}

func (r *BetaResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&betactrlv1beta1.BetaResource{}).
		Complete(r)
}
