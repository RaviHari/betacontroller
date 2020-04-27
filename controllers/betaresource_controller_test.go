/*
Copyright 2019 The Kubernetes Authors.

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
	// "time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	// "k8s.io/utils/pointer"
	// "sigs.k8s.io/controller-runtime/pkg/client"
	// "github.com/RaviHari/test-controller/controllers"
	// "sigs.k8s.io/controller-runtime/pkg/event"

	betactrlv1beta1 "github.com/RaviHari/beta-controller/api/v1beta1"
)

var _ = Context("Inside of a new namespace", func() {
	ctx := context.TODO()
	// ns := SetupTest(ctx)
	ns := &core.Namespace{}
	ns.Name = "test"

	Describe("when no resources exist", func() {

		It("should create a new Testresource with the specified name", func() {
			key := types.NamespacedName{
				Name:      "betaresource1",
				Namespace: ns.Name,
			}

			tinstance := &betactrlv1beta1.BetaResource{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "betaresource1",
					Namespace: key.Name,
				},
				Spec: betactrlv1beta1.BetaResourceSpec{
					Foo: "Bar1",
				},
			}

			// By("creating an API obj")
			// Expect(k8sClient.Create(ctx, created)).Should(Succeed())
			err := k8sClient.Create(ctx, tinstance)
			Expect(err).NotTo(HaveOccurred(), "failed to create test BetaResource resource")

			// By("deleting the created object")
			// Expect(k8sClient.Delete(ctx, created)).To(Succeed())
			// err = k8sClient.Delete(ctx, tinstance)
			// Expect(err).NotTo(HaveOccurred(), "failed to delete test BetaResource resource")
			// time.Sleep(1 * time.Second)
			// Expect(k8sClient.Get(ctx, key, tinstance)).ToNot(Succeed())

		})
	})

})
